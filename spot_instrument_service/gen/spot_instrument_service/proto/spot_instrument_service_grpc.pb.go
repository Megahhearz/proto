// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: spot_instrument_service/proto/spot_instrument_service.proto

package spot_instrument_service_proto

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
	SpotInstrumentService_ViewMarkets_FullMethodName = "/spot_instrument_service_proto.SpotInstrumentService/ViewMarkets"
)

// SpotInstrumentServiceClient is the client API for SpotInstrumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotInstrumentServiceClient interface {
	ViewMarkets(ctx context.Context, in *ViewMarketsRequest, opts ...grpc.CallOption) (*ViewMarketsResponse, error)
}

type spotInstrumentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotInstrumentServiceClient(cc grpc.ClientConnInterface) SpotInstrumentServiceClient {
	return &spotInstrumentServiceClient{cc}
}

func (c *spotInstrumentServiceClient) ViewMarkets(ctx context.Context, in *ViewMarketsRequest, opts ...grpc.CallOption) (*ViewMarketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ViewMarketsResponse)
	err := c.cc.Invoke(ctx, SpotInstrumentService_ViewMarkets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotInstrumentServiceServer is the server API for SpotInstrumentService service.
// All implementations must embed UnimplementedSpotInstrumentServiceServer
// for forward compatibility.
type SpotInstrumentServiceServer interface {
	ViewMarkets(context.Context, *ViewMarketsRequest) (*ViewMarketsResponse, error)
	mustEmbedUnimplementedSpotInstrumentServiceServer()
}

// UnimplementedSpotInstrumentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSpotInstrumentServiceServer struct{}

func (UnimplementedSpotInstrumentServiceServer) ViewMarkets(context.Context, *ViewMarketsRequest) (*ViewMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewMarkets not implemented")
}
func (UnimplementedSpotInstrumentServiceServer) mustEmbedUnimplementedSpotInstrumentServiceServer() {}
func (UnimplementedSpotInstrumentServiceServer) testEmbeddedByValue()                               {}

// UnsafeSpotInstrumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotInstrumentServiceServer will
// result in compilation errors.
type UnsafeSpotInstrumentServiceServer interface {
	mustEmbedUnimplementedSpotInstrumentServiceServer()
}

func RegisterSpotInstrumentServiceServer(s grpc.ServiceRegistrar, srv SpotInstrumentServiceServer) {
	// If the following call pancis, it indicates UnimplementedSpotInstrumentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SpotInstrumentService_ServiceDesc, srv)
}

func _SpotInstrumentService_ViewMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewMarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotInstrumentServiceServer).ViewMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpotInstrumentService_ViewMarkets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotInstrumentServiceServer).ViewMarkets(ctx, req.(*ViewMarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpotInstrumentService_ServiceDesc is the grpc.ServiceDesc for SpotInstrumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpotInstrumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spot_instrument_service_proto.SpotInstrumentService",
	HandlerType: (*SpotInstrumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewMarkets",
			Handler:    _SpotInstrumentService_ViewMarkets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spot_instrument_service/proto/spot_instrument_service.proto",
}
