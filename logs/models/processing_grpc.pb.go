// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: logs/processing.proto

package models

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

// LogProcessClient is the client API for LogProcess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogProcessClient interface {
	ProcessGzip(ctx context.Context, opts ...grpc.CallOption) (LogProcess_ProcessGzipClient, error)
}

type logProcessClient struct {
	cc grpc.ClientConnInterface
}

func NewLogProcessClient(cc grpc.ClientConnInterface) LogProcessClient {
	return &logProcessClient{cc}
}

func (c *logProcessClient) ProcessGzip(ctx context.Context, opts ...grpc.CallOption) (LogProcess_ProcessGzipClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogProcess_ServiceDesc.Streams[0], "/models.LogProcess/ProcessGzip", opts...)
	if err != nil {
		return nil, err
	}
	x := &logProcessProcessGzipClient{stream}
	return x, nil
}

type LogProcess_ProcessGzipClient interface {
	Send(*GzipChunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type logProcessProcessGzipClient struct {
	grpc.ClientStream
}

func (x *logProcessProcessGzipClient) Send(m *GzipChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logProcessProcessGzipClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogProcessServer is the server API for LogProcess service.
// All implementations must embed UnimplementedLogProcessServer
// for forward compatibility
type LogProcessServer interface {
	ProcessGzip(LogProcess_ProcessGzipServer) error
	mustEmbedUnimplementedLogProcessServer()
}

// UnimplementedLogProcessServer must be embedded to have forward compatible implementations.
type UnimplementedLogProcessServer struct {
}

func (UnimplementedLogProcessServer) ProcessGzip(LogProcess_ProcessGzipServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessGzip not implemented")
}
func (UnimplementedLogProcessServer) mustEmbedUnimplementedLogProcessServer() {}

// UnsafeLogProcessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogProcessServer will
// result in compilation errors.
type UnsafeLogProcessServer interface {
	mustEmbedUnimplementedLogProcessServer()
}

func RegisterLogProcessServer(s grpc.ServiceRegistrar, srv LogProcessServer) {
	s.RegisterService(&LogProcess_ServiceDesc, srv)
}

func _LogProcess_ProcessGzip_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogProcessServer).ProcessGzip(&logProcessProcessGzipServer{stream})
}

type LogProcess_ProcessGzipServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*GzipChunk, error)
	grpc.ServerStream
}

type logProcessProcessGzipServer struct {
	grpc.ServerStream
}

func (x *logProcessProcessGzipServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logProcessProcessGzipServer) Recv() (*GzipChunk, error) {
	m := new(GzipChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogProcess_ServiceDesc is the grpc.ServiceDesc for LogProcess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogProcess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.LogProcess",
	HandlerType: (*LogProcessServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessGzip",
			Handler:       _LogProcess_ProcessGzip_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "logs/processing.proto",
}
