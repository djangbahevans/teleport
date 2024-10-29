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
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: accessgraph/v1alpha/azure.proto

package accessgraphv1alpha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AzureResourceList is a list of Azure resources
type AzureResourceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*AzureResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *AzureResourceList) Reset() {
	*x = AzureResourceList{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureResourceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureResourceList) ProtoMessage() {}

func (x *AzureResourceList) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureResourceList.ProtoReflect.Descriptor instead.
func (*AzureResourceList) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{0}
}

func (x *AzureResourceList) GetResources() []*AzureResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// AWSResource is a list of AWS resources supported by the access graph.
type AzureResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resource:
	//
	//	*AzureResource_VirtualMachine
	//	*AzureResource_User
	//	*AzureResource_RoleDefinition
	//	*AzureResource_RoleAssignment
	Resource isAzureResource_Resource `protobuf_oneof:"resource"`
}

func (x *AzureResource) Reset() {
	*x = AzureResource{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureResource) ProtoMessage() {}

func (x *AzureResource) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureResource.ProtoReflect.Descriptor instead.
func (*AzureResource) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{1}
}

func (m *AzureResource) GetResource() isAzureResource_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *AzureResource) GetVirtualMachine() *AzureVirtualMachine {
	if x, ok := x.GetResource().(*AzureResource_VirtualMachine); ok {
		return x.VirtualMachine
	}
	return nil
}

func (x *AzureResource) GetUser() *AzureUser {
	if x, ok := x.GetResource().(*AzureResource_User); ok {
		return x.User
	}
	return nil
}

func (x *AzureResource) GetRoleDefinition() *AzureRoleDefinition {
	if x, ok := x.GetResource().(*AzureResource_RoleDefinition); ok {
		return x.RoleDefinition
	}
	return nil
}

func (x *AzureResource) GetRoleAssignment() *AzureRoleAssignment {
	if x, ok := x.GetResource().(*AzureResource_RoleAssignment); ok {
		return x.RoleAssignment
	}
	return nil
}

type isAzureResource_Resource interface {
	isAzureResource_Resource()
}

type AzureResource_VirtualMachine struct {
	VirtualMachine *AzureVirtualMachine `protobuf:"bytes,1,opt,name=virtual_machine,json=virtualMachine,proto3,oneof"`
}

type AzureResource_User struct {
	User *AzureUser `protobuf:"bytes,2,opt,name=user,proto3,oneof"`
}

type AzureResource_RoleDefinition struct {
	RoleDefinition *AzureRoleDefinition `protobuf:"bytes,3,opt,name=role_definition,json=roleDefinition,proto3,oneof"`
}

type AzureResource_RoleAssignment struct {
	RoleAssignment *AzureRoleAssignment `protobuf:"bytes,4,opt,name=role_assignment,json=roleAssignment,proto3,oneof"`
}

func (*AzureResource_VirtualMachine) isAzureResource_Resource() {}

func (*AzureResource_User) isAzureResource_Resource() {}

func (*AzureResource_RoleDefinition) isAzureResource_Resource() {}

func (*AzureResource_RoleAssignment) isAzureResource_Resource() {}

// AzureVirtualMachine is an Azure virtual machine
type AzureVirtualMachine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionId string                 `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	LastSyncTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	Name           string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AzureVirtualMachine) Reset() {
	*x = AzureVirtualMachine{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureVirtualMachine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureVirtualMachine) ProtoMessage() {}

func (x *AzureVirtualMachine) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureVirtualMachine.ProtoReflect.Descriptor instead.
func (*AzureVirtualMachine) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{2}
}

func (x *AzureVirtualMachine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AzureVirtualMachine) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureVirtualMachine) GetLastSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSyncTime
	}
	return nil
}

func (x *AzureVirtualMachine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// AzureUser is an Azure user
type AzureUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionId string                 `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	LastSyncTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	DisplayName    string                 `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *AzureUser) Reset() {
	*x = AzureUser{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureUser) ProtoMessage() {}

func (x *AzureUser) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureUser.ProtoReflect.Descriptor instead.
func (*AzureUser) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{3}
}

func (x *AzureUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AzureUser) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureUser) GetLastSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSyncTime
	}
	return nil
}

func (x *AzureUser) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// AzureRoleAssignment links an Azure principal to a role definition with a scope
type AzureRoleAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionId   string                 `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	LastSyncTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	PrincipalId      string                 `protobuf:"bytes,4,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
	RoleDefinitionId string                 `protobuf:"bytes,5,opt,name=role_definition_id,json=roleDefinitionId,proto3" json:"role_definition_id,omitempty"`
	Scope            string                 `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *AzureRoleAssignment) Reset() {
	*x = AzureRoleAssignment{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureRoleAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureRoleAssignment) ProtoMessage() {}

func (x *AzureRoleAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureRoleAssignment.ProtoReflect.Descriptor instead.
func (*AzureRoleAssignment) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{4}
}

func (x *AzureRoleAssignment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AzureRoleAssignment) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureRoleAssignment) GetLastSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSyncTime
	}
	return nil
}

func (x *AzureRoleAssignment) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

func (x *AzureRoleAssignment) GetRoleDefinitionId() string {
	if x != nil {
		return x.RoleDefinitionId
	}
	return ""
}

func (x *AzureRoleAssignment) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

