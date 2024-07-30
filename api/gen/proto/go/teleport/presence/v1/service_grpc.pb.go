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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.0
// - protoc             (unknown)
// source: teleport/presence/v1/service.proto

package presencev1

import (
	context "context"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PresenceService_GetRemoteCluster_FullMethodName    = "/teleport.presence.v1.PresenceService/GetRemoteCluster"
	PresenceService_ListRemoteClusters_FullMethodName  = "/teleport.presence.v1.PresenceService/ListRemoteClusters"
	PresenceService_UpdateRemoteCluster_FullMethodName = "/teleport.presence.v1.PresenceService/UpdateRemoteCluster"
	PresenceService_DeleteRemoteCluster_FullMethodName = "/teleport.presence.v1.PresenceService/DeleteRemoteCluster"
)

// PresenceServiceClient is the client API for PresenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PresenceService provides methods to manage presence of RemoteClusters
type PresenceServiceClient interface {
	// GetRemoteCluster retrieves a RemoteCluster by name.
	GetRemoteCluster(ctx context.Context, in *GetRemoteClusterRequest, opts ...grpc.CallOption) (*types.RemoteClusterV3, error)
	// ListRemoteClusters retrieves a page of RemoteClusters.
	ListRemoteClusters(ctx context.Context, in *ListRemoteClustersRequest, opts ...grpc.CallOption) (*ListRemoteClustersResponse, error)
	// UpdateRemoteCluster updates an existing RemoteCluster.
	UpdateRemoteCluster(ctx context.Context, in *UpdateRemoteClusterRequest, opts ...grpc.CallOption) (*types.RemoteClusterV3, error)
	// DeleteRemoteCluster removes an existing RemoteCluster by name.
	DeleteRemoteCluster(ctx context.Context, in *DeleteRemoteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type presenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceServiceClient(cc grpc.ClientConnInterface) PresenceServiceClient {
	return &presenceServiceClient{cc}
}

func (c *presenceServiceClient) GetRemoteCluster(ctx context.Context, in *GetRemoteClusterRequest, opts ...grpc.CallOption) (*types.RemoteClusterV3, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.RemoteClusterV3)
	err := c.cc.Invoke(ctx, PresenceService_GetRemoteCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceServiceClient) ListRemoteClusters(ctx context.Context, in *ListRemoteClustersRequest, opts ...grpc.CallOption) (*ListRemoteClustersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRemoteClustersResponse)
	err := c.cc.Invoke(ctx, PresenceService_ListRemoteClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceServiceClient) UpdateRemoteCluster(ctx context.Context, in *UpdateRemoteClusterRequest, opts ...grpc.CallOption) (*types.RemoteClusterV3, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(types.RemoteClusterV3)
	err := c.cc.Invoke(ctx, PresenceService_UpdateRemoteCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceServiceClient) DeleteRemoteCluster(ctx context.Context, in *DeleteRemoteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PresenceService_DeleteRemoteCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenceServiceServer is the server API for PresenceService service.
// All implementations must embed UnimplementedPresenceServiceServer
// for forward compatibility.
//
// PresenceService provides methods to manage presence of RemoteClusters
type PresenceServiceServer interface {
	// GetRemoteCluster retrieves a RemoteCluster by name.
	GetRemoteCluster(context.Context, *GetRemoteClusterRequest) (*types.RemoteClusterV3, error)
	// ListRemoteClusters retrieves a page of RemoteClusters.
	ListRemoteClusters(context.Context, *ListRemoteClustersRequest) (*ListRemoteClustersResponse, error)
	// UpdateRemoteCluster updates an existing RemoteCluster.
	UpdateRemoteCluster(context.Context, *UpdateRemoteClusterRequest) (*types.RemoteClusterV3, error)
	// DeleteRemoteCluster removes an existing RemoteCluster by name.
	DeleteRemoteCluster(context.Context, *DeleteRemoteClusterRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPresenceServiceServer()
}

// UnimplementedPresenceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPresenceServiceServer struct{}

func (UnimplementedPresenceServiceServer) GetRemoteCluster(context.Context, *GetRemoteClusterRequest) (*types.RemoteClusterV3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemoteCluster not implemented")
}
func (UnimplementedPresenceServiceServer) ListRemoteClusters(context.Context, *ListRemoteClustersRequest) (*ListRemoteClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemoteClusters not implemented")
}
func (UnimplementedPresenceServiceServer) UpdateRemoteCluster(context.Context, *UpdateRemoteClusterRequest) (*types.RemoteClusterV3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRemoteCluster not implemented")
}
func (UnimplementedPresenceServiceServer) DeleteRemoteCluster(context.Context, *DeleteRemoteClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRemoteCluster not implemented")
}
func (UnimplementedPresenceServiceServer) mustEmbedUnimplementedPresenceServiceServer() {}
func (UnimplementedPresenceServiceServer) testEmbeddedByValue()                         {}

// UnsafePresenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresenceServiceServer will
// result in compilation errors.
type UnsafePresenceServiceServer interface {
	mustEmbedUnimplementedPresenceServiceServer()
}

func RegisterPresenceServiceServer(s grpc.ServiceRegistrar, srv PresenceServiceServer) {
	// If the following call pancis, it indicates UnimplementedPresenceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PresenceService_ServiceDesc, srv)
}

func _PresenceService_GetRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).GetRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresenceService_GetRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).GetRemoteCluster(ctx, req.(*GetRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenceService_ListRemoteClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRemoteClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).ListRemoteClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresenceService_ListRemoteClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).ListRemoteClusters(ctx, req.(*ListRemoteClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenceService_UpdateRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).UpdateRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresenceService_UpdateRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).UpdateRemoteCluster(ctx, req.(*UpdateRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenceService_DeleteRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).DeleteRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresenceService_DeleteRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).DeleteRemoteCluster(ctx, req.(*DeleteRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PresenceService_ServiceDesc is the grpc.ServiceDesc for PresenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PresenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.presence.v1.PresenceService",
	HandlerType: (*PresenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRemoteCluster",
			Handler:    _PresenceService_GetRemoteCluster_Handler,
		},
		{
			MethodName: "ListRemoteClusters",
			Handler:    _PresenceService_ListRemoteClusters_Handler,
		},
		{
			MethodName: "UpdateRemoteCluster",
			Handler:    _PresenceService_UpdateRemoteCluster_Handler,
		},
		{
			MethodName: "DeleteRemoteCluster",
			Handler:    _PresenceService_DeleteRemoteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/presence/v1/service.proto",
}