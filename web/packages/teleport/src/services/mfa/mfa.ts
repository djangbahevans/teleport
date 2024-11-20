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

import cfg from 'teleport/config';
import api from 'teleport/services/api';
import auth, { makeWebauthnCreationResponse, MfaChallengeResponse } from 'teleport/services/auth';

import {
  MfaDevice,
  AddNewTotpDeviceRequest,
  AddNewHardwareDeviceRequest,
  SaveNewHardwareDeviceRequest,
} from './types';
import makeMfaDevice from './makeMfaDevice';

class MfaService {
  fetchDevicesWithToken(tokenId: string): Promise<MfaDevice[]> {
    return api
      .get(cfg.getMfaDevicesWithTokenUrl(tokenId))
      .then(devices => devices.map(makeMfaDevice));
  }

  removeDeviceWithToken(tokenId: string, deviceName: string) {
    return api.delete(cfg.getMfaDeviceUrl(tokenId, deviceName));
  }

  removeDevice(deviceName: string, existingMfaResponse: MfaChallengeResponse) {
    return api.delete(cfg.api.mfaDevicesPath, {
      deviceName,
      existingMfaResponse,
    });
  }

  fetchDevices(): Promise<MfaDevice[]> {
    return api
      .get(cfg.api.mfaDevicesPath)
      .then(devices => devices.map(makeMfaDevice));
  }

  addNewTotpDevice(req: AddNewTotpDeviceRequest) {
    return api.post(cfg.api.mfaDevicesPath, req);
  }

  saveNewWebAuthnDevice(req: SaveNewHardwareDeviceRequest) {
    return auth.checkWebauthnSupport().then(() => {
      const request = {
        ...req.addRequest,
        webauthnRegisterResponse: makeWebauthnCreationResponse(req.credential),
      };

      return api.post(cfg.api.mfaDevicesPath, request);
    });
  }

  addNewWebauthnDevice(req: AddNewHardwareDeviceRequest) {
    return auth.createNewWebAuthnDevice(req).then(credential => {
      this.saveNewWebAuthnDevice({ addRequest: req, credential });
    });
  }
}

export default MfaService;
