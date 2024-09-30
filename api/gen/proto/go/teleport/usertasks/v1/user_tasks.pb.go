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
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: teleport/usertasks/v1/user_tasks.proto

package usertasksv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// UserTask is a resource that represents an action to be completed by the user.
// UserTasks are a unit of work for users to act upon issues related to other resources.
// As an example, when auto-enrolling EC2 instances using the Discovery Service
// a UserTask is created to let the user know that something failed on a set of instances.
// The user can then mark the task as resolved after following the recommendation/fixing steps.
type UserTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind of resource represented.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Mandatory field for all resources. Not populated for this resource type.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// The version of the resource being represented.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Common metadata that all resources share.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The configured properties of UserTask.
	Spec *UserTaskSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *UserTask) Reset() {
	*x = UserTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTask) ProtoMessage() {}

func (x *UserTask) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTask.ProtoReflect.Descriptor instead.
func (*UserTask) Descriptor() ([]byte, []int) {
	return file_teleport_usertasks_v1_user_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *UserTask) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *UserTask) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *UserTask) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UserTask) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UserTask) GetSpec() *UserTaskSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// UserTaskSpec contains the properties of the UserTask.
type UserTaskSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Integration is the integration name that originated this task.
	Integration string `protobuf:"bytes,1,opt,name=integration,proto3" json:"integration,omitempty"`
	// TaskType indicates the type of task.
	// Examples: discover-ec2, discover-rds, discover-eks
	TaskType string `protobuf:"bytes,2,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// IssueType is an identifier for the type of issue that happened.
	// Example for discover-ec2: SSM_AGENT_NOT_AVAILABLE
	IssueType string `protobuf:"bytes,3,opt,name=issue_type,json=issueType,proto3" json:"issue_type,omitempty"`
	// State indicates the task state.
	// When the task is created, it starts with OPEN.
	// Users can mark it as RESOLVED.
	// If the issue happens again (eg, new discover iteration faces the same error), it will move to OPEN again.
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// DiscoverEC2 contains the AWS EC2 instances that failed to auto enroll into the cluster.
	// Present when TaskType is discover-ec2.
	DiscoverEc2 *DiscoverEC2 `protobuf:"bytes,5,opt,name=discover_ec2,json=discoverEc2,proto3" json:"discover_ec2,omitempty"`
}

func (x *UserTaskSpec) Reset() {
	*x = UserTaskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTaskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTaskSpec) ProtoMessage() {}

func (x *UserTaskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTaskSpec.ProtoReflect.Descriptor instead.
func (*UserTaskSpec) Descriptor() ([]byte, []int) {
	return file_teleport_usertasks_v1_user_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *UserTaskSpec) GetIntegration() string {
	if x != nil {
		return x.Integration
	}
	return ""
}

func (x *UserTaskSpec) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *UserTaskSpec) GetIssueType() string {
	if x != nil {
		return x.IssueType
	}
	return ""
}

func (x *UserTaskSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UserTaskSpec) GetDiscoverEc2() *DiscoverEC2 {
	if x != nil {
		return x.DiscoverEc2
	}
	return nil
}

// DiscoverEC2 contains the instances that failed to auto-enroll into the cluster.
type DiscoverEC2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instances maps an instance id to the result of enrolling that instance into teleport.
	Instances map[string]*DiscoverEC2Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DiscoverEC2) Reset() {
	*x = DiscoverEC2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverEC2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverEC2) ProtoMessage() {}

func (x *DiscoverEC2) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverEC2.ProtoReflect.Descriptor instead.
func (*DiscoverEC2) Descriptor() ([]byte, []int) {
	return file_teleport_usertasks_v1_user_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverEC2) GetInstances() map[string]*DiscoverEC2Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

// DiscoverEC2Instance contains the result of enrolling an AWS EC2 Instance.
type DiscoverEC2Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// InstanceID is the EC2 Instance ID that uniquely identifies the instance.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Name is the instance Name.
	// Might be empty, if the instance doesn't have the Name tag.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// AccountID is the AWS Account ID for this instance.
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Region is the AWS Region where this issue is happening.
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// InvocationURL is the URL that points to the invocation.
	// Empty if there was an error before installing the
	InvocationUrl string `protobuf:"bytes,5,opt,name=invocation_url,json=invocationUrl,proto3" json:"invocation_url,omitempty"`
	// DiscoveryConfig is the discovery config name that originated this instance enrollment.
	DiscoveryConfig string `protobuf:"bytes,6,opt,name=discovery_config,json=discoveryConfig,proto3" json:"discovery_config,omitempty"`
	// DiscoveryGroup is the DiscoveryGroup name that originated this task.
	DiscoveryGroup string `protobuf:"bytes,7,opt,name=discovery_group,json=discoveryGroup,proto3" json:"discovery_group,omitempty"`
	// SyncTime is the timestamp when the error was produced.
	SyncTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=sync_time,json=syncTime,proto3" json:"sync_time,omitempty"`
}

