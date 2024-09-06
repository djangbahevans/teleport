// Copyright 2024 Gravitational, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/autoupdate/v1/autoupdate_service.proto

package autoupdate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for GetAutoUpdateConfig.
type GetAutoUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAutoUpdateConfigRequest) Reset() {
	*x = GetAutoUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAutoUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAutoUpdateConfigRequest) ProtoMessage() {}

func (x *GetAutoUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAutoUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*GetAutoUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{0}
}

// Request for CreateAutoUpdateConfig.
type CreateAutoUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AutoUpdateConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *CreateAutoUpdateConfigRequest) Reset() {
	*x = CreateAutoUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAutoUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAutoUpdateConfigRequest) ProtoMessage() {}

func (x *CreateAutoUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAutoUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateAutoUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAutoUpdateConfigRequest) GetConfig() *AutoUpdateConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for UpdateAutoUpdateConfig.
type UpdateAutoUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AutoUpdateConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdateAutoUpdateConfigRequest) Reset() {
	*x = UpdateAutoUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAutoUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAutoUpdateConfigRequest) ProtoMessage() {}

func (x *UpdateAutoUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAutoUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateAutoUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAutoUpdateConfigRequest) GetConfig() *AutoUpdateConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for UpsertAutoUpdateConfig.
type UpsertAutoUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AutoUpdateConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpsertAutoUpdateConfigRequest) Reset() {
	*x = UpsertAutoUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAutoUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAutoUpdateConfigRequest) ProtoMessage() {}

func (x *UpsertAutoUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAutoUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpsertAutoUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertAutoUpdateConfigRequest) GetConfig() *AutoUpdateConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for DeleteAutoUpdateConfig.
type DeleteAutoUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAutoUpdateConfigRequest) Reset() {
	*x = DeleteAutoUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAutoUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoUpdateConfigRequest) ProtoMessage() {}

func (x *DeleteAutoUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteAutoUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{4}
}

// Request for GetAutoUpdateVersion.
type GetAutoUpdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAutoUpdateVersionRequest) Reset() {
	*x = GetAutoUpdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAutoUpdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAutoUpdateVersionRequest) ProtoMessage() {}

func (x *GetAutoUpdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAutoUpdateVersionRequest.ProtoReflect.Descriptor instead.
func (*GetAutoUpdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{5}
}

// Request for CreateAutoUpdateVersion.
type CreateAutoUpdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *AutoUpdateVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CreateAutoUpdateVersionRequest) Reset() {
	*x = CreateAutoUpdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAutoUpdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAutoUpdateVersionRequest) ProtoMessage() {}

func (x *CreateAutoUpdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAutoUpdateVersionRequest.ProtoReflect.Descriptor instead.
func (*CreateAutoUpdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAutoUpdateVersionRequest) GetVersion() *AutoUpdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// Request for UpdateAutoUpdateConfig.
type UpdateAutoUpdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *AutoUpdateVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UpdateAutoUpdateVersionRequest) Reset() {
	*x = UpdateAutoUpdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAutoUpdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAutoUpdateVersionRequest) ProtoMessage() {}

func (x *UpdateAutoUpdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAutoUpdateVersionRequest.ProtoReflect.Descriptor instead.
func (*UpdateAutoUpdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAutoUpdateVersionRequest) GetVersion() *AutoUpdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// Request for UpsertAutoUpdateVersion.
type UpsertAutoUpdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *AutoUpdateVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UpsertAutoUpdateVersionRequest) Reset() {
	*x = UpsertAutoUpdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAutoUpdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAutoUpdateVersionRequest) ProtoMessage() {}

func (x *UpsertAutoUpdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAutoUpdateVersionRequest.ProtoReflect.Descriptor instead.
func (*UpsertAutoUpdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpsertAutoUpdateVersionRequest) GetVersion() *AutoUpdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// Request for DeleteAutoUpdateVersion.
type DeleteAutoUpdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAutoUpdateVersionRequest) Reset() {
	*x = DeleteAutoUpdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAutoUpdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoUpdateVersionRequest) ProtoMessage() {}

