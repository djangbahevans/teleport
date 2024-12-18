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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/assist.proto

package userpreferencesv1

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

// AssistViewMode is the way the assistant is displayed.
type AssistViewMode int32

const (
	AssistViewMode_ASSIST_VIEW_MODE_UNSPECIFIED AssistViewMode = 0
	// DOCKED is the assistant is docked to the right hand side of the screen.
	AssistViewMode_ASSIST_VIEW_MODE_DOCKED AssistViewMode = 1
	// POPUP is the assistant is displayed as a popup.
	AssistViewMode_ASSIST_VIEW_MODE_POPUP AssistViewMode = 2
	// POPUP_EXPANDED is the assistant is displayed as a popup and expanded.
	AssistViewMode_ASSIST_VIEW_MODE_POPUP_EXPANDED AssistViewMode = 3
	// POPUP_EXPANDED_SIDEBAR_VISIBLE is the assistant is displayed as a popup and expanded with the sidebar visible.
	AssistViewMode_ASSIST_VIEW_MODE_POPUP_EXPANDED_SIDEBAR_VISIBLE AssistViewMode = 4
)

// Enum value maps for AssistViewMode.
var (
	AssistViewMode_name = map[int32]string{
		0: "ASSIST_VIEW_MODE_UNSPECIFIED",
		1: "ASSIST_VIEW_MODE_DOCKED",
		2: "ASSIST_VIEW_MODE_POPUP",
		3: "ASSIST_VIEW_MODE_POPUP_EXPANDED",
		4: "ASSIST_VIEW_MODE_POPUP_EXPANDED_SIDEBAR_VISIBLE",
	}
	AssistViewMode_value = map[string]int32{
		"ASSIST_VIEW_MODE_UNSPECIFIED":                    0,
		"ASSIST_VIEW_MODE_DOCKED":                         1,
		"ASSIST_VIEW_MODE_POPUP":                          2,
		"ASSIST_VIEW_MODE_POPUP_EXPANDED":                 3,
		"ASSIST_VIEW_MODE_POPUP_EXPANDED_SIDEBAR_VISIBLE": 4,
	}
)

func (x AssistViewMode) Enum() *AssistViewMode {
	p := new(AssistViewMode)
	*p = x
	return p
}

func (x AssistViewMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssistViewMode) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_userpreferences_v1_assist_proto_enumTypes[0].Descriptor()
}

func (AssistViewMode) Type() protoreflect.EnumType {
	return &file_teleport_userpreferences_v1_assist_proto_enumTypes[0]
}

func (x AssistViewMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssistViewMode.Descriptor instead.
func (AssistViewMode) EnumDescriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_assist_proto_rawDescGZIP(), []int{0}
}

// AssistUserPreferences is the user preferences for Assist.
type AssistUserPreferences struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// preferredLogins is an array of the logins a user would prefer to use when running a command, ordered by preference.
	PreferredLogins []string `protobuf:"bytes,1,rep,name=preferred_logins,json=preferredLogins,proto3" json:"preferred_logins,omitempty"`
	// viewMode is the way the assistant is displayed.
	ViewMode      AssistViewMode `protobuf:"varint,2,opt,name=view_mode,json=viewMode,proto3,enum=teleport.userpreferences.v1.AssistViewMode" json:"view_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssistUserPreferences) Reset() {
	*x = AssistUserPreferences{}
	mi := &file_teleport_userpreferences_v1_assist_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssistUserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistUserPreferences) ProtoMessage() {}

func (x *AssistUserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_assist_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistUserPreferences.ProtoReflect.Descriptor instead.
func (*AssistUserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_assist_proto_rawDescGZIP(), []int{0}
}

func (x *AssistUserPreferences) GetPreferredLogins() []string {
	if x != nil {
		return x.PreferredLogins
	}
	return nil
}

func (x *AssistUserPreferences) GetViewMode() AssistViewMode {
	if x != nil {
		return x.ViewMode
	}
	return AssistViewMode_ASSIST_VIEW_MODE_UNSPECIFIED
}

var File_teleport_userpreferences_v1_assist_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_assist_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x8c, 0x01, 0x0a, 0x15, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x76, 0x69,
	0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x2a, 0xc5, 0x01, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x53, 0x53,
	0x49, 0x53, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41,
	0x53, 0x53, 0x49, 0x53, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x44, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x53, 0x53, 0x49,
	0x53, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x50,
	0x55, 0x50, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x5f, 0x56,
	0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x50, 0x55, 0x50, 0x5f, 0x45,
	0x58, 0x50, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x33, 0x0a, 0x2f, 0x41, 0x53, 0x53,
	0x49, 0x53, 0x54, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f,
	0x50, 0x55, 0x50, 0x5f, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x53, 0x49, 0x44,
	0x45, 0x42, 0x41, 0x52, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x42, 0x59,
	0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_teleport_userpreferences_v1_assist_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_assist_proto_rawDescData = file_teleport_userpreferences_v1_assist_proto_rawDesc
)

func file_teleport_userpreferences_v1_assist_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_assist_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_assist_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userpreferences_v1_assist_proto_rawDescData)
	})
	return file_teleport_userpreferences_v1_assist_proto_rawDescData
}

var file_teleport_userpreferences_v1_assist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_userpreferences_v1_assist_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_userpreferences_v1_assist_proto_goTypes = []any{
	(AssistViewMode)(0),           // 0: teleport.userpreferences.v1.AssistViewMode
	(*AssistUserPreferences)(nil), // 1: teleport.userpreferences.v1.AssistUserPreferences
}
var file_teleport_userpreferences_v1_assist_proto_depIdxs = []int32{
	0, // 0: teleport.userpreferences.v1.AssistUserPreferences.view_mode:type_name -> teleport.userpreferences.v1.AssistViewMode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_assist_proto_init() }
func file_teleport_userpreferences_v1_assist_proto_init() {
	if File_teleport_userpreferences_v1_assist_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_userpreferences_v1_assist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userpreferences_v1_assist_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_assist_proto_depIdxs,
		EnumInfos:         file_teleport_userpreferences_v1_assist_proto_enumTypes,
		MessageInfos:      file_teleport_userpreferences_v1_assist_proto_msgTypes,
	}.Build()
	File_teleport_userpreferences_v1_assist_proto = out.File
	file_teleport_userpreferences_v1_assist_proto_rawDesc = nil
	file_teleport_userpreferences_v1_assist_proto_goTypes = nil
	file_teleport_userpreferences_v1_assist_proto_depIdxs = nil
}
