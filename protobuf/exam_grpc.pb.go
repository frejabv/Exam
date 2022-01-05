// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// ExamClient is the client API for Exam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type examClient struct {
	cc grpc.ClientConnInterface
}

func NewExamClient(cc grpc.ClientConnInterface) ExamClient {
	return &examClient{cc}
}

func (c *examClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := c.cc.Invoke(ctx, "/exam.Exam/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/exam.Exam/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/exam.Exam/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExamServer is the server API for Exam service.
// All implementations must embed UnimplementedExamServer
// for forward compatibility
type ExamServer interface {
	Put(context.Context, *PutRequest) (*PutReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedExamServer()
}

// UnimplementedExamServer must be embedded to have forward compatible implementations.
type UnimplementedExamServer struct {
}

func (UnimplementedExamServer) Put(context.Context, *PutRequest) (*PutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedExamServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExamServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedExamServer) mustEmbedUnimplementedExamServer() {}

// UnsafeExamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExamServer will
// result in compilation errors.
type UnsafeExamServer interface {
	mustEmbedUnimplementedExamServer()
}

func RegisterExamServer(s grpc.ServiceRegistrar, srv ExamServer) {
	s.RegisterService(&Exam_ServiceDesc, srv)
}

func _Exam_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exam.Exam/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exam.Exam/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exam.Exam/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Exam_ServiceDesc is the grpc.ServiceDesc for Exam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exam.Exam",
	HandlerType: (*ExamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Exam_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Exam_Get_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Exam_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/exam.proto",
}