// AzureRoleDefinition defines a role by its permissions
type AzureRoleDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubscriptionId string                 `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	LastSyncTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	Permissions    []*AzurePermission     `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *AzureRoleDefinition) Reset() {
	*x = AzureRoleDefinition{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureRoleDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureRoleDefinition) ProtoMessage() {}

func (x *AzureRoleDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureRoleDefinition.ProtoReflect.Descriptor instead.
func (*AzureRoleDefinition) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{5}
}

func (x *AzureRoleDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AzureRoleDefinition) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureRoleDefinition) GetLastSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSyncTime
	}
	return nil
}

func (x *AzureRoleDefinition) GetPermissions() []*AzurePermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// AzurePermission defines the actions and not (disallowed) actions for a role definition
type AzurePermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions    []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	NotActions []string `protobuf:"bytes,5,rep,name=not_actions,json=notActions,proto3" json:"not_actions,omitempty"`
}

func (x *AzurePermission) Reset() {
	*x = AzurePermission{}
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzurePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzurePermission) ProtoMessage() {}

func (x *AzurePermission) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_azure_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzurePermission.ProtoReflect.Descriptor instead.
func (*AzurePermission) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_azure_proto_rawDescGZIP(), []int{6}
}

func (x *AzurePermission) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *AzurePermission) GetNotActions() []string {
	if x != nil {
		return x.NotActions
	}
	return nil
}

var File_accessgraph_v1alpha_azure_proto protoreflect.FileDescriptor

var file_accessgraph_v1alpha_azure_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xd0,
	0x02, 0x0a, 0x0d, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x53, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0xa4, 0x01, 0x0a, 0x13, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x13, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0xd8,
	0x01, 0x0a, 0x13, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x0f, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accessgraph_v1alpha_azure_proto_rawDescOnce sync.Once
	file_accessgraph_v1alpha_azure_proto_rawDescData = file_accessgraph_v1alpha_azure_proto_rawDesc
)

func file_accessgraph_v1alpha_azure_proto_rawDescGZIP() []byte {
	file_accessgraph_v1alpha_azure_proto_rawDescOnce.Do(func() {
		file_accessgraph_v1alpha_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_accessgraph_v1alpha_azure_proto_rawDescData)
	})
	return file_accessgraph_v1alpha_azure_proto_rawDescData
}

var file_accessgraph_v1alpha_azure_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_accessgraph_v1alpha_azure_proto_goTypes = []any{
	(*AzureResourceList)(nil),     // 0: accessgraph.v1alpha.AzureResourceList
	(*AzureResource)(nil),         // 1: accessgraph.v1alpha.AzureResource
	(*AzureVirtualMachine)(nil),   // 2: accessgraph.v1alpha.AzureVirtualMachine
	(*AzureUser)(nil),             // 3: accessgraph.v1alpha.AzureUser
	(*AzureRoleAssignment)(nil),   // 4: accessgraph.v1alpha.AzureRoleAssignment
	(*AzureRoleDefinition)(nil),   // 5: accessgraph.v1alpha.AzureRoleDefinition
	(*AzurePermission)(nil),       // 6: accessgraph.v1alpha.AzurePermission
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_accessgraph_v1alpha_azure_proto_depIdxs = []int32{
	1,  // 0: accessgraph.v1alpha.AzureResourceList.resources:type_name -> accessgraph.v1alpha.AzureResource
	2,  // 1: accessgraph.v1alpha.AzureResource.virtual_machine:type_name -> accessgraph.v1alpha.AzureVirtualMachine
	3,  // 2: accessgraph.v1alpha.AzureResource.user:type_name -> accessgraph.v1alpha.AzureUser
	5,  // 3: accessgraph.v1alpha.AzureResource.role_definition:type_name -> accessgraph.v1alpha.AzureRoleDefinition
	4,  // 4: accessgraph.v1alpha.AzureResource.role_assignment:type_name -> accessgraph.v1alpha.AzureRoleAssignment
	7,  // 5: accessgraph.v1alpha.AzureVirtualMachine.last_sync_time:type_name -> google.protobuf.Timestamp
	7,  // 6: accessgraph.v1alpha.AzureUser.last_sync_time:type_name -> google.protobuf.Timestamp
	7,  // 7: accessgraph.v1alpha.AzureRoleAssignment.last_sync_time:type_name -> google.protobuf.Timestamp
	7,  // 8: accessgraph.v1alpha.AzureRoleDefinition.last_sync_time:type_name -> google.protobuf.Timestamp
	6,  // 9: accessgraph.v1alpha.AzureRoleDefinition.permissions:type_name -> accessgraph.v1alpha.AzurePermission
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_accessgraph_v1alpha_azure_proto_init() }
func file_accessgraph_v1alpha_azure_proto_init() {
	if File_accessgraph_v1alpha_azure_proto != nil {
		return
	}
	file_accessgraph_v1alpha_azure_proto_msgTypes[1].OneofWrappers = []any{
		(*AzureResource_VirtualMachine)(nil),
		(*AzureResource_User)(nil),
		(*AzureResource_RoleDefinition)(nil),
		(*AzureResource_RoleAssignment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accessgraph_v1alpha_azure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accessgraph_v1alpha_azure_proto_goTypes,
		DependencyIndexes: file_accessgraph_v1alpha_azure_proto_depIdxs,
		MessageInfos:      file_accessgraph_v1alpha_azure_proto_msgTypes,
	}.Build()
	File_accessgraph_v1alpha_azure_proto = out.File
	file_accessgraph_v1alpha_azure_proto_rawDesc = nil
	file_accessgraph_v1alpha_azure_proto_goTypes = nil
	file_accessgraph_v1alpha_azure_proto_depIdxs = nil
}
