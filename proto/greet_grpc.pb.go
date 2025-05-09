// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/greet.proto

// Defines the package name for this protocol buffer file

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GreetService_SayHello_FullMethodName                    = "/greet_service.GreetService/SayHello"
	GreetService_SayHelloServerSteam_FullMethodName         = "/greet_service.GreetService/SayHelloServerSteam"
	GreetService_SayHelloClientStream_FullMethodName        = "/greet_service.GreetService/SayHelloClientStream"
	GreetService_SayHelloBidirectionalStream_FullMethodName = "/greet_service.GreetService/SayHelloBidirectionalStream"
)

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for handling different types of greeting requests
type GreetServiceClient interface {
	// A simple unary RPC where the client sends a request, and the server responds with a greeting message
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	// A server streaming RPC where the client sends a list of names,
	// and the server responds with multiple greeting messages in a stream
	SayHelloServerSteam(ctx context.Context, in *NameList, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
	// A client streaming RPC where the client sends multiple names in a stream,
	// and the server responds with a single message containing all greetings
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, MessageList], error)
	// A bidirectional streaming RPC where both client and server exchange multiple messages in real-time
	SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreetService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayHelloServerSteam(ctx context.Context, in *NameList, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], GreetService_SayHelloServerSteam_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NameList, HelloResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloServerSteamClient = grpc.ServerStreamingClient[HelloResponse]

func (c *greetServiceClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, MessageList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], GreetService_SayHelloClientStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, MessageList]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloClientStreamClient = grpc.ClientStreamingClient[HelloRequest, MessageList]

func (c *greetServiceClient) SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], GreetService_SayHelloBidirectionalStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloBidirectionalStreamClient = grpc.BidiStreamingClient[HelloRequest, HelloResponse]

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility.
//
// Service definition for handling different types of greeting requests
type GreetServiceServer interface {
	// A simple unary RPC where the client sends a request, and the server responds with a greeting message
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	// A server streaming RPC where the client sends a list of names,
	// and the server responds with multiple greeting messages in a stream
	SayHelloServerSteam(*NameList, grpc.ServerStreamingServer[HelloResponse]) error
	// A client streaming RPC where the client sends multiple names in a stream,
	// and the server responds with a single message containing all greetings
	SayHelloClientStream(grpc.ClientStreamingServer[HelloRequest, MessageList]) error
	// A bidirectional streaming RPC where both client and server exchange multiple messages in real-time
	SayHelloBidirectionalStream(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreetServiceServer struct{}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloServerSteam(*NameList, grpc.ServerStreamingServer[HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerSteam not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloClientStream(grpc.ClientStreamingServer[HelloRequest, MessageList]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloBidirectionalStream(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStream not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}
func (UnimplementedGreetServiceServer) testEmbeddedByValue()                      {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	// If the following call pancis, it indicates UnimplementedGreetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayHelloServerSteam_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NameList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayHelloServerSteam(m, &grpc.GenericServerStream[NameList, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloServerSteamServer = grpc.ServerStreamingServer[HelloResponse]

func _GreetService_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloClientStream(&grpc.GenericServerStream[HelloRequest, MessageList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloClientStreamServer = grpc.ClientStreamingServer[HelloRequest, MessageList]

func _GreetService_SayHelloBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloBidirectionalStream(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetService_SayHelloBidirectionalStreamServer = grpc.BidiStreamingServer[HelloRequest, HelloResponse]

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerSteam",
			Handler:       _GreetService_SayHelloServerSteam_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _GreetService_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStream",
			Handler:       _GreetService_SayHelloBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}
