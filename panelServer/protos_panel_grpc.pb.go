// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package panelServer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// InstrumentClient is the client API for Instrument service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstrumentClient interface {
	SearchInstrumentInfo(ctx context.Context, in *SearchInstrumentInfoQuery, opts ...grpc.CallOption) (*SearchInstrumentInfoResponse, error)
	GetInstrumentSnapshot(ctx context.Context, in *InstrumentSnapshotQuery, opts ...grpc.CallOption) (Instrument_GetInstrumentSnapshotClient, error)
	GetInstrumentHistoricalPrice(ctx context.Context, in *HistoricalPriceQuery, opts ...grpc.CallOption) (*HistoricalPriceResponse, error)
}

type instrumentClient struct {
	cc grpc.ClientConnInterface
}

func NewInstrumentClient(cc grpc.ClientConnInterface) InstrumentClient {
	return &instrumentClient{cc}
}

func (c *instrumentClient) SearchInstrumentInfo(ctx context.Context, in *SearchInstrumentInfoQuery, opts ...grpc.CallOption) (*SearchInstrumentInfoResponse, error) {
	out := new(SearchInstrumentInfoResponse)
	err := c.cc.Invoke(ctx, "/panelServer.Instrument/SearchInstrumentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentClient) GetInstrumentSnapshot(ctx context.Context, in *InstrumentSnapshotQuery, opts ...grpc.CallOption) (Instrument_GetInstrumentSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &Instrument_ServiceDesc.Streams[0], "/panelServer.Instrument/GetInstrumentSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &instrumentGetInstrumentSnapshotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Instrument_GetInstrumentSnapshotClient interface {
	Recv() (*InstrumentSnapshotResponse, error)
	grpc.ClientStream
}

type instrumentGetInstrumentSnapshotClient struct {
	grpc.ClientStream
}

func (x *instrumentGetInstrumentSnapshotClient) Recv() (*InstrumentSnapshotResponse, error) {
	m := new(InstrumentSnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *instrumentClient) GetInstrumentHistoricalPrice(ctx context.Context, in *HistoricalPriceQuery, opts ...grpc.CallOption) (*HistoricalPriceResponse, error) {
	out := new(HistoricalPriceResponse)
	err := c.cc.Invoke(ctx, "/panelServer.Instrument/GetInstrumentHistoricalPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstrumentServer is the server API for Instrument service.
// All implementations must embed UnimplementedInstrumentServer
// for forward compatibility
type InstrumentServer interface {
	SearchInstrumentInfo(context.Context, *SearchInstrumentInfoQuery) (*SearchInstrumentInfoResponse, error)
	GetInstrumentSnapshot(*InstrumentSnapshotQuery, Instrument_GetInstrumentSnapshotServer) error
	GetInstrumentHistoricalPrice(context.Context, *HistoricalPriceQuery) (*HistoricalPriceResponse, error)
	mustEmbedUnimplementedInstrumentServer()
}

// UnimplementedInstrumentServer must be embedded to have forward compatible implementations.
type UnimplementedInstrumentServer struct {
}

func (UnimplementedInstrumentServer) SearchInstrumentInfo(context.Context, *SearchInstrumentInfoQuery) (*SearchInstrumentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchInstrumentInfo not implemented")
}
func (UnimplementedInstrumentServer) GetInstrumentSnapshot(*InstrumentSnapshotQuery, Instrument_GetInstrumentSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method GetInstrumentSnapshot not implemented")
}
func (UnimplementedInstrumentServer) GetInstrumentHistoricalPrice(context.Context, *HistoricalPriceQuery) (*HistoricalPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentHistoricalPrice not implemented")
}
func (UnimplementedInstrumentServer) mustEmbedUnimplementedInstrumentServer() {}

// UnsafeInstrumentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstrumentServer will
// result in compilation errors.
type UnsafeInstrumentServer interface {
	mustEmbedUnimplementedInstrumentServer()
}

func RegisterInstrumentServer(s grpc.ServiceRegistrar, srv InstrumentServer) {
	s.RegisterService(&Instrument_ServiceDesc, srv)
}

func _Instrument_SearchInstrumentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchInstrumentInfoQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServer).SearchInstrumentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panelServer.Instrument/SearchInstrumentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServer).SearchInstrumentInfo(ctx, req.(*SearchInstrumentInfoQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instrument_GetInstrumentSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstrumentSnapshotQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstrumentServer).GetInstrumentSnapshot(m, &instrumentGetInstrumentSnapshotServer{stream})
}

type Instrument_GetInstrumentSnapshotServer interface {
	Send(*InstrumentSnapshotResponse) error
	grpc.ServerStream
}

type instrumentGetInstrumentSnapshotServer struct {
	grpc.ServerStream
}

func (x *instrumentGetInstrumentSnapshotServer) Send(m *InstrumentSnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Instrument_GetInstrumentHistoricalPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoricalPriceQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServer).GetInstrumentHistoricalPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panelServer.Instrument/GetInstrumentHistoricalPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServer).GetInstrumentHistoricalPrice(ctx, req.(*HistoricalPriceQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Instrument_ServiceDesc is the grpc.ServiceDesc for Instrument service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Instrument_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "panelServer.Instrument",
	HandlerType: (*InstrumentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchInstrumentInfo",
			Handler:    _Instrument_SearchInstrumentInfo_Handler,
		},
		{
			MethodName: "GetInstrumentHistoricalPrice",
			Handler:    _Instrument_GetInstrumentHistoricalPrice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetInstrumentSnapshot",
			Handler:       _Instrument_GetInstrumentSnapshot_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos_panel.proto",
}

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexClient interface {
	GetIndexHistoricalPrice(ctx context.Context, in *HistoricalPriceQuery, opts ...grpc.CallOption) (*IndexHistoricalPriceResponse, error)
}

type indexClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexClient(cc grpc.ClientConnInterface) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) GetIndexHistoricalPrice(ctx context.Context, in *HistoricalPriceQuery, opts ...grpc.CallOption) (*IndexHistoricalPriceResponse, error) {
	out := new(IndexHistoricalPriceResponse)
	err := c.cc.Invoke(ctx, "/panelServer.Index/GetIndexHistoricalPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
// All implementations must embed UnimplementedIndexServer
// for forward compatibility
type IndexServer interface {
	GetIndexHistoricalPrice(context.Context, *HistoricalPriceQuery) (*IndexHistoricalPriceResponse, error)
	mustEmbedUnimplementedIndexServer()
}

// UnimplementedIndexServer must be embedded to have forward compatible implementations.
type UnimplementedIndexServer struct {
}

func (UnimplementedIndexServer) GetIndexHistoricalPrice(context.Context, *HistoricalPriceQuery) (*IndexHistoricalPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexHistoricalPrice not implemented")
}
func (UnimplementedIndexServer) mustEmbedUnimplementedIndexServer() {}

// UnsafeIndexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServer will
// result in compilation errors.
type UnsafeIndexServer interface {
	mustEmbedUnimplementedIndexServer()
}

func RegisterIndexServer(s grpc.ServiceRegistrar, srv IndexServer) {
	s.RegisterService(&Index_ServiceDesc, srv)
}

func _Index_GetIndexHistoricalPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoricalPriceQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetIndexHistoricalPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panelServer.Index/GetIndexHistoricalPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetIndexHistoricalPrice(ctx, req.(*HistoricalPriceQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Index_ServiceDesc is the grpc.ServiceDesc for Index service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Index_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "panelServer.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIndexHistoricalPrice",
			Handler:    _Index_GetIndexHistoricalPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos_panel.proto",
}
