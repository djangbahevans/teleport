/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { render, screen } from 'design/utils/testing';
import React from 'react';

import { waitFor, within } from '@testing-library/react';
import { userEvent, UserEvent } from '@testing-library/user-event';

import TeleportContext from 'teleport/teleportContext';
import { ContextProvider } from 'teleport';
import MfaService from 'teleport/services/mfa';
import auth from 'teleport/services/auth';

import { AddAuthDeviceWizardStepProps } from './AddAuthDeviceWizard';

import { AddAuthDeviceWizard } from '.';

const dummyCredential: Credential = { id: 'cred-id', type: 'public-key' };
let ctx: TeleportContext;
let user: UserEvent;
let onSuccess: jest.Mock;

beforeEach(() => {
  ctx = new TeleportContext();
  user = userEvent.setup();
  onSuccess = jest.fn();

  jest
    .spyOn(auth, 'createNewWebAuthnDevice')
    .mockResolvedValueOnce(dummyCredential);
  jest
    .spyOn(MfaService.prototype, 'saveNewWebAuthnDevice')
    .mockResolvedValueOnce(undefined);
  jest.spyOn(auth, 'createMfaRegistrationChallenge').mockResolvedValueOnce({
    qrCode: 'dummy-qr-code',
    webauthnPublicKey: {} as PublicKeyCredentialCreationOptions,
  });
  jest
    .spyOn(MfaService.prototype, 'addNewTotpDevice')
    .mockResolvedValueOnce(undefined);
});

afterEach(jest.resetAllMocks);

function TestWizard(props: Partial<AddAuthDeviceWizardStepProps> = {}) {
  return (
    <ContextProvider ctx={ctx}>
      <AddAuthDeviceWizard
        usage="passwordless"
        auth2faType="on"
        onClose={() => {}}
        onSuccess={onSuccess}
        {...props}
      />
    </ContextProvider>
  );
}

describe('flow without reauthentication', () => {
  beforeEach(() => {
    jest.spyOn(auth, 'getMfaChallenge').mockResolvedValueOnce({});
  });

  test('adds a passkey', async () => {
    render(
      <TestWizard usage="passwordless" privilegeToken="privilege-token" />
    );

    const createStep = await waitFor(() => {
      return within(screen.getByTestId('create-step'));
    });
    await user.click(
      createStep.getByRole('button', { name: 'Create a passkey' })
    );
    expect(auth.createNewWebAuthnDevice).toHaveBeenCalledWith({
      tokenId: 'privilege-token',
      deviceUsage: 'passwordless',
    });

    const saveStep = within(screen.getByTestId('save-step'));
    await user.type(saveStep.getByLabelText('Passkey Nickname'), 'new-passkey');
    await user.click(
      saveStep.getByRole('button', { name: 'Save the Passkey' })
    );
    expect(ctx.mfaService.saveNewWebAuthnDevice).toHaveBeenCalledWith({
      credential: dummyCredential,
      addRequest: {
        deviceName: 'new-passkey',
        deviceUsage: 'passwordless',
        tokenId: 'privilege-token',
      },
    });
    expect(onSuccess).toHaveBeenCalled();
  });

  test('adds a WebAuthn MFA', async () => {
    render(<TestWizard usage="mfa" privilegeToken="privilege-token" />);

    const createStep = await waitFor(() => {
      return within(screen.getByTestId('create-step'));
    });
    await user.click(createStep.getByLabelText('Hardware Device'));
    await user.click(
      createStep.getByRole('button', { name: 'Create an MFA method' })
    );
    expect(auth.createNewWebAuthnDevice).toHaveBeenCalledWith({
      tokenId: 'privilege-token',
      deviceUsage: 'mfa',
    });

    const saveStep = within(screen.getByTestId('save-step'));
    await user.type(saveStep.getByLabelText('MFA Method Name'), 'new-mfa');
    await user.click(
      saveStep.getByRole('button', { name: 'Save the MFA method' })
    );
    expect(ctx.mfaService.saveNewWebAuthnDevice).toHaveBeenCalledWith({
      credential: dummyCredential,
      addRequest: {
        deviceName: 'new-mfa',
        deviceUsage: 'mfa',
        tokenId: 'privilege-token',
      },
    });
    expect(onSuccess).toHaveBeenCalled();
  });

  test('adds an authenticator app', async () => {
    render(<TestWizard usage="mfa" privilegeToken="privilege-token" />);

    const createStep = await waitFor(() => {
      return within(screen.getByTestId('create-step'));
    });

    await user.click(createStep.getByLabelText('Authenticator App'));
    expect(createStep.getByRole('img')).toHaveAttribute(
      'src',
      'data:image/png;base64,dummy-qr-code'
    );
    await user.click(
      createStep.getByRole('button', { name: 'Create an MFA method' })
    );

    const saveStep = within(screen.getByTestId('save-step'));
    await user.type(saveStep.getByLabelText('MFA Method Name'), 'new-mfa');
    await user.type(saveStep.getByLabelText(/Authenticator Code/), '345678');
    await user.click(
      saveStep.getByRole('button', { name: 'Save the MFA method' })
    );
    expect(ctx.mfaService.addNewTotpDevice).toHaveBeenCalledWith({
      tokenId: 'privilege-token',
      secondFactorToken: '345678',
      deviceName: 'new-mfa',
    });
    expect(onSuccess).toHaveBeenCalled();
  });
});

