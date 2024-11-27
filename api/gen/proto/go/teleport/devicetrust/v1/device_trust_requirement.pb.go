// Copyright 2024 Gravitational, Inc
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
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: teleport/devicetrust/v1/device_trust_requirement.proto

package devicetrustv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TrustedDeviceRequirement indicates whether access may be hindered by the lack
// of a trusted device.
type TrustedDeviceRequirement int32

const (
	// Device requirement not determined.
	// Does not mean that a device is not required, only that the necessary data
	// was not considered.
	TrustedDeviceRequirement_TRUSTED_DEVICE_REQUIREMENT_UNSPECIFIED TrustedDeviceRequirement = 0
	// Trusted device not required.
	TrustedDeviceRequirement_TRUSTED_DEVICE_REQUIREMENT_NOT_REQUIRED TrustedDeviceRequirement = 1
	// Trusted device required by either cluster mode or user roles.
	TrustedDeviceRequirement_TRUSTED_DEVICE_REQUIREMENT_REQUIRED TrustedDeviceRequirement = 2
)

// Enum value maps for TrustedDeviceRequirement.
var (
	TrustedDeviceRequirement_name = map[int32]string{
		0: "TRUSTED_DEVICE_REQUIREMENT_UNSPECIFIED",
		1: "TRUSTED_DEVICE_REQUIREMENT_NOT_REQUIRED",
		2: "TRUSTED_DEVICE_REQUIREMENT_REQUIRED",
	}
	TrustedDeviceRequirement_value = map[string]int32{
		"TRUSTED_DEVICE_REQUIREMENT_UNSPECIFIED":  0,
		"TRUSTED_DEVICE_REQUIREMENT_NOT_REQUIRED": 1,
		"TRUSTED_DEVICE_REQUIREMENT_REQUIRED":     2,
	}
)

func (x TrustedDeviceRequirement) Enum() *TrustedDeviceRequirement {
	p := new(TrustedDeviceRequirement)
	*p = x
	return p
}

func (x TrustedDeviceRequirement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrustedDeviceRequirement) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_devicetrust_v1_device_trust_requirement_proto_enumTypes[0].Descriptor()
}

func (TrustedDeviceRequirement) Type() protoreflect.EnumType {
	return &file_teleport_devicetrust_v1_device_trust_requirement_proto_enumTypes[0]
}

func (x TrustedDeviceRequirement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrustedDeviceRequirement.Descriptor instead.
func (TrustedDeviceRequirement) EnumDescriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescGZIP(), []int{0}
}

var File_teleport_devicetrust_v1_device_trust_requirement_proto protoreflect.FileDescriptor

var file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2a, 0x9c, 0x01, 0x0a, 0x18, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x26, 0x54, 0x52, 0x55, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x52,
	0x55, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x49, 0x52, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x55, 0x53, 0x54,
	0x45, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescOnce sync.Once
	file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescData = file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDesc
)

func file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescGZIP() []byte {
	file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescOnce.Do(func() {
		file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescData)
	})
	return file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDescData
}

var file_teleport_devicetrust_v1_device_trust_requirement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_devicetrust_v1_device_trust_requirement_proto_goTypes = []any{
	(TrustedDeviceRequirement)(0), // 0: teleport.devicetrust.v1.TrustedDeviceRequirement
}
var file_teleport_devicetrust_v1_device_trust_requirement_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_devicetrust_v1_device_trust_requirement_proto_init() }
func file_teleport_devicetrust_v1_device_trust_requirement_proto_init() {
	if File_teleport_devicetrust_v1_device_trust_requirement_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_devicetrust_v1_device_trust_requirement_proto_goTypes,
		DependencyIndexes: file_teleport_devicetrust_v1_device_trust_requirement_proto_depIdxs,
		EnumInfos:         file_teleport_devicetrust_v1_device_trust_requirement_proto_enumTypes,
	}.Build()
	File_teleport_devicetrust_v1_device_trust_requirement_proto = out.File
	file_teleport_devicetrust_v1_device_trust_requirement_proto_rawDesc = nil
	file_teleport_devicetrust_v1_device_trust_requirement_proto_goTypes = nil
	file_teleport_devicetrust_v1_device_trust_requirement_proto_depIdxs = nil
}
