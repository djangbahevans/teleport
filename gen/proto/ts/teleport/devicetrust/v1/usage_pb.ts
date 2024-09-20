/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter eslint_disable,add_pb_suffix,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/devicetrust/v1/usage.proto" (package "teleport.devicetrust.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Superseded by ResourceUsageService.GetUsage.
 *
 * @generated from protobuf message teleport.devicetrust.v1.DevicesUsage
 */
export interface DevicesUsage {
}
/**
 * Superseded by ResourceUsageService.GetUsage.
 *
 * @generated from protobuf enum teleport.devicetrust.v1.AccountUsageType
 */
export enum AccountUsageType {
    /**
     * @generated from protobuf enum value: ACCOUNT_USAGE_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: ACCOUNT_USAGE_TYPE_UNLIMITED = 1;
     */
    UNLIMITED = 1,
    /**
     * @generated from protobuf enum value: ACCOUNT_USAGE_TYPE_USAGE_BASED = 2;
     */
    USAGE_BASED = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class DevicesUsage$Type extends MessageType<DevicesUsage> {
    constructor() {
        super("teleport.devicetrust.v1.DevicesUsage", []);
    }
    create(value?: PartialMessage<DevicesUsage>): DevicesUsage {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DevicesUsage>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DevicesUsage): DevicesUsage {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DevicesUsage, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.devicetrust.v1.DevicesUsage
 */
export const DevicesUsage = new DevicesUsage$Type();