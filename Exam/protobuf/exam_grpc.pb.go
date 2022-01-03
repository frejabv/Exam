// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExamClient is the client API for Exam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamClient interface {
}

type examClient struct {
	cc grpc.ClientConnInterface
}

func NewExamClient(cc grpc.ClientConnInterface) ExamClient {
	return &examClient{cc}
}

// ExamServer is the server API for Exam service.
// All implementations must embed UnimplementedExamServer
// for forward compatibility
type ExamServer interface {
	mustEmbedUnimplementedExamServer()
}

// UnimplementedExamServer must be embedded to have forward compatible implementations.
type UnimplementedExamServer struct {
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

// Exam_ServiceDesc is the grpc.ServiceDesc for Exam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mock.Exam",
	HandlerType: (*ExamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "protobuf/exam.proto",
}
