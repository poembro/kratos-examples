// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: api/blog/v1/task.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTaskGrpcServiceListDayTask = "/blog.v1.TaskGrpcService/ListDayTask"
const OperationTaskGrpcServiceListMasterTask = "/blog.v1.TaskGrpcService/ListMasterTask"
const OperationTaskGrpcServicePushDayTask = "/blog.v1.TaskGrpcService/PushDayTask"
const OperationTaskGrpcServicePushMasterTask = "/blog.v1.TaskGrpcService/PushMasterTask"
const OperationTaskGrpcServiceRecvDayTask = "/blog.v1.TaskGrpcService/RecvDayTask"

type TaskGrpcServiceHTTPServer interface {
	ListDayTask(context.Context, *ListDayTaskRequest) (*ListDayTaskReply, error)
	ListMasterTask(context.Context, *ListMasterTaskRequest) (*ListMasterTaskReply, error)
	PushDayTask(context.Context, *PushDayTaskRequest) (*Empty, error)
	PushMasterTask(context.Context, *PushMasterTaskRequest) (*Empty, error)
	RecvDayTask(context.Context, *RecvDayTaskRequest) (*Empty, error)
}

func RegisterTaskGrpcServiceHTTPServer(s *http.Server, srv TaskGrpcServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/task/day/list", _TaskGrpcService_ListDayTask0_HTTP_Handler(srv))
	r.POST("/v1/task/day/push", _TaskGrpcService_PushDayTask0_HTTP_Handler(srv))
	r.POST("/v1/task/day/recv", _TaskGrpcService_RecvDayTask0_HTTP_Handler(srv))
	r.GET("/v1/task/master/list", _TaskGrpcService_ListMasterTask0_HTTP_Handler(srv))
	r.POST("/v1/task/master/push", _TaskGrpcService_PushMasterTask0_HTTP_Handler(srv))
}

func _TaskGrpcService_ListDayTask0_HTTP_Handler(srv TaskGrpcServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDayTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGrpcServiceListDayTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDayTask(ctx, req.(*ListDayTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDayTaskReply)
		return ctx.Result(200, reply)
	}
}

func _TaskGrpcService_PushDayTask0_HTTP_Handler(srv TaskGrpcServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushDayTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGrpcServicePushDayTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushDayTask(ctx, req.(*PushDayTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _TaskGrpcService_RecvDayTask0_HTTP_Handler(srv TaskGrpcServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RecvDayTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGrpcServiceRecvDayTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecvDayTask(ctx, req.(*RecvDayTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _TaskGrpcService_ListMasterTask0_HTTP_Handler(srv TaskGrpcServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMasterTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGrpcServiceListMasterTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMasterTask(ctx, req.(*ListMasterTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMasterTaskReply)
		return ctx.Result(200, reply)
	}
}

func _TaskGrpcService_PushMasterTask0_HTTP_Handler(srv TaskGrpcServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushMasterTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGrpcServicePushMasterTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushMasterTask(ctx, req.(*PushMasterTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type TaskGrpcServiceHTTPClient interface {
	ListDayTask(ctx context.Context, req *ListDayTaskRequest, opts ...http.CallOption) (rsp *ListDayTaskReply, err error)
	ListMasterTask(ctx context.Context, req *ListMasterTaskRequest, opts ...http.CallOption) (rsp *ListMasterTaskReply, err error)
	PushDayTask(ctx context.Context, req *PushDayTaskRequest, opts ...http.CallOption) (rsp *Empty, err error)
	PushMasterTask(ctx context.Context, req *PushMasterTaskRequest, opts ...http.CallOption) (rsp *Empty, err error)
	RecvDayTask(ctx context.Context, req *RecvDayTaskRequest, opts ...http.CallOption) (rsp *Empty, err error)
}

type TaskGrpcServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskGrpcServiceHTTPClient(client *http.Client) TaskGrpcServiceHTTPClient {
	return &TaskGrpcServiceHTTPClientImpl{client}
}

func (c *TaskGrpcServiceHTTPClientImpl) ListDayTask(ctx context.Context, in *ListDayTaskRequest, opts ...http.CallOption) (*ListDayTaskReply, error) {
	var out ListDayTaskReply
	pattern := "/v1/task/day/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGrpcServiceListDayTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskGrpcServiceHTTPClientImpl) ListMasterTask(ctx context.Context, in *ListMasterTaskRequest, opts ...http.CallOption) (*ListMasterTaskReply, error) {
	var out ListMasterTaskReply
	pattern := "/v1/task/master/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGrpcServiceListMasterTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskGrpcServiceHTTPClientImpl) PushDayTask(ctx context.Context, in *PushDayTaskRequest, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/v1/task/day/push"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskGrpcServicePushDayTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskGrpcServiceHTTPClientImpl) PushMasterTask(ctx context.Context, in *PushMasterTaskRequest, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/v1/task/master/push"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskGrpcServicePushMasterTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskGrpcServiceHTTPClientImpl) RecvDayTask(ctx context.Context, in *RecvDayTaskRequest, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/v1/task/day/recv"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskGrpcServiceRecvDayTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}