func (x *DeleteAutoUpdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoUpdateVersionRequest.ProtoReflect.Descriptor instead.
func (*DeleteAutoUpdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{9}
}

var File_teleport_autoupdate_v1_autoupdate_service_proto protoreflect.FileDescriptor

var file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x61, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x61, 0x0a, 0x1d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a,
	0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x1e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75,
	0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x1e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xbf, 0x09,
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x79, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x79, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x79,
	0x0a, 0x16, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x67, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72,
	0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescOnce sync.Once
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData = file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc
)

func file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP() []byte {
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescOnce.Do(func() {
		file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData)
	})
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData
}

var file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes = []any{
	(*GetAutoUpdateConfigRequest)(nil),     // 0: teleport.autoupdate.v1.GetAutoUpdateConfigRequest
	(*CreateAutoUpdateConfigRequest)(nil),  // 1: teleport.autoupdate.v1.CreateAutoUpdateConfigRequest
	(*UpdateAutoUpdateConfigRequest)(nil),  // 2: teleport.autoupdate.v1.UpdateAutoUpdateConfigRequest
	(*UpsertAutoUpdateConfigRequest)(nil),  // 3: teleport.autoupdate.v1.UpsertAutoUpdateConfigRequest
	(*DeleteAutoUpdateConfigRequest)(nil),  // 4: teleport.autoupdate.v1.DeleteAutoUpdateConfigRequest
	(*GetAutoUpdateVersionRequest)(nil),    // 5: teleport.autoupdate.v1.GetAutoUpdateVersionRequest
	(*CreateAutoUpdateVersionRequest)(nil), // 6: teleport.autoupdate.v1.CreateAutoUpdateVersionRequest
	(*UpdateAutoUpdateVersionRequest)(nil), // 7: teleport.autoupdate.v1.UpdateAutoUpdateVersionRequest
	(*UpsertAutoUpdateVersionRequest)(nil), // 8: teleport.autoupdate.v1.UpsertAutoUpdateVersionRequest
	(*DeleteAutoUpdateVersionRequest)(nil), // 9: teleport.autoupdate.v1.DeleteAutoUpdateVersionRequest
	(*AutoUpdateConfig)(nil),               // 10: teleport.autoupdate.v1.AutoUpdateConfig
	(*AutoUpdateVersion)(nil),              // 11: teleport.autoupdate.v1.AutoUpdateVersion
	(*emptypb.Empty)(nil),                  // 12: google.protobuf.Empty
}
var file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs = []int32{
	10, // 0: teleport.autoupdate.v1.CreateAutoUpdateConfigRequest.config:type_name -> teleport.autoupdate.v1.AutoUpdateConfig
	10, // 1: teleport.autoupdate.v1.UpdateAutoUpdateConfigRequest.config:type_name -> teleport.autoupdate.v1.AutoUpdateConfig
	10, // 2: teleport.autoupdate.v1.UpsertAutoUpdateConfigRequest.config:type_name -> teleport.autoupdate.v1.AutoUpdateConfig
	11, // 3: teleport.autoupdate.v1.CreateAutoUpdateVersionRequest.version:type_name -> teleport.autoupdate.v1.AutoUpdateVersion
	11, // 4: teleport.autoupdate.v1.UpdateAutoUpdateVersionRequest.version:type_name -> teleport.autoupdate.v1.AutoUpdateVersion
	11, // 5: teleport.autoupdate.v1.UpsertAutoUpdateVersionRequest.version:type_name -> teleport.autoupdate.v1.AutoUpdateVersion
	0,  // 6: teleport.autoupdate.v1.AutoUpdateService.GetAutoUpdateConfig:input_type -> teleport.autoupdate.v1.GetAutoUpdateConfigRequest
	1,  // 7: teleport.autoupdate.v1.AutoUpdateService.CreateAutoUpdateConfig:input_type -> teleport.autoupdate.v1.CreateAutoUpdateConfigRequest
	2,  // 8: teleport.autoupdate.v1.AutoUpdateService.UpdateAutoUpdateConfig:input_type -> teleport.autoupdate.v1.UpdateAutoUpdateConfigRequest
	3,  // 9: teleport.autoupdate.v1.AutoUpdateService.UpsertAutoUpdateConfig:input_type -> teleport.autoupdate.v1.UpsertAutoUpdateConfigRequest
	4,  // 10: teleport.autoupdate.v1.AutoUpdateService.DeleteAutoUpdateConfig:input_type -> teleport.autoupdate.v1.DeleteAutoUpdateConfigRequest
	5,  // 11: teleport.autoupdate.v1.AutoUpdateService.GetAutoUpdateVersion:input_type -> teleport.autoupdate.v1.GetAutoUpdateVersionRequest
	6,  // 12: teleport.autoupdate.v1.AutoUpdateService.CreateAutoUpdateVersion:input_type -> teleport.autoupdate.v1.CreateAutoUpdateVersionRequest
	7,  // 13: teleport.autoupdate.v1.AutoUpdateService.UpdateAutoUpdateVersion:input_type -> teleport.autoupdate.v1.UpdateAutoUpdateVersionRequest
	8,  // 14: teleport.autoupdate.v1.AutoUpdateService.UpsertAutoUpdateVersion:input_type -> teleport.autoupdate.v1.UpsertAutoUpdateVersionRequest
	9,  // 15: teleport.autoupdate.v1.AutoUpdateService.DeleteAutoUpdateVersion:input_type -> teleport.autoupdate.v1.DeleteAutoUpdateVersionRequest
	10, // 16: teleport.autoupdate.v1.AutoUpdateService.GetAutoUpdateConfig:output_type -> teleport.autoupdate.v1.AutoUpdateConfig
	10, // 17: teleport.autoupdate.v1.AutoUpdateService.CreateAutoUpdateConfig:output_type -> teleport.autoupdate.v1.AutoUpdateConfig
	10, // 18: teleport.autoupdate.v1.AutoUpdateService.UpdateAutoUpdateConfig:output_type -> teleport.autoupdate.v1.AutoUpdateConfig
	10, // 19: teleport.autoupdate.v1.AutoUpdateService.UpsertAutoUpdateConfig:output_type -> teleport.autoupdate.v1.AutoUpdateConfig
	12, // 20: teleport.autoupdate.v1.AutoUpdateService.DeleteAutoUpdateConfig:output_type -> google.protobuf.Empty
	11, // 21: teleport.autoupdate.v1.AutoUpdateService.GetAutoUpdateVersion:output_type -> teleport.autoupdate.v1.AutoUpdateVersion
	11, // 22: teleport.autoupdate.v1.AutoUpdateService.CreateAutoUpdateVersion:output_type -> teleport.autoupdate.v1.AutoUpdateVersion
	11, // 23: teleport.autoupdate.v1.AutoUpdateService.UpdateAutoUpdateVersion:output_type -> teleport.autoupdate.v1.AutoUpdateVersion
	11, // 24: teleport.autoupdate.v1.AutoUpdateService.UpsertAutoUpdateVersion:output_type -> teleport.autoupdate.v1.AutoUpdateVersion
	12, // 25: teleport.autoupdate.v1.AutoUpdateService.DeleteAutoUpdateVersion:output_type -> google.protobuf.Empty
	16, // [16:26] is the sub-list for method output_type
	6,  // [6:16] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_autoupdate_v1_autoupdate_service_proto_init() }
func file_teleport_autoupdate_v1_autoupdate_service_proto_init() {
	if File_teleport_autoupdate_v1_autoupdate_service_proto != nil {
		return
	}
	file_teleport_autoupdate_v1_autoupdate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAutoUpdateConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAutoUpdateConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAutoUpdateConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAutoUpdateConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAutoUpdateConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAutoUpdateVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAutoUpdateVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAutoUpdateVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAutoUpdateVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAutoUpdateVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes,
		DependencyIndexes: file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs,
		MessageInfos:      file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes,
	}.Build()
	File_teleport_autoupdate_v1_autoupdate_service_proto = out.File
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc = nil
	file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes = nil
	file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs = nil
}
