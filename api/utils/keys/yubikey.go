//go:build piv && !pivtest

/*
Copyright 2022 Gravitational, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keys

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-piv/piv-go/piv"
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api"
	attestation "github.com/gravitational/teleport/api/gen/proto/go/attestation/v1"
	"github.com/gravitational/teleport/api/utils/retryutils"
)

const (
	// PIVCardTypeYubiKey is the PIV card type assigned to yubiKeys.
	PIVCardTypeYubiKey = "yubikey"

	// PIV connections are closed after a short delay so that the program
	// has a chance to reclaim the connection before it is closed completely.
	releaseConnectionDelay = 5 * time.Second
)

// Cache keys to prevent reconnecting to PIV module to discover a known key.
//
// Additionally, this allows the program to cache the key's PIN (if applicable)
// after the user is prompted the first time, preventing redundant prompts when
// the key is retrieved multiple times.
//
// Note: in most cases the connection caches the PIN itself, and connections can be
// reclaimed before they are fully closed (within a few seconds). However, in uncommon
// setups, this PIN caching does not actually work as expected, so we handle it instead.
// See https://github.com/go-piv/piv-go/issues/47
var (
	cachedKeys   = map[piv.Slot]*PrivateKey{}
	cachedKeysMu sync.Mutex
)

// getOrGenerateYubiKeyPrivateKey connects to a connected yubiKey and gets a private key
// matching the given touch requirement. This private key will either be newly generated
// or previously generated by a Teleport client and reused.
func getOrGenerateYubiKeyPrivateKey(ctx context.Context, requiredKeyPolicy PrivateKeyPolicy, slot PIVSlot, prompt HardwareKeyPrompt) (*PrivateKey, error) {
	if prompt == nil {
		prompt = &cliPrompt{}
	}
	cachedKeysMu.Lock()
	defer cachedKeysMu.Unlock()

	// Get the default PIV slot or the piv slot requested.
	pivSlot, err := GetDefaultKeySlot(requiredKeyPolicy)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	if slot != "" {
		pivSlot, err = slot.parse()
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}

	// If the program has already retrieved and cached this key, return it.
	if key, ok := cachedKeys[pivSlot]; ok && key.GetPrivateKeyPolicy() == requiredKeyPolicy {
		return key, nil
	}

	// Use the first yubiKey we find.
	y, err := FindYubiKey(0, prompt)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// If PIN is required, check that PIN and PUK are not the defaults.
	if requiredKeyPolicy.isHardwareKeyPINVerified() {
		if err := y.checkOrSetPIN(ctx); err != nil {
			return nil, trace.Wrap(err)
		}
	}

	promptOverwriteSlot := func(msg string) error {
		promptQuestion := fmt.Sprintf("%v\nWould you like to overwrite this slot's private key and certificate?", msg)
		if confirmed, confirmErr := prompt.ConfirmSlotOverwrite(ctx, promptQuestion); confirmErr != nil {
			return trace.Wrap(confirmErr)
		} else if !confirmed {
			return trace.Wrap(trace.CompareFailed(msg), "user declined to overwrite slot")
		}
		return nil
	}

	// If a custom slot was not specified, check for a key in the
	// default slot for the given policy and generate a new one if needed.
	if slot == "" {
		pivSlot, err = GetDefaultKeySlot(requiredKeyPolicy)
		if err != nil {
			return nil, trace.Wrap(err)
		}

		// Check the client certificate in the slot.
		switch cert, err := y.getCertificate(pivSlot); {
		case err == nil && (len(cert.Subject.Organization) == 0 || cert.Subject.Organization[0] != certOrgName):
			// Unknown cert found, prompt the user before we overwrite the slot.
			if err := promptOverwriteSlot(nonTeleportCertificateMessage(pivSlot, cert)); err != nil {
				return nil, trace.Wrap(err)
			}

			// user confirmed, generate a new key.
			fallthrough
		case errors.Is(err, piv.ErrNotFound):
			// no cert found, generate a new key.
			priv, err := y.generatePrivateKeyAndCert(pivSlot, requiredKeyPolicy)
			return priv, trace.Wrap(err)
		case err != nil:
			return nil, trace.Wrap(err)
		}
	}

	// Get the key in the slot, or generate a new one if needed.
	priv, err := y.getPrivateKey(pivSlot)
	switch {
	case err == nil && !requiredKeyPolicy.IsSatisfiedBy(priv.GetPrivateKeyPolicy()):
		// Key does not meet the required key policy, prompt the user before we overwrite the slot.
		msg := fmt.Sprintf("private key in YubiKey PIV slot %q does not meet private key policy %q.", pivSlot, requiredKeyPolicy)
		if err := promptOverwriteSlot(msg); err != nil {
			return nil, trace.Wrap(err)
		}

		// user confirmed, generate a new key.
		fallthrough
	case trace.IsNotFound(err):
		// no key found, generate a new key.
		priv, err = y.generatePrivateKeyAndCert(pivSlot, requiredKeyPolicy)
		return priv, trace.Wrap(err)
	case err != nil:
		return nil, trace.Wrap(err)
	}

	return priv, nil
}

func GetDefaultKeySlot(policy PrivateKeyPolicy) (piv.Slot, error) {
	switch policy {
	case PrivateKeyPolicyHardwareKey:
		// private_key_policy: hardware_key -> 9a
		return piv.SlotAuthentication, nil
	case PrivateKeyPolicyHardwareKeyTouch:
		// private_key_policy: hardware_key_touch -> 9c
		return piv.SlotSignature, nil
	case PrivateKeyPolicyHardwareKeyTouchAndPIN:
		// private_key_policy: hardware_key_touch_and_pin -> 9d
		return piv.SlotKeyManagement, nil
	case PrivateKeyPolicyHardwareKeyPIN:
		// private_key_policy: hardware_key_pin -> 9e
		return piv.SlotCardAuthentication, nil
	default:
		return piv.Slot{}, trace.BadParameter("unexpected private key policy %v", policy)
	}
}

func getKeyPolicies(policy PrivateKeyPolicy) (piv.TouchPolicy, piv.PINPolicy, error) {
	switch policy {
	case PrivateKeyPolicyHardwareKey:
		return piv.TouchPolicyNever, piv.PINPolicyNever, nil
	case PrivateKeyPolicyHardwareKeyTouch:
		return piv.TouchPolicyCached, piv.PINPolicyNever, nil
	case PrivateKeyPolicyHardwareKeyPIN:
		return piv.TouchPolicyNever, piv.PINPolicyOnce, nil
	case PrivateKeyPolicyHardwareKeyTouchAndPIN:
		return piv.TouchPolicyCached, piv.PINPolicyOnce, nil
	default:
		return piv.TouchPolicyNever, piv.PINPolicyNever, trace.BadParameter("unexpected private key policy %v", policy)
	}
}

func nonTeleportCertificateMessage(slot piv.Slot, cert *x509.Certificate) string {
	// Gather a small list of user-readable x509 certificate fields to display to the user.
	sum := sha256.Sum256(cert.Raw)
	fingerPrint := hex.EncodeToString(sum[:])
	return fmt.Sprintf(`Certificate in YubiKey PIV slot %q is not a Teleport client cert:
Slot %s:
	Algorithm:		%v
	Subject DN:		%v
	Issuer DN:		%v
	Serial:			%v
	Fingerprint:	%v
	Not before:		%v
	Not after:		%v
`,
		slot, slot,
		cert.SignatureAlgorithm,
		cert.Subject,
		cert.Issuer,
		cert.SerialNumber,
		fingerPrint,
		cert.NotBefore,
		cert.NotAfter,
	)
}

// YubiKeyPrivateKey is a YubiKey PIV private key. Cryptographical operations open
// a new temporary connection to the PIV card to perform the operation.
type YubiKeyPrivateKey struct {
	// YubiKey is a specific YubiKey PIV module.
	*YubiKey

	pivSlot piv.Slot
	signMux sync.Mutex

	slotCert        *x509.Certificate
	attestationCert *x509.Certificate
	attestation     *piv.Attestation
}

// yubiKeyPrivateKeyData is marshalable data used to retrieve a specific yubiKey PIV private key.
type yubiKeyPrivateKeyData struct {
	SerialNumber uint32 `json:"serial_number"`
	SlotKey      uint32 `json:"slot_key"`
}

func parseYubiKeyPrivateKeyData(keyDataBytes []byte, prompt HardwareKeyPrompt) (*PrivateKey, error) {
	if prompt == nil {
		prompt = &cliPrompt{}
	}
	cachedKeysMu.Lock()
	defer cachedKeysMu.Unlock()

	var keyData yubiKeyPrivateKeyData
	if err := json.Unmarshal(keyDataBytes, &keyData); err != nil {
		return nil, trace.Wrap(err)
	}

	pivSlot, err := parsePIVSlot(keyData.SlotKey)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// If the program has already retrieved and cached this key, return it.
	if key, ok := cachedKeys[pivSlot]; ok {
		return key, nil
	}

	y, err := FindYubiKey(keyData.SerialNumber, prompt)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	priv, err := y.getPrivateKey(pivSlot)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return priv, nil
}

// Public returns the public key corresponding to this private key.
func (y *YubiKeyPrivateKey) Public() crypto.PublicKey {
	return y.slotCert.PublicKey
}

// WarmupHardwareKey performs a bogus sign() call to prompt the user for
// a PIN/touch (if needed).
func (y *YubiKeyPrivateKey) WarmupHardwareKey(ctx context.Context) error {
	hash := sha256.Sum256(make([]byte, 256))
	_, err := y.sign(ctx, rand.Reader, hash[:], crypto.SHA256)
	return trace.Wrap(err, "failed to access a YubiKey private key")
}

// Sign implements crypto.Signer.
func (y *YubiKeyPrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signature, err := y.sign(ctx, rand, digest, opts)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return signature, nil
}

// YubiKeys require touch when signing with a private key that requires touch.
// Unfortunately, there is no good way to check whether touch is cached by the
// PIV module at a given time. In order to require touch only when needed, we
// prompt for touch after a short delay when we expect the request would succeed
// if touch were not required.
//
// There are some X factors which determine how long a request may take, such as the
// YubiKey model and firmware version, so the delays below may need to be adjusted to
// suit more models. The durations mentioned below were retrieved from testing with a
// YubiKey 5 nano (5.2.7) and a YubiKey NFC (5.4.3).
const (
	// piv.ECDSAPrivateKey.Sign consistently takes ~70 milliseconds. However, 200ms
	// should be imperceptible the the user and should avoid misfired prompts for
	// slower cards (if there are any).
	signTouchPromptDelay = time.Millisecond * 200
)

func (y *YubiKeyPrivateKey) sign(ctx context.Context, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	// To prevent concurrent calls to sign from failing due to PIV only handling a
	// single connection, use a lock to queue through signature requests one at a time.
	y.signMux.Lock()
	defer y.signMux.Unlock()
	ctx, cancel := context.WithCancelCause(ctx)

	// Lock the connection for the entire duration of the sign process.
	// Without this, the connection will be released after releaseConnectionDelay,
	// leading to a failure when providing PIN or touch input:
	// "verify pin: transmitting request: the supplied handle was invalid".
	release, err := y.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()

	var touchPromptDelayTimer *time.Timer
	if y.attestation.TouchPolicy != piv.TouchPolicyNever {
		touchPromptDelayTimer = time.NewTimer(signTouchPromptDelay)
		defer touchPromptDelayTimer.Stop()

		go func() {
			select {
			case <-touchPromptDelayTimer.C:
				// Prompt for touch after a delay, in case the function succeeds without touch due to a cached touch.
				err := y.prompt.Touch(ctx)
				if err != nil {
					// Cancel the entire function when an error occurs.
					// This is typically used for aborting the prompt.
					cancel(trace.Wrap(err))
				}
				return
			case <-ctx.Done():
				// touch cached, skip prompt.
				return
			}
		}()
	}

	promptPIN := func() (string, error) {
		// touch prompt delay is disrupted by pin prompts. To prevent misfired
		// touch prompts, pause the timer for the duration of the pin prompt.
		if touchPromptDelayTimer != nil {
			if touchPromptDelayTimer.Stop() {
				defer touchPromptDelayTimer.Reset(signTouchPromptDelay)
			}
		}
		pass, err := y.prompt.AskPIN(ctx, PINRequired)
		return pass, trace.Wrap(err)
	}

	auth := piv.KeyAuth{
		PINPrompt: promptPIN,
		PINPolicy: y.attestation.PINPolicy,
	}

	// YubiKeys with firmware version 5.3.1 have a bug where insVerify(0x20, 0x00, 0x80, nil)
	// clears the PIN cache instead of performing a non-mutable check. This causes the signature
	// with pin policy "once" to fail unless PIN is provided for each call. We can avoid this bug
	// by skipping the insVerify check and instead manually retrying with a PIN prompt only when
	// the signature fails.
	manualRetryWithPIN := false
	fw531 := piv.Version{Major: 5, Minor: 3, Patch: 1}
	if auth.PINPolicy == piv.PINPolicyOnce && y.attestation.Version == fw531 {
		// Set the keys PIN policy to never to skip the insVerify check. If PIN was provided in
		// a previous recent call, the signature will succeed as expected of the "once" policy.
		auth.PINPolicy = piv.PINPolicyNever
		manualRetryWithPIN = true
	}

	privateKey, err := y.privateKey(y.pivSlot, y.Public(), auth)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	signer, ok := privateKey.(crypto.Signer)
	if !ok {
		return nil, trace.BadParameter("private key type %T does not implement crypto.Signer", privateKey)
	}

	// For generic auth errors, such as when PIN is not provided, the smart card returns the error code 0x6982.
	// The piv-go library wraps error codes like this with a user readable message: "security status not satisfied".
	const pivGenericAuthErrCodeString = "6982"

	signature, err := abandonableSign(ctx, signer, rand, digest, opts)
	switch {
	case err == nil:
		return signature, nil
	case manualRetryWithPIN && strings.Contains(err.Error(), pivGenericAuthErrCodeString):
		pin, err := promptPIN()
		if err != nil {
			return nil, trace.Wrap(err)
		}
		if err := y.verifyPIN(pin); err != nil {
			return nil, trace.Wrap(err)
		}
		signature, err := abandonableSign(ctx, signer, rand, digest, opts)
		return signature, trace.Wrap(err)
	default:
		return nil, trace.Wrap(err)
	}
}

// abandonableSign is a wrapper around signer.Sign.
// It enhances the functionality of signer.Sign by allowing the caller to stop
// waiting for the result if the provided context is canceled.
// It is especially important for WarmupHardwareKey,
// where waiting for the user providing a PIN/touch could block program termination.
// Important: this function only abandons the signer.Sign result, doesn't cancel it.
func abandonableSign(ctx context.Context, signer crypto.Signer, rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	type signResult struct {
		signature []byte
		err       error
	}

	signResultCh := make(chan signResult)
	go func() {
		if err := ctx.Err(); err != nil {
			return
		}
		signature, err := signer.Sign(rand, digest, opts)
		select {
		case <-ctx.Done():
		case signResultCh <- signResult{signature: signature, err: trace.Wrap(err)}:
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-signResultCh:
		return result.signature, trace.Wrap(result.err)
	}
}

func (y *YubiKeyPrivateKey) toPrivateKey() (*PrivateKey, error) {
	keyPEM, err := y.keyPEM()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return NewPrivateKey(y, keyPEM)
}

func (y *YubiKeyPrivateKey) keyPEM() ([]byte, error) {
	keyDataBytes, err := json.Marshal(yubiKeyPrivateKeyData{
		SerialNumber: y.serialNumber,
		SlotKey:      y.pivSlot.Key,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:    pivYubiKeyPrivateKeyType,
		Headers: nil,
		Bytes:   keyDataBytes,
	}), nil
}

// GetAttestationStatement returns an AttestationStatement for this YubiKeyPrivateKey.
func (y *YubiKeyPrivateKey) GetAttestationStatement() *AttestationStatement {
	return &AttestationStatement{
		AttestationStatement: &attestation.AttestationStatement_YubikeyAttestationStatement{
			YubikeyAttestationStatement: &attestation.YubiKeyAttestationStatement{
				SlotCert:        y.slotCert.Raw,
				AttestationCert: y.attestationCert.Raw,
			},
		},
	}
}

// GetPrivateKeyPolicy returns the PrivateKeyPolicy supported by this YubiKeyPrivateKey.
func (y *YubiKeyPrivateKey) GetPrivateKeyPolicy() PrivateKeyPolicy {
	return GetPrivateKeyPolicyFromAttestation(y.attestation)
}

// GetPrivateKeyPolicyFromAttestation returns the PrivateKeyPolicy satisfied by the given hardware key attestation.
func GetPrivateKeyPolicyFromAttestation(att *piv.Attestation) PrivateKeyPolicy {
	isTouchPolicy := att.TouchPolicy == piv.TouchPolicyCached ||
		att.TouchPolicy == piv.TouchPolicyAlways

	isPINPolicy := att.PINPolicy == piv.PINPolicyOnce ||
		att.PINPolicy == piv.PINPolicyAlways

	switch {
	case isPINPolicy && isTouchPolicy:
		return PrivateKeyPolicyHardwareKeyTouchAndPIN
	case isPINPolicy:
		return PrivateKeyPolicyHardwareKeyPIN
	case isTouchPolicy:
		return PrivateKeyPolicyHardwareKeyTouch
	default:
		return PrivateKeyPolicyHardwareKey
	}
}

// YubiKey is a specific YubiKey PIV card.
type YubiKey struct {
	// conn is a shared YubiKey PIV connection.
	//
	// PIV connections claim an exclusive lock on the PIV module until closed.
	// In order to improve connection sharing for this program without locking
	// out other programs during extended program executions (like "tsh proxy ssh"),
	// this connections is opportunistically formed and released after being
	// unused for a few seconds.
	*sharedPIVConnection
	// serialNumber is the yubiKey's 8 digit serial number.
	serialNumber uint32
	prompt       HardwareKeyPrompt
}

func newYubiKey(card string, prompt HardwareKeyPrompt) (*YubiKey, error) {
	y := &YubiKey{
		sharedPIVConnection: &sharedPIVConnection{
			card: card,
		},
		prompt: prompt,
	}

	serialNumber, err := y.serial()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	y.serialNumber = serialNumber
	return y, nil
}

// Reset resets the YubiKey PIV module to default settings.
func (y *YubiKey) Reset() error {
	err := y.reset()
	return trace.Wrap(err)
}

// generatePrivateKeyAndCert generates a new private key and client metadata cert in the given PIV slot.
func (y *YubiKey) generatePrivateKeyAndCert(slot piv.Slot, requiredKeyPolicy PrivateKeyPolicy) (*PrivateKey, error) {
	if err := y.generatePrivateKey(slot, requiredKeyPolicy); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := y.SetMetadataCertificate(slot, pkix.Name{
		Organization:       []string{certOrgName},
		OrganizationalUnit: []string{api.Version},
	}); err != nil {
		return nil, trace.Wrap(err)
	}

	return y.getPrivateKey(slot)
}

// SetMetadataCertificate creates a self signed certificate and stores it in the YubiKey's
// PIV certificate slot. This certificate is purely used as metadata to determine when a
// slot is in used by a Teleport Client and is not fit to be used in cryptographic operations.
// This cert is also useful for users to discern where the key came with tools like `ykman piv info`.
func (y *YubiKey) SetMetadataCertificate(slot piv.Slot, subject pkix.Name) error {
	cert, err := SelfSignedMetadataCertificate(subject)
	if err != nil {
		return trace.Wrap(err)
	}

	err = y.setCertificate(piv.DefaultManagementKey, slot, cert)
	return trace.Wrap(err)
}

// getCertificate gets a certificate from the given PIV slot.
func (y *YubiKey) getCertificate(slot piv.Slot) (*x509.Certificate, error) {
	cert, err := y.certificate(slot)
	return cert, trace.Wrap(err)
}

// generatePrivateKey generates a new private key in the given PIV slot.
func (y *YubiKey) generatePrivateKey(slot piv.Slot, requiredKeyPolicy PrivateKeyPolicy) error {
	touchPolicy, pinPolicy, err := getKeyPolicies(requiredKeyPolicy)
	if err != nil {
		return trace.Wrap(err)
	}

	opts := piv.Key{
		Algorithm:   piv.AlgorithmEC256,
		PINPolicy:   pinPolicy,
		TouchPolicy: touchPolicy,
	}

	_, err = y.generateKey(piv.DefaultManagementKey, slot, opts)
	return trace.Wrap(err)
}

// getPrivateKey gets an existing private key from the given PIV slot.
func (y *YubiKey) getPrivateKey(slot piv.Slot) (*PrivateKey, error) {
	slotCert, err := y.attest(slot)
	if errors.Is(err, piv.ErrNotFound) {
		return nil, trace.NotFound("private key in YubiKey PIV slot %q not found.", slot.String())
	} else if err != nil {
		return nil, trace.Wrap(err)
	}

	attCert, err := y.attestationCertificate()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	attestation, err := piv.Verify(attCert, slotCert)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	priv := &YubiKeyPrivateKey{
		YubiKey:         y,
		pivSlot:         slot,
		slotCert:        slotCert,
		attestationCert: attCert,
		attestation:     attestation,
	}

	keyPEM, err := priv.keyPEM()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	key, err := NewPrivateKey(priv, keyPEM)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	cachedKeys[slot] = key
	return key, nil
}

// SetPIN sets the YubiKey PIV PIN. This doesn't require user interaction like touch, just the correct old PIN.
func (y *YubiKey) SetPIN(oldPin, newPin string) error {
	err := y.setPIN(oldPin, newPin)
	return trace.Wrap(err)
}

// checkOrSetPIN prompts the user for PIN and verifies it with the YubiKey.
// If the user provides the default PIN, they will be prompted to set a
// non-default PIN and PUK before continuing.
func (y *YubiKey) checkOrSetPIN(ctx context.Context) error {
	pin, err := y.prompt.AskPIN(ctx, PINOptional)
	if err != nil {
		return trace.Wrap(err)
	}

	switch pin {
	case piv.DefaultPIN:
		fmt.Fprintf(os.Stderr, "The default PIN %q is not supported.\n", piv.DefaultPIN)
		fallthrough
	case "":
		if pin, err = y.setPINAndPUKFromDefault(ctx, y.prompt); err != nil {
			return trace.Wrap(err)
		}
	}

	return trace.Wrap(y.verifyPIN(pin))
}

type sharedPIVConnection struct {
	// card is a reader name used to find and connect to this yubiKey.
	// This value may change between OS's, or with other system changes.
	card string

	// conn is the shared PIV connection.
	conn              *piv.YubiKey
	mu                sync.Mutex
	activeConnections int
}

// connect establishes a connection to a YubiKey PIV module and returns a release function.
// The release function should be called to properly close the shared connection.
// The connection is not immediately terminated, allowing other callers to
// use it before it's released.
// The YubiKey PIV module itself takes some additional time to handle closed
// connections, so we use a retry loop to give the PIV module time to close prior connections.
func (c *sharedPIVConnection) connect() (func(), error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	release := func() {
		go func() {
			time.Sleep(releaseConnectionDelay)

			c.mu.Lock()
			defer c.mu.Unlock()

			c.activeConnections--
			if c.activeConnections == 0 {
				c.conn.Close()
				c.conn = nil
			}
		}()
	}

	if c.conn != nil {
		c.activeConnections++
		return release, nil
	}

	linearRetry, err := retryutils.NewLinear(retryutils.LinearConfig{
		// If a PIV connection has just been closed, it take ~5 ms to become
		// available to new connections. For this reason, we initially wait a
		// short 10ms before stepping up to a longer 50ms retry.
		First: time.Millisecond * 10,
		Step:  time.Millisecond * 10,
		// Since PIV modules only allow a single connection, it is a bottleneck
		// resource. To maximize usage, we use a short 50ms retry to catch the
		// connection opening up as soon as possible.
		Max: time.Millisecond * 50,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// Backoff and retry for a time slightly greater than releaseConnectionDelay.
	retryCtx, cancel := context.WithTimeout(context.Background(), releaseConnectionDelay+100*time.Millisecond)
	defer cancel()

	err = linearRetry.For(retryCtx, func() error {
		c.conn, err = piv.Open(c.card)
		if err != nil && !isRetryError(err) {
			return retryutils.PermanentRetryError(err)
		}
		return trace.Wrap(err)
	})
	if trace.IsLimitExceeded(err) {
		// Using PIV synchronously causes issues since only one connection is allowed at a time.
		// This shouldn't be an issue for `tsh` which primarily runs consecutively, but Teleport
		// Connect works through callbacks, etc. and may try to open multiple connections at a time.
		// If this error is being emitted more than rarely, the 1 second timeout may need to be increased.
		//
		// It's also possible that the user is running another PIV program, which may hold the PIV
		// connection indefinitely (yubikey-agent). In this case, user action is necessary, so we
		// alert them with this issue.
		return nil, trace.LimitExceeded("could not connect to YubiKey as another application is using it. Please try again once the program that uses the YubiKey, such as yubikey-agent is closed")
	} else if err != nil {
		return nil, trace.Wrap(err)
	}

	c.activeConnections++
	return release, nil
}

func (c *sharedPIVConnection) privateKey(slot piv.Slot, public crypto.PublicKey, auth piv.KeyAuth) (crypto.PrivateKey, error) {
	release, err := c.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()
	privateKey, err := c.conn.PrivateKey(slot, public, auth)
	return privateKey, trace.Wrap(err)
}

func (c *sharedPIVConnection) serial() (uint32, error) {
	release, err := c.connect()
	if err != nil {
		return 0, trace.Wrap(err)
	}
	defer release()
	serial, err := c.conn.Serial()
	return serial, trace.Wrap(err)
}

func (c *sharedPIVConnection) reset() error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	// Clear cached keys.
	cachedKeys = make(map[piv.Slot]*PrivateKey)
	return trace.Wrap(c.conn.Reset())
}

func (c *sharedPIVConnection) setCertificate(key [24]byte, slot piv.Slot, cert *x509.Certificate) error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	return trace.Wrap(c.conn.SetCertificate(key, slot, cert))
}

func (c *sharedPIVConnection) certificate(slot piv.Slot) (*x509.Certificate, error) {
	release, err := c.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()
	cert, err := c.conn.Certificate(slot)
	return cert, trace.Wrap(err)
}

func (c *sharedPIVConnection) generateKey(key [24]byte, slot piv.Slot, opts piv.Key) (crypto.PublicKey, error) {
	release, err := c.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()
	pubKey, err := c.conn.GenerateKey(key, slot, opts)
	return pubKey, trace.Wrap(err)
}

func (c *sharedPIVConnection) attest(slot piv.Slot) (*x509.Certificate, error) {
	release, err := c.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()
	cert, err := c.conn.Attest(slot)
	return cert, trace.Wrap(err)
}

func (c *sharedPIVConnection) attestationCertificate() (*x509.Certificate, error) {
	release, err := c.connect()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	defer release()
	cert, err := c.conn.AttestationCertificate()
	return cert, trace.Wrap(err)
}

func (c *sharedPIVConnection) setPIN(oldPIN string, newPIN string) error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	return trace.Wrap(c.conn.SetPIN(oldPIN, newPIN))
}

func (c *sharedPIVConnection) setPUK(oldPUK string, newPUK string) error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	return trace.Wrap(c.conn.SetPUK(oldPUK, newPUK))
}

func (c *sharedPIVConnection) unblock(puk string, newPIN string) error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	return trace.Wrap(c.conn.Unblock(puk, newPIN))
}

func (c *sharedPIVConnection) verifyPIN(pin string) error {
	release, err := c.connect()
	if err != nil {
		return trace.Wrap(err)
	}
	defer release()
	return trace.Wrap(c.conn.VerifyPIN(pin))
}

func (c *sharedPIVConnection) setPINAndPUKFromDefault(ctx context.Context, prompt HardwareKeyPrompt) (string, error) {
	pinAndPUK, err := prompt.ChangePIN(ctx)
	if err != nil {
		return "", trace.Wrap(err)
	}
	// YubiKey requires that PIN and PUK be 6-8 characters.
	// Verify that we get valid values from the prompt.
	if !isPINLengthValid(pinAndPUK.PIN) {
		return "", trace.BadParameter("PIN must be 6-8 characters long")
	}
	if pinAndPUK.PIN == piv.DefaultPIN {
		return "", trace.BadParameter("The default PIN is not supported")
	}
	if !isPINLengthValid(pinAndPUK.PUK) {
		return "", trace.BadParameter("PUK must be 6-8 characters long")
	}
	if pinAndPUK.PUK == piv.DefaultPUK {
		return "", trace.BadParameter("The default PUK is not supported")
	}

	if pinAndPUK.PUKChanged {
		if err := c.setPUK(piv.DefaultPUK, pinAndPUK.PUK); err != nil {
			return "", trace.Wrap(err)
		}
	}

	if err := c.unblock(pinAndPUK.PUK, pinAndPUK.PIN); err != nil {
		return "", trace.Wrap(err)
	}

	return pinAndPUK.PIN, nil
}

func isRetryError(err error) bool {
	const retryError = "connecting to smart card: the smart card cannot be accessed because of other connections outstanding"
	return strings.Contains(err.Error(), retryError)
}

// FindYubiKey finds a yubiKey PIV card by serial number. If no serial
// number is provided, the first yubiKey found will be returned.
func FindYubiKey(serialNumber uint32, prompt HardwareKeyPrompt) (*YubiKey, error) {
	yubiKeyCards, err := findYubiKeyCards()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if len(yubiKeyCards) == 0 {
		if serialNumber != 0 {
			return nil, trace.ConnectionProblem(nil, "no YubiKey device connected with serial number %d", serialNumber)
		}
		return nil, trace.ConnectionProblem(nil, "no YubiKey device connected")
	}

	for _, card := range yubiKeyCards {
		y, err := newYubiKey(card, prompt)
		if err != nil {
			return nil, trace.Wrap(err)
		}

		if serialNumber == 0 || y.serialNumber == serialNumber {
			return y, nil
		}
	}

	return nil, trace.ConnectionProblem(nil, "no YubiKey device connected with serial number %d", serialNumber)
}

// findYubiKeyCards returns a list of connected yubiKey PIV card names.
func findYubiKeyCards() ([]string, error) {
	cards, err := piv.Cards()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	var yubiKeyCards []string
	for _, card := range cards {
		if strings.Contains(strings.ToLower(card), PIVCardTypeYubiKey) {
			yubiKeyCards = append(yubiKeyCards, card)
		}
	}

	return yubiKeyCards, nil
}

func (s PIVSlot) validate() error {
	_, err := s.parse()
	return trace.Wrap(err)
}

func (s PIVSlot) parse() (piv.Slot, error) {
	slotKey, err := strconv.ParseUint(string(s), 16, 32)
	if err != nil {
		return piv.Slot{}, trace.Wrap(err)
	}

	return parsePIVSlot(uint32(slotKey))
}

func parsePIVSlot(slotKey uint32) (piv.Slot, error) {
	switch slotKey {
	case piv.SlotAuthentication.Key:
		return piv.SlotAuthentication, nil
	case piv.SlotSignature.Key:
		return piv.SlotSignature, nil
	case piv.SlotKeyManagement.Key:
		return piv.SlotKeyManagement, nil
	case piv.SlotCardAuthentication.Key:
		return piv.SlotCardAuthentication, nil
	default:
		retiredSlot, ok := piv.RetiredKeyManagementSlot(slotKey)
		if !ok {
			return piv.Slot{}, trace.BadParameter("slot %X does not exist", slotKey)
		}
		return retiredSlot, nil
	}
}

// certOrgName is used to identify Teleport Client self-signed certificates stored in yubiKey PIV slots.
const certOrgName = "teleport"

func SelfSignedMetadataCertificate(subject pkix.Name) (*x509.Certificate, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit) // see crypto/tls/generate_cert.go
	if err != nil {
		return nil, trace.Wrap(err)
	}
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		PublicKey:    priv.Public(),
	}

	if cert.Raw, err = x509.CreateCertificate(rand.Reader, cert, cert, priv.Public(), priv); err != nil {
		return nil, trace.Wrap(err)
	}
	return cert, nil
}

// IsHardware returns true if [k] is a hardware PIV key.
func (k *PrivateKey) IsHardware() bool {
	switch k.Signer.(type) {
	case *YubiKeyPrivateKey:
		return true
	}
	return false
}