func (x *DiscoverEC2Instance) Reset() {
	*x = DiscoverEC2Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverEC2Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverEC2Instance) ProtoMessage() {}

func (x *DiscoverEC2Instance) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_usertasks_v1_user_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverEC2Instance.ProtoReflect.Descriptor instead.
func (*DiscoverEC2Instance) Descriptor() ([]byte, []int) {
	return file_teleport_usertasks_v1_user_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverEC2Instance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *DiscoverEC2Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscoverEC2Instance) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *DiscoverEC2Instance) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DiscoverEC2Instance) GetInvocationUrl() string {
	if x != nil {
		return x.InvocationUrl
	}
	return ""
}

func (x *DiscoverEC2Instance) GetDiscoveryConfig() string {
	if x != nil {
		return x.DiscoveryConfig
	}
	return ""
}

func (x *DiscoverEC2Instance) GetDiscoveryGroup() string {
	if x != nil {
		return x.DiscoveryGroup
	}
	return ""
}

func (x *DiscoverEC2Instance) GetSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SyncTime
	}
	return nil
}

var File_teleport_usertasks_v1_user_tasks_proto protoreflect.FileDescriptor

var file_teleport_usertasks_v1_user_tasks_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xc9, 0x01, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x63,
	0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x43, 0x32, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x63, 0x32, 0x22, 0xc8, 0x01, 0x0a, 0x0b, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x43, 0x32, 0x12, 0x4f, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x43, 0x32, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x68, 0x0a, 0x0e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x43, 0x32,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xb5, 0x02, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x45, 0x43, 0x32, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x29,
	0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x56, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_usertasks_v1_user_tasks_proto_rawDescOnce sync.Once
	file_teleport_usertasks_v1_user_tasks_proto_rawDescData = file_teleport_usertasks_v1_user_tasks_proto_rawDesc
)

func file_teleport_usertasks_v1_user_tasks_proto_rawDescGZIP() []byte {
	file_teleport_usertasks_v1_user_tasks_proto_rawDescOnce.Do(func() {
		file_teleport_usertasks_v1_user_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_usertasks_v1_user_tasks_proto_rawDescData)
	})
	return file_teleport_usertasks_v1_user_tasks_proto_rawDescData
}

var file_teleport_usertasks_v1_user_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_usertasks_v1_user_tasks_proto_goTypes = []interface{}{
	(*UserTask)(nil),              // 0: teleport.usertasks.v1.UserTask
	(*UserTaskSpec)(nil),          // 1: teleport.usertasks.v1.UserTaskSpec
	(*DiscoverEC2)(nil),           // 2: teleport.usertasks.v1.DiscoverEC2
	(*DiscoverEC2Instance)(nil),   // 3: teleport.usertasks.v1.DiscoverEC2Instance
	nil,                           // 4: teleport.usertasks.v1.DiscoverEC2.InstancesEntry
	(*v1.Metadata)(nil),           // 5: teleport.header.v1.Metadata
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_teleport_usertasks_v1_user_tasks_proto_depIdxs = []int32{
	5, // 0: teleport.usertasks.v1.UserTask.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.usertasks.v1.UserTask.spec:type_name -> teleport.usertasks.v1.UserTaskSpec
	2, // 2: teleport.usertasks.v1.UserTaskSpec.discover_ec2:type_name -> teleport.usertasks.v1.DiscoverEC2
	4, // 3: teleport.usertasks.v1.DiscoverEC2.instances:type_name -> teleport.usertasks.v1.DiscoverEC2.InstancesEntry
	6, // 4: teleport.usertasks.v1.DiscoverEC2Instance.sync_time:type_name -> google.protobuf.Timestamp
	3, // 5: teleport.usertasks.v1.DiscoverEC2.InstancesEntry.value:type_name -> teleport.usertasks.v1.DiscoverEC2Instance
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_usertasks_v1_user_tasks_proto_init() }
func file_teleport_usertasks_v1_user_tasks_proto_init() {
	if File_teleport_usertasks_v1_user_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_usertasks_v1_user_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTask); i {
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
		file_teleport_usertasks_v1_user_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTaskSpec); i {
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
		file_teleport_usertasks_v1_user_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverEC2); i {
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
		file_teleport_usertasks_v1_user_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverEC2Instance); i {
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
			RawDescriptor: file_teleport_usertasks_v1_user_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_usertasks_v1_user_tasks_proto_goTypes,
		DependencyIndexes: file_teleport_usertasks_v1_user_tasks_proto_depIdxs,
		MessageInfos:      file_teleport_usertasks_v1_user_tasks_proto_msgTypes,
	}.Build()
	File_teleport_usertasks_v1_user_tasks_proto = out.File
	file_teleport_usertasks_v1_user_tasks_proto_rawDesc = nil
	file_teleport_usertasks_v1_user_tasks_proto_goTypes = nil
	file_teleport_usertasks_v1_user_tasks_proto_depIdxs = nil
}