describe('flow with reauthentication', () => {
  beforeEach(() => {
    jest.spyOn(auth, 'getMfaChallenge').mockResolvedValueOnce({
      totpChallenge: true,
      webauthnPublicKey: {} as PublicKeyCredentialRequestOptions,
    });
    jest.spyOn(auth, 'getMfaChallengeResponse').mockResolvedValueOnce({});
    jest
      .spyOn(auth, 'createPrivilegeToken')
      .mockResolvedValueOnce('privilege-token');
  });

  test('adds a passkey with WebAuthn reauthentication', async () => {
    render(<TestWizard usage="passwordless" />);

    const reauthenticateStep = await waitFor(() => {
      return within(screen.getByTestId('reauthenticate-step'));
    });

    await user.click(reauthenticateStep.getByText('Verify my identity'));

    const createStep = await waitFor(() => {
      return within(screen.getByTestId('create-step'));
    });
    await user.click(
      createStep.getByRole('button', { name: 'Create a passkey' })
    );
    expect(auth.createNewWebAuthnDevice).toHaveBeenCalledWith({
      tokenId: 'privilege-token',
      deviceUsage: 'passwordless',
    });

    const saveStep = within(screen.getByTestId('save-step'));
    await user.type(saveStep.getByLabelText('Passkey Nickname'), 'new-passkey');
    await user.click(
      saveStep.getByRole('button', { name: 'Save the Passkey' })
    );
    expect(ctx.mfaService.saveNewWebAuthnDevice).toHaveBeenCalledWith({
      credential: dummyCredential,
      addRequest: {
        deviceName: 'new-passkey',
        deviceUsage: 'passwordless',
        tokenId: 'privilege-token',
      },
    });
    expect(onSuccess).toHaveBeenCalled();
  });

  test('adds a passkey with OTP reauthentication', async () => {
    render(<TestWizard usage="passwordless" />);

    const reauthenticateStep = await waitFor(() => {
      return within(screen.getByTestId('reauthenticate-step'));
    });

    await user.click(reauthenticateStep.getByText('Authenticator App'));
    await user.type(
      reauthenticateStep.getByLabelText('Authenticator Code'),
      '654987'
    );
    await user.click(reauthenticateStep.getByText('Verify my identity'));

    const createStep = await waitFor(() => {
      return within(screen.getByTestId('create-step'));
    });
    await user.click(
      createStep.getByRole('button', { name: 'Create a passkey' })
    );
    expect(auth.createNewWebAuthnDevice).toHaveBeenCalledWith({
      tokenId: 'privilege-token',
      deviceUsage: 'passwordless',
    });

    const saveStep = within(screen.getByTestId('save-step'));
    await user.type(saveStep.getByLabelText('Passkey Nickname'), 'new-passkey');
    await user.click(
      saveStep.getByRole('button', { name: 'Save the Passkey' })
    );
    expect(ctx.mfaService.saveNewWebAuthnDevice).toHaveBeenCalledWith({
      credential: dummyCredential,
      addRequest: {
        deviceName: 'new-passkey',
        deviceUsage: 'passwordless',
        tokenId: 'privilege-token',
      },
    });
    expect(onSuccess).toHaveBeenCalled();
  });

  test('shows reauthentication options', async () => {
    render(<TestWizard usage="mfa" />);

    const reauthenticateStep = await waitFor(() => {
      return within(screen.getByTestId('reauthenticate-step'));
    });

    expect(
      reauthenticateStep.queryByLabelText(/passkey or security key/i)
    ).toBeVisible();
    expect(
      reauthenticateStep.queryByLabelText(/authenticator app/i)
    ).toBeVisible();
  });
});
