/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

import { useCallback, useEffect, useState } from 'react';
import { Attempt, makeEmptyAttempt, useAsync } from 'shared/hooks/useAsync';

import auth from 'teleport/services/auth';
import { MfaChallengeScope } from 'teleport/services/auth/auth';

import {
  DeviceType,
  DeviceUsage,
  getMfaChallengeOptions,
  MfaAuthenticateChallenge,
  MfaChallengeResponse,
  MfaOption,
} from 'teleport/services/mfa';

export default function useReAuthenticate({
  challengeScope,
  onMfaResponse,
}: ReauthProps): ReauthState {
  const [mfaOptions, setMfaOptions] = useState<MfaOption[]>();
  const [challengeState, setChallengeState] = useState<challengeState>();

  const [initAttempt, init] = useAsync(async () => {
    const challenge = await auth.getMfaChallenge({
      scope: challengeScope,
    });

    setChallengeState({ challenge, deviceUsage: 'mfa' });
    setMfaOptions(getMfaChallengeOptions(challenge));
  });

  useEffect(() => {
    init();
  }, []);

  const getChallenge = useCallback(
    async (deviceUsage: DeviceUsage = 'mfa') => {
      // If the challenge state is empty, used, or has different args,
      // retrieve a new mfa challenge and set it in the state.
      if (!challengeState || challengeState.deviceUsage != deviceUsage) {
        const challenge = await auth.getMfaChallenge({
          scope: challengeScope,
          userVerificationRequirement:
            deviceUsage === 'passwordless' ? 'required' : 'discouraged',
        });
        setChallengeState({
          challenge,
          deviceUsage,
        });
      }
      return challengeState.challenge;
    },
    [challengeScope, challengeState]
  );

  const [submitAttempt, submitWithMfa, setSubmitAttempt] = useAsync(
    useCallback(
      async (
        mfaType?: DeviceType,
        deviceUsage?: DeviceUsage,
        totpCode?: string
      ) => {
        const challenge = await getChallenge(deviceUsage);

        let response: MfaChallengeResponse;
        try {
          response = await auth.getMfaChallengeResponse(
            challenge,
            mfaType,
            totpCode
          );
        } catch (err) {
          throw new Error(getReAuthenticationErrorMessage(err));
        }

        try {
          onMfaResponse(response);
        } finally {
          setChallengeState(null);
        }
      },
      [getChallenge, onMfaResponse]
    )
  );

  function clearSubmitAttempt() {
    setSubmitAttempt(makeEmptyAttempt());
  }

  return {
    initAttempt,
    mfaOptions,
    submitWithMfa,
    submitAttempt,
    clearSubmitAttempt,
  };
}

export type ReauthProps = {
  challengeScope: MfaChallengeScope;
  onMfaResponse(res: MfaChallengeResponse): void;
};

export type ReauthState = {
  initAttempt: Attempt<any>;
  mfaOptions: MfaOption[];
  submitWithMfa: (
    mfaType?: DeviceType,
    deviceUsage?: DeviceUsage,
    totpCode?: string
  ) => Promise<[void, Error]>;
  submitAttempt: Attempt<void>;
  clearSubmitAttempt: () => void;
};

type challengeState = {
  challenge: MfaAuthenticateChallenge;
  deviceUsage: DeviceUsage;
};

function getReAuthenticationErrorMessage(err: Error): string {
  if (err.message.includes('attempt was made to use an object that is not')) {
    // Catch a webauthn frontend error that occurs on Firefox and replace it with a more helpful error message.
    return 'The two-factor device you used is not registered on this account. You must verify using a device that has already been registered.';
  }

  if (err.message === 'invalid totp token') {
    // This message relies on the status message produced by the auth server in
    // lib/auth/Server.checkOTP function. Please keep these in sync.
    return 'Invalid authenticator code';
  }

  return err.message;
}
