/*
Copyright 2021 Gravitational, Inc.

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

package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/defaults"
	"github.com/gravitational/teleport/api/utils"
	"github.com/gravitational/teleport/api/utils/keys"
	"github.com/gravitational/teleport/api/utils/tlsutils"
)

var (
	// ErrPasswordlessRequiresWebauthn is issued if a passwordless challenge is
	// requested but WebAuthn isn't enabled.
	ErrPasswordlessRequiresWebauthn = &trace.BadParameterError{
		Message: "passwordless requires WebAuthn",
	}

	// ErrPasswordlessDisabledBySettings is issued if a passwordless challenge is
	// requested but passwordless is disabled by cluster settings.
	// See AuthPreferenceV2.AuthPreferenceV2.
	ErrPasswordlessDisabledBySettings = &trace.BadParameterError{
		Message: "passwordless disabled by cluster settings",
	}

	// ErrPassswordlessLoginBySSOUser is issued if an SSO user tries to login
	// using passwordless.
	ErrPassswordlessLoginBySSOUser = &trace.AccessDeniedError{
		Message: "SSO user cannot login using passwordless",
	}
)

// AuthPreference defines the authentication preferences for a specific
// cluster. It defines the type (local, oidc) and second factor (off, otp, oidc).
// AuthPreference is a configuration resource, never create more than one instance
// of it.
type AuthPreference interface {
	// Resource provides common resource properties.
	ResourceWithOrigin

	// GetType gets the type of authentication: local, saml, or oidc.
	GetType() string
	// SetType sets the type of authentication: local, saml, or oidc.
	SetType(string)

	// GetSecondFactor gets the type of second factor.
	// Deprecated: use GetSecondFactors
	GetSecondFactor() constants.SecondFactorType
	// SetSecondFactor sets the type of second factor.
	// Deprecated: use SetSecondFactors
	SetSecondFactor(constants.SecondFactorType)
	// GetSecondFactors gets a list of supported second factors.
	GetSecondFactors() []SecondFactorType
	// SetSecondFactors sets the list of supported second factors.
	SetSecondFactors(...SecondFactorType)
	// GetPreferredLocalMFA returns the preferred MFA method from available methods.
	// This is often used as a tip to clients to pick an MFA method.
	// Preference order: WebAuthn > OTP
	GetPreferredLocalMFA() string
	// IsSecondFactorEnabled checks if second factor is enabled.
	IsSecondFactorEnabled() bool
	// IsSecondFactorEnforced checks if second factor is enforced.
	IsSecondFactorEnforced() bool
	// IsSecondFactorTOTPAllowed checks if users are allowed to register TOTP devices.
	IsSecondFactorTOTPAllowed() bool
	// IsSecondFactorWebauthnAllowed checks if users are allowed to register
	// Webauthn devices.
	IsSecondFactorWebauthnAllowed() bool
	// IsAdminActionMFAEnforced checks if admin action MFA is enforced.
	IsAdminActionMFAEnforced() bool

	// GetConnectorName gets the name of the OIDC or SAML connector to use. If
	// this value is empty, we fall back to the first connector in the backend.
	GetConnectorName() string
	// SetConnectorName sets the name of the OIDC or SAML connector to use. If
	// this value is empty, we fall back to the first connector in the backend.
	SetConnectorName(string)

	// GetU2F gets the U2F configuration settings.
	GetU2F() (*U2F, error)
	// SetU2F sets the U2F configuration settings.
	SetU2F(*U2F)

	// GetWebauthn returns the Webauthn configuration settings.
	GetWebauthn() (*Webauthn, error)
	// SetWebauthn sets the Webauthn configuration settings.
	SetWebauthn(*Webauthn)

	// GetAllowPasswordless returns if passwordless is allowed by cluster
	// settings.
	GetAllowPasswordless() bool
	// SetAllowPasswordless sets the value of the allow passwordless setting.
	SetAllowPasswordless(b bool)

	// GetAllowHeadless returns if headless is allowed by cluster settings.
	GetAllowHeadless() bool
	// SetAllowHeadless sets the value of the allow headless setting.
	SetAllowHeadless(b bool)

	// SetRequireMFAType sets the type of MFA requirement enforced for this cluster.
	SetRequireMFAType(RequireMFAType)
	// GetRequireMFAType returns the type of MFA requirement enforced for this cluster.
	GetRequireMFAType() RequireMFAType

	// GetPrivateKeyPolicy returns the configured private key policy for the cluster.
	GetPrivateKeyPolicy() keys.PrivateKeyPolicy

	// GetHardwareKey returns the hardware key settings configured for the cluster.
	GetHardwareKey() (*HardwareKey, error)
	// GetPIVSlot returns the configured piv slot for the cluster.
	GetPIVSlot() keys.PIVSlot
	// GetHardwareKeySerialNumberValidation returns the cluster's hardware key
	// serial number validation settings.
	GetHardwareKeySerialNumberValidation() (*HardwareKeySerialNumberValidation, error)

	// GetDisconnectExpiredCert returns disconnect expired certificate setting
	GetDisconnectExpiredCert() bool
	// SetDisconnectExpiredCert sets disconnect client with expired certificate setting
	SetDisconnectExpiredCert(bool)

	// GetAllowLocalAuth gets if local authentication is allowed.
	GetAllowLocalAuth() bool
	// SetAllowLocalAuth sets if local authentication is allowed.
	SetAllowLocalAuth(bool)

	// GetMessageOfTheDay fetches the MOTD
	GetMessageOfTheDay() string
	// SetMessageOfTheDay sets the MOTD
	SetMessageOfTheDay(string)

	// GetLockingMode gets the cluster-wide locking mode default.
	GetLockingMode() constants.LockingMode
	// SetLockingMode sets the cluster-wide locking mode default.
	SetLockingMode(constants.LockingMode)

	// GetDeviceTrust returns the cluster device trust settings, or nil if no
	// explicit configurations are present.
	GetDeviceTrust() *DeviceTrust
	// SetDeviceTrust sets the cluster device trust settings.
	SetDeviceTrust(*DeviceTrust)

	// IsSAMLIdPEnabled returns true if the SAML IdP is enabled.
	IsSAMLIdPEnabled() bool
	// SetSAMLIdPEnabled sets the SAML IdP to enabled.
	SetSAMLIdPEnabled(bool)

	// GetDefaultSessionTTL retrieves the max session ttl
	GetDefaultSessionTTL() Duration
	// SetDefaultSessionTTL sets the max session ttl
	SetDefaultSessionTTL(Duration)

	// GetOktaSyncPeriod returns the duration between Okta synchronization calls if the Okta service is running.
	GetOktaSyncPeriod() time.Duration
	// SetOktaSyncPeriod sets the duration between Okta synchronzation calls.
	SetOktaSyncPeriod(timeBetweenSyncs time.Duration)

	// GetSignatureAlgorithmSuite gets the signature algorithm suite.
	GetSignatureAlgorithmSuite() SignatureAlgorithmSuite
	// SetSignatureAlgorithmSuite sets the signature algorithm suite.
	SetSignatureAlgorithmSuite(SignatureAlgorithmSuite)
	// SetDefaultSignatureAlgorithmSuite sets default signature algorithm suite
	// based on the params. This is meant for a default auth preference in a
	// brand new cluster or after resetting the auth preference.
	SetDefaultSignatureAlgorithmSuite(SignatureAlgorithmSuiteParams)
	// CheckSignatureAlgorithmSuite returns an error if the current signature
	// algorithm suite is incompatible with [params].
	CheckSignatureAlgorithmSuite(SignatureAlgorithmSuiteParams) error

	// String represents a human readable version of authentication settings.
	String() string

	// Clone makes a deep copy of the AuthPreference.
	Clone() AuthPreference
}

// NewAuthPreference is a convenience method to to create AuthPreferenceV2.
func NewAuthPreference(spec AuthPreferenceSpecV2) (AuthPreference, error) {
	return newAuthPreferenceWithLabels(spec, map[string]string{})
}

// NewAuthPreferenceFromConfigFile is a convenience method to create
// AuthPreferenceV2 labeled as originating from config file.
func NewAuthPreferenceFromConfigFile(spec AuthPreferenceSpecV2) (AuthPreference, error) {
	return newAuthPreferenceWithLabels(spec, map[string]string{
		OriginLabel: OriginConfigFile,
	})
}

// NewAuthPreferenceWithLabels is a convenience method to create
// AuthPreferenceV2 with a specific map of labels.
func newAuthPreferenceWithLabels(spec AuthPreferenceSpecV2, labels map[string]string) (AuthPreference, error) {
	pref := &AuthPreferenceV2{
		Metadata: Metadata{
			Labels: labels,
		},
		Spec: spec,
	}
	if err := pref.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}
	return pref, nil
}

// DefaultAuthPreference returns the default authentication preferences.
func DefaultAuthPreference() AuthPreference {
	authPref, _ := newAuthPreferenceWithLabels(AuthPreferenceSpecV2{
		// This is useful as a static value, but the real default signature
		// algorithm suite depends on the cluster FIPS and HSM settings, and
		// gets written by [AuthPreferenceV2.SetDefaultSignatureAlgorithmSuite]
		// wherever a default auth preference will actually be persisted.
		// It is set here so that many existing tests using this get the
		// benefits of the balanced-v1 suite.
		SignatureAlgorithmSuite: SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_BALANCED_V1,
	}, map[string]string{
		OriginLabel: OriginDefaults,
	})
	return authPref
}

// GetVersion returns resource version.
func (c *AuthPreferenceV2) GetVersion() string {
	return c.Version
}

// GetName returns the name of the resource.
func (c *AuthPreferenceV2) GetName() string {
	return c.Metadata.Name
}

// SetName sets the name of the resource.
func (c *AuthPreferenceV2) SetName(e string) {
	c.Metadata.Name = e
}

// SetExpiry sets expiry time for the object.
func (c *AuthPreferenceV2) SetExpiry(expires time.Time) {
	c.Metadata.SetExpiry(expires)
}

// Expiry returns object expiry setting.
func (c *AuthPreferenceV2) Expiry() time.Time {
	return c.Metadata.Expiry()
}

// GetMetadata returns object metadata.
func (c *AuthPreferenceV2) GetMetadata() Metadata {
	return c.Metadata
}

// GetRevision returns the revision
func (c *AuthPreferenceV2) GetRevision() string {
	return c.Metadata.GetRevision()
}

// SetRevision sets the revision
func (c *AuthPreferenceV2) SetRevision(rev string) {
	c.Metadata.SetRevision(rev)
}

// Origin returns the origin value of the resource.
func (c *AuthPreferenceV2) Origin() string {
	return c.Metadata.Origin()
}

// SetOrigin sets the origin value of the resource.
func (c *AuthPreferenceV2) SetOrigin(origin string) {
	c.Metadata.SetOrigin(origin)
}

// GetKind returns resource kind.
func (c *AuthPreferenceV2) GetKind() string {
	return c.Kind
}

// GetSubKind returns resource subkind.
func (c *AuthPreferenceV2) GetSubKind() string {
	return c.SubKind
}

// SetSubKind sets resource subkind.
func (c *AuthPreferenceV2) SetSubKind(sk string) {
	c.SubKind = sk
}

// GetType returns the type of authentication.
func (c *AuthPreferenceV2) GetType() string {
	return c.Spec.Type
}

// SetType sets the type of authentication.
func (c *AuthPreferenceV2) SetType(s string) {
	c.Spec.Type = s
}

// GetSecondFactors gets a list of supported second factors.
func (c *AuthPreferenceV2) GetSecondFactors() []SecondFactorType {
	if c.Spec.SecondFactors != nil {
		return c.Spec.SecondFactors
	}

	return SecondFactorsFromLegacySecondFactor(c.Spec.SecondFactor, c.Spec.Webauthn != nil)
}

// SecondFactorsFromLegacySecondFactor returns the list of SecondFactorTypes supported by the given second factor type.
func SecondFactorsFromLegacySecondFactor(sf constants.SecondFactorType, webauthnConfigured bool) []SecondFactorType {
	switch sf {
	case constants.SecondFactorOff:
		return nil
	case constants.SecondFactorOptional, constants.SecondFactorOn:
		if !webauthnConfigured {
			return []SecondFactorType{SecondFactorType_SECOND_FACTOR_TYPE_OTP}
		}
		return []SecondFactorType{SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN, SecondFactorType_SECOND_FACTOR_TYPE_OTP}
	case constants.SecondFactorOTP:
		return []SecondFactorType{SecondFactorType_SECOND_FACTOR_TYPE_OTP}
	case constants.SecondFactorWebauthn:
		return []SecondFactorType{SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN}
	default:
		slog.WarnContext(context.Background(), "Found unknown second_factor setting", "second_factor", sf)
		return nil
	}
}

// SetSecondFactors sets the list of supported second factors.
func (c *AuthPreferenceV2) SetSecondFactors(s ...SecondFactorType) {
	c.Spec.SecondFactors = s
}

// GetSecondFactor returns the type of second factor.
func (c *AuthPreferenceV2) GetSecondFactor() constants.SecondFactorType {
	return c.Spec.SecondFactor
}

// SetSecondFactor sets the type of second factor.
func (c *AuthPreferenceV2) SetSecondFactor(s constants.SecondFactorType) {
	c.Spec.SecondFactor = s
}

// GetPreferredLocalMFA returns the preferred MFA method from available methods.
// This is often used as a tip to clients to pick an MFA method.
// Preference order: WebAuthn > OTP
func (c *AuthPreferenceV2) GetPreferredLocalMFA() string {
	if c.IsSecondFactorWebauthnAllowed() {
		return SecondFactorTypeWebauthnString
	}

	if c.IsSecondFactorTOTPAllowed() {
		return SecondFactorTypeOTPString
	}

	return ""
}

// IsSecondFactorEnforced checks if second factor is enabled.
//
// TODO(Joerger): outside of tests, second factor should always be enabled.
// All calls should be removed and the old off/optional second factors removed.
func (c *AuthPreferenceV2) IsSecondFactorEnabled() bool {
	return len(c.GetSecondFactors()) > 0
}

// IsSecondFactorEnforced checks if second factor is enforced.
//
// TODO(Joerger): outside of tests, second factor should always be enforced.
// All calls should be removed and the old off/optional second factors removed.
func (c *AuthPreferenceV2) IsSecondFactorEnforced() bool {
	return len(c.GetSecondFactors()) > 0 && c.Spec.SecondFactor != constants.SecondFactorOptional
}

// IsSecondFactorTOTPAllowed checks if users are allowed to register TOTP devices.
func (c *AuthPreferenceV2) IsSecondFactorTOTPAllowed() bool {
	return slices.Contains(c.GetSecondFactors(), SecondFactorType_SECOND_FACTOR_TYPE_OTP)
}

// IsSecondFactorWebauthnAllowed checks if users are allowed to register
// Webauthn devices.
func (c *AuthPreferenceV2) IsSecondFactorWebauthnAllowed() bool {
	return slices.Contains(c.GetSecondFactors(), SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN)
}

// IsAdminActionMFAEnforced checks if admin action MFA is enforced.
func (c *AuthPreferenceV2) IsAdminActionMFAEnforced() bool {
	// OTP is not supported for Admin MFA.
	return c.IsSecondFactorEnforced() && !c.IsSecondFactorTOTPAllowed()
}

// GetConnectorName gets the name of the OIDC or SAML connector to use. If
// this value is empty, we fall back to the first connector in the backend.
func (c *AuthPreferenceV2) GetConnectorName() string {
	return c.Spec.ConnectorName
}

// SetConnectorName sets the name of the OIDC or SAML connector to use. If
// this value is empty, we fall back to the first connector in the backend.
func (c *AuthPreferenceV2) SetConnectorName(cn string) {
	c.Spec.ConnectorName = cn
}

// GetU2F gets the U2F configuration settings.
func (c *AuthPreferenceV2) GetU2F() (*U2F, error) {
	if c.Spec.U2F == nil {
		return nil, trace.NotFound("U2F is not configured in this cluster, please contact your administrator and ask them to follow https://goteleport.com/docs/access-controls/guides/u2f/")
	}
	return c.Spec.U2F, nil
}

// SetU2F sets the U2F configuration settings.
func (c *AuthPreferenceV2) SetU2F(u2f *U2F) {
	c.Spec.U2F = u2f
}

func (c *AuthPreferenceV2) GetWebauthn() (*Webauthn, error) {
	if c.Spec.Webauthn == nil {
		return nil, trace.NotFound("Webauthn is not configured in this cluster, please contact your administrator and ask them to follow https://goteleport.com/docs/access-controls/guides/webauthn/")
	}
	return c.Spec.Webauthn, nil
}

func (c *AuthPreferenceV2) SetWebauthn(w *Webauthn) {
	c.Spec.Webauthn = w
}

func (c *AuthPreferenceV2) GetAllowPasswordless() bool {
	return c.Spec.AllowPasswordless != nil && c.Spec.AllowPasswordless.Value
}

func (c *AuthPreferenceV2) SetAllowPasswordless(b bool) {
	c.Spec.AllowPasswordless = NewBoolOption(b)
}

func (c *AuthPreferenceV2) GetAllowHeadless() bool {
	return c.Spec.AllowHeadless != nil && c.Spec.AllowHeadless.Value
}

func (c *AuthPreferenceV2) SetAllowHeadless(b bool) {
	c.Spec.AllowHeadless = NewBoolOption(b)
}

// SetRequireMFAType sets the type of MFA requirement enforced for this cluster.
func (c *AuthPreferenceV2) SetRequireMFAType(t RequireMFAType) {
	c.Spec.RequireMFAType = t
}

// GetRequireMFAType returns the type of MFA requirement enforced for this cluster.
func (c *AuthPreferenceV2) GetRequireMFAType() RequireMFAType {
	return c.Spec.RequireMFAType
}

// GetPrivateKeyPolicy returns the configured private key policy for the cluster.
func (c *AuthPreferenceV2) GetPrivateKeyPolicy() keys.PrivateKeyPolicy {
	switch c.Spec.RequireMFAType {
	case RequireMFAType_SESSION_AND_HARDWARE_KEY:
		return keys.PrivateKeyPolicyHardwareKey
	case RequireMFAType_HARDWARE_KEY_TOUCH:
		return keys.PrivateKeyPolicyHardwareKeyTouch
	case RequireMFAType_HARDWARE_KEY_PIN:
		return keys.PrivateKeyPolicyHardwareKeyPIN
	case RequireMFAType_HARDWARE_KEY_TOUCH_AND_PIN:
		return keys.PrivateKeyPolicyHardwareKeyTouchAndPIN
	default:
		return keys.PrivateKeyPolicyNone
	}
}

// GetHardwareKey returns the hardware key settings configured for the cluster.
func (c *AuthPreferenceV2) GetHardwareKey() (*HardwareKey, error) {
	if c.Spec.HardwareKey == nil {
		return nil, trace.NotFound("Hardware key support is not configured in this cluster")
	}
	return c.Spec.HardwareKey, nil
}

// GetPIVSlot returns the configured piv slot for the cluster.
func (c *AuthPreferenceV2) GetPIVSlot() keys.PIVSlot {
	if hk, err := c.GetHardwareKey(); err == nil {
		return keys.PIVSlot(hk.PIVSlot)
	}
	return ""
}

// GetHardwareKeySerialNumberValidation returns the cluster's hardware key
// serial number validation settings.
func (c *AuthPreferenceV2) GetHardwareKeySerialNumberValidation() (*HardwareKeySerialNumberValidation, error) {
	if c.Spec.HardwareKey == nil || c.Spec.HardwareKey.SerialNumberValidation == nil {
		return nil, trace.NotFound("Hardware key serial number validation is not configured in this cluster")
	}
	return c.Spec.HardwareKey.SerialNumberValidation, nil
}

// GetDisconnectExpiredCert returns disconnect expired certificate setting
func (c *AuthPreferenceV2) GetDisconnectExpiredCert() bool {
	return c.Spec.DisconnectExpiredCert.Value
}

// SetDisconnectExpiredCert sets disconnect client with expired certificate setting
func (c *AuthPreferenceV2) SetDisconnectExpiredCert(b bool) {
	c.Spec.DisconnectExpiredCert = NewBoolOption(b)
}

// GetAllowLocalAuth gets if local authentication is allowed.
func (c *AuthPreferenceV2) GetAllowLocalAuth() bool {
	return c.Spec.AllowLocalAuth.Value
}

// SetAllowLocalAuth gets if local authentication is allowed.
func (c *AuthPreferenceV2) SetAllowLocalAuth(b bool) {
	c.Spec.AllowLocalAuth = NewBoolOption(b)
}

// GetMessageOfTheDay gets the current Message Of The Day. May be empty.
func (c *AuthPreferenceV2) GetMessageOfTheDay() string {
	return c.Spec.MessageOfTheDay
}

// SetMessageOfTheDay sets the current Message Of The Day. May be empty.
func (c *AuthPreferenceV2) SetMessageOfTheDay(motd string) {
	c.Spec.MessageOfTheDay = motd
}

// GetLockingMode gets the cluster-wide locking mode default.
func (c *AuthPreferenceV2) GetLockingMode() constants.LockingMode {
	return c.Spec.LockingMode
}

// SetLockingMode sets the cluster-wide locking mode default.
func (c *AuthPreferenceV2) SetLockingMode(mode constants.LockingMode) {
	c.Spec.LockingMode = mode
}

// GetDeviceTrust returns the cluster device trust settings, or nil if no
// explicit configurations are present.
func (c *AuthPreferenceV2) GetDeviceTrust() *DeviceTrust {
	if c == nil {
		return nil
	}
	return c.Spec.DeviceTrust
}

// SetDeviceTrust sets the cluster device trust settings.
func (c *AuthPreferenceV2) SetDeviceTrust(dt *DeviceTrust) {
	c.Spec.DeviceTrust = dt
}

// IsSAMLIdPEnabled returns true if the SAML IdP is enabled.
func (c *AuthPreferenceV2) IsSAMLIdPEnabled() bool {
	return c.Spec.IDP.SAML.Enabled.Value
}

// SetSAMLIdPEnabled sets the SAML IdP to enabled.
func (c *AuthPreferenceV2) SetSAMLIdPEnabled(enabled bool) {
	c.Spec.IDP.SAML.Enabled = NewBoolOption(enabled)
}

// SetDefaultSessionTTL sets the default session ttl
func (c *AuthPreferenceV2) SetDefaultSessionTTL(sessionTTL Duration) {
	c.Spec.DefaultSessionTTL = sessionTTL
}

// GetDefaultSessionTTL retrieves the default session ttl
func (c *AuthPreferenceV2) GetDefaultSessionTTL() Duration {
	return c.Spec.DefaultSessionTTL
}

// GetOktaSyncPeriod returns the duration between Okta synchronization calls if the Okta service is running.
func (c *AuthPreferenceV2) GetOktaSyncPeriod() time.Duration {
	return c.Spec.Okta.SyncPeriod.Duration()
}

// SetOktaSyncPeriod sets the duration between Okta synchronzation calls.
func (c *AuthPreferenceV2) SetOktaSyncPeriod(syncPeriod time.Duration) {
	c.Spec.Okta.SyncPeriod = Duration(syncPeriod)
}

// setStaticFields sets static resource header and metadata fields.
func (c *AuthPreferenceV2) setStaticFields() {
	c.Kind = KindClusterAuthPreference
	c.Version = V2
	c.Metadata.Name = MetaNameClusterAuthPreference
}

// GetSignatureAlgorithmSuite gets the signature algorithm suite.
func (c *AuthPreferenceV2) GetSignatureAlgorithmSuite() SignatureAlgorithmSuite {
	return c.Spec.SignatureAlgorithmSuite
}

// SetSignatureAlgorithmSuite sets the signature algorithm suite.
func (c *AuthPreferenceV2) SetSignatureAlgorithmSuite(suite SignatureAlgorithmSuite) {
	c.Spec.SignatureAlgorithmSuite = suite
}

// SignatureAlgorithmSuiteParams is a set of parameters used to determine if a
// configured signature algorithm suite is valid, or to set a default signature
// algorithm suite.
type SignatureAlgorithmSuiteParams struct {
	// FIPS should be true if running in FIPS mode.
	FIPS bool
	// UsingHSMOrKMS should be true if the auth server is configured to
	// use an HSM or KMS.
	UsingHSMOrKMS bool
}

// SetDefaultSignatureAlgorithmSuite sets default signature algorithm suite
// based on the params. This is meant for a default auth preference in a
// brand new cluster or after resetting the auth preference.
func (c *AuthPreferenceV2) SetDefaultSignatureAlgorithmSuite(params SignatureAlgorithmSuiteParams) {
	switch {
	case params.FIPS:
		c.SetSignatureAlgorithmSuite(SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_FIPS_V1)
	case params.UsingHSMOrKMS:
		c.SetSignatureAlgorithmSuite(SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_HSM_V1)
	default:
		c.SetSignatureAlgorithmSuite(SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_BALANCED_V1)
	}
}

var (
	errNonFIPSSignatureAlgorithmSuite = &trace.BadParameterError{Message: `non-FIPS compliant authentication setting: "signature_algorithm_suite" must be "fips-v1" or "legacy"`}
	errNonHSMSignatureAlgorithmSuite  = &trace.BadParameterError{Message: `configured "signature_algorithm_suite" is unsupported when "ca_key_params" configures an HSM or KMS, supported values: ["hsm-v1", "fips-v1", "legacy"]`}
)

// CheckSignatureAlgorithmSuite returns an error if the current signature
// algorithm suite is incompatible with [params].
func (c *AuthPreferenceV2) CheckSignatureAlgorithmSuite(params SignatureAlgorithmSuiteParams) error {
	switch c.GetSignatureAlgorithmSuite() {
	case SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_UNSPECIFIED,
		SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_LEGACY,
		SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_FIPS_V1:
		// legacy, fips-v1, and unspecified are always valid.
	case SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_HSM_V1:
		if params.FIPS {
			return trace.Wrap(errNonFIPSSignatureAlgorithmSuite)
		}
	case SignatureAlgorithmSuite_SIGNATURE_ALGORITHM_SUITE_BALANCED_V1:
		if params.FIPS {
			return trace.Wrap(errNonFIPSSignatureAlgorithmSuite)
		}
		if params.UsingHSMOrKMS {
			return trace.Wrap(errNonHSMSignatureAlgorithmSuite)
		}
	default:
		return trace.Errorf("unhandled signature_algorithm_suite %q: this is a bug", c.GetSignatureAlgorithmSuite())
	}
	return nil
}

// CheckAndSetDefaults verifies the constraints for AuthPreference.
func (c *AuthPreferenceV2) CheckAndSetDefaults() error {
	c.setStaticFields()
	if err := c.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}

	if c.Spec.Type == "" {
		c.Spec.Type = constants.Local
	}
	if c.Spec.AllowLocalAuth == nil {
		c.Spec.AllowLocalAuth = NewBoolOption(true)
	}
	if c.Spec.DisconnectExpiredCert == nil {
		c.Spec.DisconnectExpiredCert = NewBoolOption(false)
	}
	if c.Spec.LockingMode == "" {
		c.Spec.LockingMode = constants.LockingModeBestEffort
	}
	if c.Origin() == "" {
		c.SetOrigin(OriginDynamic)
	}

	if c.Spec.DefaultSessionTTL == 0 {
		c.Spec.DefaultSessionTTL = Duration(defaults.CertDuration)
	}

	switch c.Spec.Type {
	case constants.Local, constants.OIDC, constants.SAML, constants.Github:
		// Note that "type:local" and "local_auth:false" is considered a valid
		// setting, as it is a common idiom for clusters that rely on dynamic
		// configuration.
	default:
		return trace.BadParameter("authentication type %q not supported", c.Spec.Type)
	}

	if c.Spec.SecondFactor == constants.SecondFactorU2F {
		const deprecationMessage = `` +
			`Second Factor "u2f" is deprecated and marked for removal, using "webauthn" instead. ` +
			`Please update your configuration to use WebAuthn. ` +
			`Refer to https://goteleport.com/docs/access-controls/guides/webauthn/`
		slog.WarnContext(context.Background(), deprecationMessage)
		c.Spec.SecondFactor = constants.SecondFactorWebauthn
	}

	// If U2F is present validate it, we can derive Webauthn from it.
	if c.Spec.U2F != nil {
		if err := c.Spec.U2F.Check(); err != nil {
			return trace.Wrap(err)
		}
		if c.Spec.Webauthn == nil {
			// Not a problem, try to derive from U2F.
			c.Spec.Webauthn = &Webauthn{}
		}
		if err := c.Spec.Webauthn.CheckAndSetDefaults(c.Spec.U2F); err != nil {
			return trace.Wrap(err)
		}
	}

	// Set SecondFactors from SecondFactor.
	if len(c.Spec.SecondFactors) == 0 {
		c.Spec.SecondFactors = SecondFactorsFromLegacySecondFactor(c.Spec.SecondFactor, c.Spec.Webauthn != nil)
	}

	// Validate expected fields for webauthn.
	hasWebauthn := c.IsSecondFactorWebauthnAllowed()
	if hasWebauthn {
		if c.Spec.Webauthn == nil {
			return trace.BadParameter("missing required webauthn configuration")
		}

		if err := c.Spec.Webauthn.CheckAndSetDefaults(c.Spec.U2F); err != nil {
			return trace.Wrap(err)
		}
	}

	// Set/validate AllowPasswordless. We need Webauthn first to do this properly.
	switch {
	case c.Spec.AllowPasswordless == nil:
		c.Spec.AllowPasswordless = NewBoolOption(hasWebauthn)
	case !hasWebauthn && c.Spec.AllowPasswordless.Value:
		return trace.BadParameter("missing required webauthn configuration for passwordless=true")
	}

	// Set/validate AllowHeadless. We need Webauthn first to do this properly.
	switch {
	case c.Spec.AllowHeadless == nil:
		c.Spec.AllowHeadless = NewBoolOption(hasWebauthn)
	case !hasWebauthn && c.Spec.AllowHeadless.Value:
		return trace.BadParameter("missing required webauthn configuration for headless=true")
	}

	// Validate connector name for type=local.
	if c.Spec.Type == constants.Local {
		switch connectorName := c.Spec.ConnectorName; connectorName {
		case "", constants.LocalConnector: // OK
		case constants.PasswordlessConnector:
			if !c.Spec.AllowPasswordless.Value {
				return trace.BadParameter("invalid local connector %q, passwordless not allowed by cluster settings", connectorName)
			}
		case constants.HeadlessConnector:
			if !c.Spec.AllowHeadless.Value {
				return trace.BadParameter("invalid local connector %q, headless not allowed by cluster settings", connectorName)
			}
		default:
			return trace.BadParameter("invalid local connector %q", connectorName)
		}
	}

	switch c.Spec.LockingMode {
	case constants.LockingModeBestEffort, constants.LockingModeStrict:
	default:
		return trace.BadParameter("locking mode %q not supported", c.Spec.LockingMode)
	}

	if dt := c.Spec.DeviceTrust; dt != nil {
		switch dt.Mode {
		case "": // OK, "default" mode. Varies depending on OSS or Enterprise.
		case constants.DeviceTrustModeOff,
			constants.DeviceTrustModeOptional,
			constants.DeviceTrustModeRequired: // OK.
		default:
			return trace.BadParameter("device trust mode %q not supported", dt.Mode)
		}

		// Ensure configured ekcert_allowed_cas are valid
		for _, pem := range dt.EKCertAllowedCAs {
			if err := isValidCertificatePEM(pem); err != nil {
				return trace.BadParameter("device trust has invalid EKCert allowed CAs entry: %v", err)
			}
		}
	}

	// TODO(Joerger): DELETE IN 17.0.0
	c.CheckSetPIVSlot()

	if hk, err := c.GetHardwareKey(); err == nil && hk.PIVSlot != "" {
		if err := keys.PIVSlot(hk.PIVSlot).Validate(); err != nil {
			return trace.Wrap(err)
		}
	}

	// Make sure the IdP section is populated.
	if c.Spec.IDP == nil {
		c.Spec.IDP = &IdPOptions{}
	}

	// Make sure the SAML section is populated.
	if c.Spec.IDP.SAML == nil {
		c.Spec.IDP.SAML = &IdPSAMLOptions{}
	}

	// Make sure the SAML enabled field is populated.
	if c.Spec.IDP.SAML.Enabled == nil {
		// Enable the IdP by default.
		c.Spec.IDP.SAML.Enabled = NewBoolOption(true)
	}

	// Make sure the Okta field is populated.
	if c.Spec.Okta == nil {
		c.Spec.Okta = &OktaOptions{}
	}

	return nil
}

// CheckSetPIVSlot ensures that the PIVSlot and Hardwarekey.PIVSlot stay in sync so that
// older versions of Teleport that do not know about Hardwarekey.PIVSlot are able to keep
// using PIVSlot and newer versions of Teleport can rely solely on Hardwarekey.PIVSlot
// without causing any service degradation.
// TODO(Joerger): DELETE IN 17.0.0
func (c *AuthPreferenceV2) CheckSetPIVSlot() {
	if c.Spec.PIVSlot != "" {
		if c.Spec.HardwareKey == nil {
			c.Spec.HardwareKey = &HardwareKey{}
		}
		c.Spec.HardwareKey.PIVSlot = c.Spec.PIVSlot
	} else if c.Spec.HardwareKey != nil && c.Spec.HardwareKey.PIVSlot != "" {
		c.Spec.PIVSlot = c.Spec.HardwareKey.PIVSlot
	}
}

// String represents a human readable version of authentication settings.
func (c *AuthPreferenceV2) String() string {
	return fmt.Sprintf("AuthPreference(Type=%q,SecondFactors=%q)", c.Spec.Type, c.GetSecondFactors())
}

// Clone returns a copy of the AuthPreference resource.
func (c *AuthPreferenceV2) Clone() AuthPreference {
	return utils.CloneProtoMsg(c)
}

func (u *U2F) Check() error {
	if u.AppID == "" {
		return trace.BadParameter("u2f configuration missing app_id")
	}
	for _, ca := range u.DeviceAttestationCAs {
		if err := isValidCertificatePEM(ca); err != nil {
			return trace.BadParameter("u2f configuration has an invalid attestation CA: %v", err)
		}
	}
	return nil
}

func (w *Webauthn) CheckAndSetDefaults(u *U2F) error {
	// RPID.
	switch {
	case w.RPID != "": // Explicit RPID
		_, err := url.Parse(w.RPID)
		if err != nil {
			return trace.BadParameter("webauthn rp_id is not a valid URI: %v", err)
		}
	case u != nil && w.RPID == "": // Infer RPID from U2F app_id
		parsedAppID, err := url.Parse(u.AppID)
		if err != nil {
			return trace.BadParameter("webauthn missing rp_id and U2F app_id is not an URL (%v)", err)
		}

		var rpID string
		switch {
		case parsedAppID.Host != "":
			rpID = parsedAppID.Host
			rpID = strings.Split(rpID, ":")[0] // Remove :port, if present
		case parsedAppID.Path == u.AppID:
			// App ID is not a proper URL, take it literally.
			rpID = u.AppID
		default:
			return trace.BadParameter("failed to infer webauthn RPID from U2F App ID (%q)", u.AppID)
		}
		slog.InfoContext(context.Background(), "WebAuthn: RPID inferred from U2F configuration", "rpid", rpID)
		w.RPID = rpID
	default:
		return trace.BadParameter("webauthn configuration missing rp_id")
	}

	// AttestationAllowedCAs.
	switch {
	case u != nil && len(u.DeviceAttestationCAs) > 0 && len(w.AttestationAllowedCAs) == 0 && len(w.AttestationDeniedCAs) == 0:
		slog.InfoContext(context.Background(), "WebAuthn: using U2F device attestation CAs as allowed CAs")
		w.AttestationAllowedCAs = u.DeviceAttestationCAs
	default:
		for _, pem := range w.AttestationAllowedCAs {
			if err := isValidCertificatePEM(pem); err != nil {
				return trace.BadParameter("webauthn allowed CAs entry invalid: %v", err)
			}
		}
	}

	// AttestationDeniedCAs.
	for _, pem := range w.AttestationDeniedCAs {
		if err := isValidCertificatePEM(pem); err != nil {
			return trace.BadParameter("webauthn denied CAs entry invalid: %v", err)
		}
	}

	return nil
}

func isValidCertificatePEM(pem string) error {
	_, err := tlsutils.ParseCertificatePEM([]byte(pem))
	return err
}

// Check validates WebauthnLocalAuth, returning an error if it's not valid.
func (wal *WebauthnLocalAuth) Check() error {
	if len(wal.UserID) == 0 {
		return trace.BadParameter("missing UserID field")
	}
	return nil
}

// NewMFADevice creates a new MFADevice with the given name. Caller must set
// the Device field in the returned MFADevice.
func NewMFADevice(name, id string, addedAt time.Time) *MFADevice {
	return &MFADevice{
		Metadata: Metadata{
			Name: name,
		},
		Id:       id,
		AddedAt:  addedAt,
		LastUsed: addedAt,
	}
}

// setStaticFields sets static resource header and metadata fields.
func (d *MFADevice) setStaticFields() {
	d.Kind = KindMFADevice
	d.Version = V1
}

// CheckAndSetDefaults validates MFADevice fields and populates empty fields
// with default values.
func (d *MFADevice) CheckAndSetDefaults() error {
	d.setStaticFields()
	if err := d.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}
	if d.Id == "" {
		return trace.BadParameter("MFADevice missing ID field")
	}
	if d.AddedAt.IsZero() {
		return trace.BadParameter("MFADevice missing AddedAt field")
	}
	if d.LastUsed.IsZero() {
		return trace.BadParameter("MFADevice missing LastUsed field")
	}
	if d.LastUsed.Before(d.AddedAt) {
		return trace.BadParameter("MFADevice LastUsed field must be earlier than AddedAt")
	}
	if d.Device == nil {
		return trace.BadParameter("MFADevice missing Device field")
	}
	if err := checkWebauthnDevice(d); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func checkWebauthnDevice(d *MFADevice) error {
	wrapper, ok := d.Device.(*MFADevice_Webauthn)
	if !ok {
		return nil
	}
	switch webDev := wrapper.Webauthn; {
	case webDev == nil:
		return trace.BadParameter("MFADevice has malformed WebauthnDevice")
	case len(webDev.CredentialId) == 0:
		return trace.BadParameter("WebauthnDevice missing CredentialId field")
	case len(webDev.PublicKeyCbor) == 0:
		return trace.BadParameter("WebauthnDevice missing PublicKeyCbor field")
	default:
		return nil
	}
}

func (d *MFADevice) GetKind() string         { return d.Kind }
func (d *MFADevice) GetSubKind() string      { return d.SubKind }
func (d *MFADevice) SetSubKind(sk string)    { d.SubKind = sk }
func (d *MFADevice) GetVersion() string      { return d.Version }
func (d *MFADevice) GetMetadata() Metadata   { return d.Metadata }
func (d *MFADevice) GetName() string         { return d.Metadata.GetName() }
func (d *MFADevice) SetName(n string)        { d.Metadata.SetName(n) }
func (d *MFADevice) GetRevision() string     { return d.Metadata.GetRevision() }
func (d *MFADevice) SetRevision(rev string)  { d.Metadata.SetRevision(rev) }
func (d *MFADevice) Expiry() time.Time       { return d.Metadata.Expiry() }
func (d *MFADevice) SetExpiry(exp time.Time) { d.Metadata.SetExpiry(exp) }

// MFAType returns the human-readable name of the MFA protocol of this device.
func (d *MFADevice) MFAType() string {
	switch d.Device.(type) {
	case *MFADevice_Totp:
		return "TOTP"
	case *MFADevice_U2F:
		return "U2F"
	case *MFADevice_Webauthn:
		return "WebAuthn"
	default:
		return "unknown"
	}
}

func (d *MFADevice) MarshalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := (&jsonpb.Marshaler{}).Marshal(buf, d)
	return buf.Bytes(), trace.Wrap(err)
}

func (d *MFADevice) UnmarshalJSON(buf []byte) error {
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	err := unmarshaler.Unmarshal(bytes.NewReader(buf), d)
	return trace.Wrap(err)
}

// IsSessionMFARequired returns whether this RequireMFAType requires per-session MFA.
func (r RequireMFAType) IsSessionMFARequired() bool {
	return r != RequireMFAType_OFF
}

// MarshalJSON marshals RequireMFAType to boolean or string.
func (r *RequireMFAType) MarshalYAML() (interface{}, error) {
	val, err := r.encode()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return val, nil
}

// UnmarshalYAML supports parsing RequireMFAType from boolean or alias.
func (r *RequireMFAType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val interface{}
	err := unmarshal(&val)
	if err != nil {
		return trace.Wrap(err)
	}

	err = r.decode(val)
	return trace.Wrap(err)
}

// MarshalJSON marshals RequireMFAType to boolean or string.
func (r *RequireMFAType) MarshalJSON() ([]byte, error) {
	val, err := r.encode()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	out, err := json.Marshal(val)
	return out, trace.Wrap(err)
}

// UnmarshalJSON supports parsing RequireMFAType from boolean or alias.
func (r *RequireMFAType) UnmarshalJSON(data []byte) error {
	var val interface{}
	err := json.Unmarshal(data, &val)
	if err != nil {
		return trace.Wrap(err)
	}

	err = r.decode(val)
	return trace.Wrap(err)
}

const (
	// RequireMFATypeHardwareKeyString is the string representation of RequireMFATypeHardwareKey
	RequireMFATypeHardwareKeyString = "hardware_key"
	// RequireMFATypeHardwareKeyTouchString is the string representation of RequireMFATypeHardwareKeyTouch
	RequireMFATypeHardwareKeyTouchString = "hardware_key_touch"
	// RequireMFATypeHardwareKeyPINString is the string representation of RequireMFATypeHardwareKeyPIN
	RequireMFATypeHardwareKeyPINString = "hardware_key_pin"
	// RequireMFATypeHardwareKeyTouchAndPINString is the string representation of RequireMFATypeHardwareKeyTouchAndPIN
	RequireMFATypeHardwareKeyTouchAndPINString = "hardware_key_touch_and_pin"
)

// encode RequireMFAType into a string or boolean. This is necessary for
// backwards compatibility with the json/yaml tag "require_session_mfa",
// which used to be a boolean.
func (r *RequireMFAType) encode() (interface{}, error) {
	switch *r {
	case RequireMFAType_OFF:
		return false, nil
	case RequireMFAType_SESSION:
		return true, nil
	case RequireMFAType_SESSION_AND_HARDWARE_KEY:
		return RequireMFATypeHardwareKeyString, nil
	case RequireMFAType_HARDWARE_KEY_TOUCH:
		return RequireMFATypeHardwareKeyTouchString, nil
	case RequireMFAType_HARDWARE_KEY_PIN:
		return RequireMFATypeHardwareKeyPINString, nil
	case RequireMFAType_HARDWARE_KEY_TOUCH_AND_PIN:
		return RequireMFATypeHardwareKeyTouchAndPINString, nil
	default:
		return nil, trace.BadParameter("RequireMFAType invalid value %v", *r)
	}
}

// decode RequireMFAType from a string or boolean. This is necessary for
// backwards compatibility with the json/yaml tag "require_session_mfa",
// which used to be a boolean.
func (r *RequireMFAType) decode(val any) error {
	err := decodeEnum(r, val, map[any]RequireMFAType{
		"":                                   RequireMFAType_OFF, // default to off
		false:                                RequireMFAType_OFF,
		true:                                 RequireMFAType_SESSION,
		RequireMFATypeHardwareKeyString:      RequireMFAType_SESSION_AND_HARDWARE_KEY,
		RequireMFATypeHardwareKeyTouchString: RequireMFAType_HARDWARE_KEY_TOUCH,
		RequireMFATypeHardwareKeyPINString:   RequireMFAType_HARDWARE_KEY_PIN,
		RequireMFATypeHardwareKeyTouchAndPINString: RequireMFAType_HARDWARE_KEY_TOUCH_AND_PIN,
	}, RequireMFAType_name)
	return trace.Wrap(err, "failed to decode require mfa type")
}

// MarshalJSON marshals SecondFactorType to boolean or string.
func (s *SecondFactorType) MarshalYAML() (interface{}, error) {
	val, err := s.encode()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return val, nil
}

// UnmarshalYAML supports parsing SecondFactorType from boolean or alias.
func (s *SecondFactorType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val interface{}
	err := unmarshal(&val)
	if err != nil {
		return trace.Wrap(err)
	}

	err = s.decode(val)
	return trace.Wrap(err)
}

// MarshalJSON marshals SecondFactorType to boolean or string.
func (s *SecondFactorType) MarshalJSON() ([]byte, error) {
	val, err := s.encode()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	out, err := json.Marshal(val)
	return out, trace.Wrap(err)
}

// UnmarshalJSON supports parsing SecondFactorType from boolean or alias.
func (s *SecondFactorType) UnmarshalJSON(data []byte) error {
	var val interface{}
	err := json.Unmarshal(data, &val)
	if err != nil {
		return trace.Wrap(err)
	}

	err = s.decode(val)
	return trace.Wrap(err)
}

const (
	// SecondFactorTypeOTPString is the string representation of SecondFactorType_SECOND_FACTOR_TYPE_OTP
	SecondFactorTypeOTPString = "otp"
	// SecondFactorTypeWebauthnString is the string representation of SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN
	SecondFactorTypeWebauthnString = "webauthn"
	// SecondFactorTypeSSOString is the string representation of SecondFactorType_SECOND_FACTOR_TYPE_SSO
	SecondFactorTypeSSOString = "sso"
)

// encode SecondFactorType into a string or boolean. This is necessary for
// backwards compatibility with the json/yaml tag "require_session_mfa",
// which used to be a boolean.
func (s *SecondFactorType) encode() (string, error) {
	switch *s {
	case SecondFactorType_SECOND_FACTOR_TYPE_UNSPECIFIED:
		return "", nil
	case SecondFactorType_SECOND_FACTOR_TYPE_OTP:
		return SecondFactorTypeOTPString, nil
	case SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN:
		return SecondFactorTypeWebauthnString, nil
	case SecondFactorType_SECOND_FACTOR_TYPE_SSO:
		return SecondFactorTypeSSOString, nil
	default:
		return "", trace.BadParameter("SecondFactorType invalid value %v", *s)
	}
}

// decode SecondFactorType from a string or boolean. This is necessary for
// backwards compatibility with the json/yaml tag "require_session_mfa",
// which used to be a boolean.
func (s *SecondFactorType) decode(val any) error {
	err := decodeEnum(s, val, map[any]SecondFactorType{
		"":                             SecondFactorType_SECOND_FACTOR_TYPE_UNSPECIFIED,
		SecondFactorTypeOTPString:      SecondFactorType_SECOND_FACTOR_TYPE_OTP,
		SecondFactorTypeWebauthnString: SecondFactorType_SECOND_FACTOR_TYPE_WEBAUTHN,
		SecondFactorTypeSSOString:      SecondFactorType_SECOND_FACTOR_TYPE_SSO,
	}, SecondFactorType_name)
	return trace.Wrap(err, "failed to decode SecondFactorType")
}
