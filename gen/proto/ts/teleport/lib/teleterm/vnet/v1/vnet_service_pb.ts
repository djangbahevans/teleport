/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter long_type_number,eslint_disable,add_pb_suffix,client_grpc1,server_grpc1,ts_nocheck
// @generated from protobuf file "teleport/lib/teleterm/vnet/v1/vnet_service.proto" (package "teleport.lib.teleterm.vnet.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Request for Start.
 *
 * @generated from protobuf message teleport.lib.teleterm.vnet.v1.StartRequest
 */
export interface StartRequest {
    /**
     * @generated from protobuf field: string root_cluster_uri = 1;
     */
    rootClusterUri: string;
}
/**
 * Response for Start.
 *
 * @generated from protobuf message teleport.lib.teleterm.vnet.v1.StartResponse
 */
export interface StartResponse {
}
/**
 * Request for Stop.
 *
 * @generated from protobuf message teleport.lib.teleterm.vnet.v1.StopRequest
 */
export interface StopRequest {
    /**
     * @generated from protobuf field: string root_cluster_uri = 1;
     */
    rootClusterUri: string;
}
/**
 * Response for Stop.
 *
 * @generated from protobuf message teleport.lib.teleterm.vnet.v1.StopResponse
 */
export interface StopResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class StartRequest$Type extends MessageType<StartRequest> {
    constructor() {
        super("teleport.lib.teleterm.vnet.v1.StartRequest", [
            { no: 1, name: "root_cluster_uri", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<StartRequest>): StartRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.rootClusterUri = "";
        if (value !== undefined)
            reflectionMergePartial<StartRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StartRequest): StartRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string root_cluster_uri */ 1:
                    message.rootClusterUri = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: StartRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string root_cluster_uri = 1; */
        if (message.rootClusterUri !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.rootClusterUri);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.vnet.v1.StartRequest
 */
export const StartRequest = new StartRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StartResponse$Type extends MessageType<StartResponse> {
    constructor() {
        super("teleport.lib.teleterm.vnet.v1.StartResponse", []);
    }
    create(value?: PartialMessage<StartResponse>): StartResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<StartResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StartResponse): StartResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: StartResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.vnet.v1.StartResponse
 */
export const StartResponse = new StartResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StopRequest$Type extends MessageType<StopRequest> {
    constructor() {
        super("teleport.lib.teleterm.vnet.v1.StopRequest", [
            { no: 1, name: "root_cluster_uri", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<StopRequest>): StopRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.rootClusterUri = "";
        if (value !== undefined)
            reflectionMergePartial<StopRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StopRequest): StopRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string root_cluster_uri */ 1:
                    message.rootClusterUri = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: StopRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string root_cluster_uri = 1; */
        if (message.rootClusterUri !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.rootClusterUri);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.vnet.v1.StopRequest
 */
export const StopRequest = new StopRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StopResponse$Type extends MessageType<StopResponse> {
    constructor() {
        super("teleport.lib.teleterm.vnet.v1.StopResponse", []);
    }
    create(value?: PartialMessage<StopResponse>): StopResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<StopResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StopResponse): StopResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: StopResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message teleport.lib.teleterm.vnet.v1.StopResponse
 */
export const StopResponse = new StopResponse$Type();
/**
 * @generated ServiceType for protobuf service teleport.lib.teleterm.vnet.v1.VnetService
 */
export const VnetService = new ServiceType("teleport.lib.teleterm.vnet.v1.VnetService", [
    { name: "Start", options: {}, I: StartRequest, O: StartResponse },
    { name: "Stop", options: {}, I: StopRequest, O: StopResponse }
]);
