//
// Teleport
// Copyright (C) 2024  Gravitational, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: accessgraph/v1alpha/entra.proto

package accessgraphv1alpha

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

// EntraSyncOperation is a request to sync Entra resources
type EntraSyncOperation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntraSyncOperation) Reset() {
	*x = EntraSyncOperation{}
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntraSyncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntraSyncOperation) ProtoMessage() {}

func (x *EntraSyncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntraSyncOperation.ProtoReflect.Descriptor instead.
func (*EntraSyncOperation) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_entra_proto_rawDescGZIP(), []int{0}
}

// EntraResourceList is a request that contains resources to be sync.
type EntraResourceList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// resources is a list of entra resources to sync.
	Resources     []*EntraResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntraResourceList) Reset() {
	*x = EntraResourceList{}
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntraResourceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntraResourceList) ProtoMessage() {}

func (x *EntraResourceList) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntraResourceList.ProtoReflect.Descriptor instead.
func (*EntraResourceList) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_entra_proto_rawDescGZIP(), []int{1}
}

func (x *EntraResourceList) GetResources() []*EntraResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// EntraResource represents a Entra resource.
type EntraResource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Resource:
	//
	//	*EntraResource_Application
	Resource      isEntraResource_Resource `protobuf_oneof:"resource"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntraResource) Reset() {
	*x = EntraResource{}
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntraResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntraResource) ProtoMessage() {}

func (x *EntraResource) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntraResource.ProtoReflect.Descriptor instead.
func (*EntraResource) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_entra_proto_rawDescGZIP(), []int{2}
}

func (x *EntraResource) GetResource() isEntraResource_Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *EntraResource) GetApplication() *EntraApplication {
	if x != nil {
		if x, ok := x.Resource.(*EntraResource_Application); ok {
			return x.Application
		}
	}
	return nil
}

type isEntraResource_Resource interface {
	isEntraResource_Resource()
}

type EntraResource_Application struct {
	// application represents an Entra ID enterprise application.
	Application *EntraApplication `protobuf:"bytes,1,opt,name=application,proto3,oneof"`
}

func (*EntraResource_Application) isEntraResource_Resource() {}

// EntraApplication represents an Entra ID enterprise application together with its service principal.
type EntraApplication struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id is the unique Entra object ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// app_id is the application ID.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// display_name is a human-friendly application name.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// tenant_id is the ID of Entra tenant that this application is under.
	TenantId string `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// signing_certificates is a list of SAML signing certificates for this app.
	SigningCertificates []string `protobuf:"bytes,5,rep,name=signing_certificates,json=signingCertificates,proto3" json:"signing_certificates,omitempty"`
	// federated_sso_v2 contains payload from the /ApplicationSso/{servicePrincipalId}/FederatedSSOV2 endpoint.
	// It is exposed from the internal plugin cache as an opaque JSON blob.
	FederatedSsoV2 string `protobuf:"bytes,6,opt,name=federated_sso_v2,json=federatedSsoV2,proto3" json:"federated_sso_v2,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EntraApplication) Reset() {
	*x = EntraApplication{}
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntraApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntraApplication) ProtoMessage() {}

func (x *EntraApplication) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_entra_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntraApplication.ProtoReflect.Descriptor instead.
func (*EntraApplication) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_entra_proto_rawDescGZIP(), []int{3}
}

func (x *EntraApplication) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntraApplication) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *EntraApplication) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *EntraApplication) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *EntraApplication) GetSigningCertificates() []string {
	if x != nil {
		return x.SigningCertificates
	}
	return nil
}

func (x *EntraApplication) GetFederatedSsoV2() string {
	if x != nil {
		return x.FederatedSsoV2
	}
	return ""
}

var File_accessgraph_v1alpha_entra_proto protoreflect.FileDescriptor

var file_accessgraph_v1alpha_entra_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x53,
	0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x11,
	0x45, 0x6e, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x61,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x61, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x10,
	0x45, 0x6e, 0x74, 0x72, 0x61, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x73, 0x6f, 0x5f, 0x76, 0x32, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x73, 0x6f, 0x56, 0x32, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accessgraph_v1alpha_entra_proto_rawDescOnce sync.Once
	file_accessgraph_v1alpha_entra_proto_rawDescData = file_accessgraph_v1alpha_entra_proto_rawDesc
)

func file_accessgraph_v1alpha_entra_proto_rawDescGZIP() []byte {
	file_accessgraph_v1alpha_entra_proto_rawDescOnce.Do(func() {
		file_accessgraph_v1alpha_entra_proto_rawDescData = protoimpl.X.CompressGZIP(file_accessgraph_v1alpha_entra_proto_rawDescData)
	})
	return file_accessgraph_v1alpha_entra_proto_rawDescData
}

var file_accessgraph_v1alpha_entra_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accessgraph_v1alpha_entra_proto_goTypes = []any{
	(*EntraSyncOperation)(nil), // 0: accessgraph.v1alpha.EntraSyncOperation
	(*EntraResourceList)(nil),  // 1: accessgraph.v1alpha.EntraResourceList
	(*EntraResource)(nil),      // 2: accessgraph.v1alpha.EntraResource
	(*EntraApplication)(nil),   // 3: accessgraph.v1alpha.EntraApplication
}
var file_accessgraph_v1alpha_entra_proto_depIdxs = []int32{
	2, // 0: accessgraph.v1alpha.EntraResourceList.resources:type_name -> accessgraph.v1alpha.EntraResource
	3, // 1: accessgraph.v1alpha.EntraResource.application:type_name -> accessgraph.v1alpha.EntraApplication
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accessgraph_v1alpha_entra_proto_init() }
func file_accessgraph_v1alpha_entra_proto_init() {
	if File_accessgraph_v1alpha_entra_proto != nil {
		return
	}
	file_accessgraph_v1alpha_entra_proto_msgTypes[2].OneofWrappers = []any{
		(*EntraResource_Application)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accessgraph_v1alpha_entra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accessgraph_v1alpha_entra_proto_goTypes,
		DependencyIndexes: file_accessgraph_v1alpha_entra_proto_depIdxs,
		MessageInfos:      file_accessgraph_v1alpha_entra_proto_msgTypes,
	}.Build()
	File_accessgraph_v1alpha_entra_proto = out.File
	file_accessgraph_v1alpha_entra_proto_rawDesc = nil
	file_accessgraph_v1alpha_entra_proto_goTypes = nil
	file_accessgraph_v1alpha_entra_proto_depIdxs = nil
}
