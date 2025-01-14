// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/blog/v1/task.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskGrpcService_ListDayTask_FullMethodName    = "/blog.v1.TaskGrpcService/ListDayTask"
	TaskGrpcService_PushDayTask_FullMethodName    = "/blog.v1.TaskGrpcService/PushDayTask"
	TaskGrpcService_RecvDayTask_FullMethodName    = "/blog.v1.TaskGrpcService/RecvDayTask"
	TaskGrpcService_ListMasterTask_FullMethodName = "/blog.v1.TaskGrpcService/ListMasterTask"
	TaskGrpcService_PushMasterTask_FullMethodName = "/blog.v1.TaskGrpcService/PushMasterTask"
)

// TaskGrpcServiceClient is the client API for TaskGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskGrpcServiceClient interface {
	ListDayTask(ctx context.Context, in *ListDayTaskRequest, opts ...grpc.CallOption) (*ListDayTaskReply, error)
	PushDayTask(ctx context.Context, in *PushDayTaskRequest, opts ...grpc.CallOption) (*Empty, error)
	RecvDayTask(ctx context.Context, in *RecvDayTaskRequest, opts ...grpc.CallOption) (*Empty, error)
	ListMasterTask(ctx context.Context, in *ListMasterTaskRequest, opts ...grpc.CallOption) (*ListMasterTaskReply, error)
	PushMasterTask(ctx context.Context, in *PushMasterTaskRequest, opts ...grpc.CallOption) (*Empty, error)
}

type taskGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskGrpcServiceClient(cc grpc.ClientConnInterface) TaskGrpcServiceClient {
	return &taskGrpcServiceClient{cc}
}

func (c *taskGrpcServiceClient) ListDayTask(ctx context.Context, in *ListDayTaskRequest, opts ...grpc.CallOption) (*ListDayTaskReply, error) {
	out := new(ListDayTaskReply)
	err := c.cc.Invoke(ctx, TaskGrpcService_ListDayTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskGrpcServiceClient) PushDayTask(ctx context.Context, in *PushDayTaskRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TaskGrpcService_PushDayTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskGrpcServiceClient) RecvDayTask(ctx context.Context, in *RecvDayTaskRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TaskGrpcService_RecvDayTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskGrpcServiceClient) ListMasterTask(ctx context.Context, in *ListMasterTaskRequest, opts ...grpc.CallOption) (*ListMasterTaskReply, error) {
	out := new(ListMasterTaskReply)
	err := c.cc.Invoke(ctx, TaskGrpcService_ListMasterTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskGrpcServiceClient) PushMasterTask(ctx context.Context, in *PushMasterTaskRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TaskGrpcService_PushMasterTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskGrpcServiceServer is the server API for TaskGrpcService service.
// All implementations must embed UnimplementedTaskGrpcServiceServer
// for forward compatibility
type TaskGrpcServiceServer interface {
	ListDayTask(context.Context, *ListDayTaskRequest) (*ListDayTaskReply, error)
	PushDayTask(context.Context, *PushDayTaskRequest) (*Empty, error)
	RecvDayTask(context.Context, *RecvDayTaskRequest) (*Empty, error)
	ListMasterTask(context.Context, *ListMasterTaskRequest) (*ListMasterTaskReply, error)
	PushMasterTask(context.Context, *PushMasterTaskRequest) (*Empty, error)
	mustEmbedUnimplementedTaskGrpcServiceServer()
}

// UnimplementedTaskGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskGrpcServiceServer struct {
}

func (UnimplementedTaskGrpcServiceServer) ListDayTask(context.Context, *ListDayTaskRequest) (*ListDayTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDayTask not implemented")
}
func (UnimplementedTaskGrpcServiceServer) PushDayTask(context.Context, *PushDayTaskRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushDayTask not implemented")
}
func (UnimplementedTaskGrpcServiceServer) RecvDayTask(context.Context, *RecvDayTaskRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecvDayTask not implemented")
}
func (UnimplementedTaskGrpcServiceServer) ListMasterTask(context.Context, *ListMasterTaskRequest) (*ListMasterTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMasterTask not implemented")
}
func (UnimplementedTaskGrpcServiceServer) PushMasterTask(context.Context, *PushMasterTaskRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMasterTask not implemented")
}
func (UnimplementedTaskGrpcServiceServer) mustEmbedUnimplementedTaskGrpcServiceServer() {}

// UnsafeTaskGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskGrpcServiceServer will
// result in compilation errors.
type UnsafeTaskGrpcServiceServer interface {
	mustEmbedUnimplementedTaskGrpcServiceServer()
}

func RegisterTaskGrpcServiceServer(s grpc.ServiceRegistrar, srv TaskGrpcServiceServer) {
	s.RegisterService(&TaskGrpcService_ServiceDesc, srv)
}

func _TaskGrpcService_ListDayTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDayTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskGrpcServiceServer).ListDayTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskGrpcService_ListDayTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskGrpcServiceServer).ListDayTask(ctx, req.(*ListDayTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskGrpcService_PushDayTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushDayTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskGrpcServiceServer).PushDayTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskGrpcService_PushDayTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskGrpcServiceServer).PushDayTask(ctx, req.(*PushDayTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskGrpcService_RecvDayTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecvDayTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskGrpcServiceServer).RecvDayTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskGrpcService_RecvDayTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskGrpcServiceServer).RecvDayTask(ctx, req.(*RecvDayTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskGrpcService_ListMasterTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMasterTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskGrpcServiceServer).ListMasterTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskGrpcService_ListMasterTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskGrpcServiceServer).ListMasterTask(ctx, req.(*ListMasterTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskGrpcService_PushMasterTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMasterTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskGrpcServiceServer).PushMasterTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskGrpcService_PushMasterTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskGrpcServiceServer).PushMasterTask(ctx, req.(*PushMasterTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskGrpcService_ServiceDesc is the grpc.ServiceDesc for TaskGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.v1.TaskGrpcService",
	HandlerType: (*TaskGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDayTask",
			Handler:    _TaskGrpcService_ListDayTask_Handler,
		},
		{
			MethodName: "PushDayTask",
			Handler:    _TaskGrpcService_PushDayTask_Handler,
		},
		{
			MethodName: "RecvDayTask",
			Handler:    _TaskGrpcService_RecvDayTask_Handler,
		},
		{
			MethodName: "ListMasterTask",
			Handler:    _TaskGrpcService_ListMasterTask_Handler,
		},
		{
			MethodName: "PushMasterTask",
			Handler:    _TaskGrpcService_PushMasterTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog/v1/task.proto",
}
