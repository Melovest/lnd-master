// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package invoicesrpc

import (
	context "context"
	lnrpc "github.com/lightningnetwork/lnd/lnrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InvoicesClient is the client API for Invoices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicesClient interface {
	// SubscribeSingleInvoice returns a uni-directional stream (server -> client)
	// to notify the client of state transitions of the specified invoice.
	// Initially the current invoice state is always sent out.
	SubscribeSingleInvoice(ctx context.Context, in *SubscribeSingleInvoiceRequest, opts ...grpc.CallOption) (Invoices_SubscribeSingleInvoiceClient, error)
	// lncli: `cancelinvoice`
	// CancelInvoice cancels a currently open invoice. If the invoice is already
	// canceled, this call will succeed. If the invoice is already settled, it will
	// fail.
	CancelInvoice(ctx context.Context, in *CancelInvoiceMsg, opts ...grpc.CallOption) (*CancelInvoiceResp, error)
	// lncli: `addholdinvoice`
	// AddHoldInvoice creates a hold invoice. It ties the invoice to the hash
	// supplied in the request.
	AddHoldInvoice(ctx context.Context, in *AddHoldInvoiceRequest, opts ...grpc.CallOption) (*AddHoldInvoiceResp, error)
	// lncli: `settleinvoice`
	// SettleInvoice settles an accepted invoice. If the invoice is already
	// settled, this call will succeed.
	SettleInvoice(ctx context.Context, in *SettleInvoiceMsg, opts ...grpc.CallOption) (*SettleInvoiceResp, error)
	// LookupInvoiceV2 attempts to look up at invoice. An invoice can be referenced
	// using either its payment hash, payment address, or set ID.
	LookupInvoiceV2(ctx context.Context, in *LookupInvoiceMsg, opts ...grpc.CallOption) (*lnrpc.Invoice, error)
	// HtlcModifier is a bidirectional streaming RPC that allows a client to
	// intercept and modify the HTLCs that attempt to settle the given invoice. The
	// server will send HTLCs of invoices to the client and the client can modify
	// some aspects of the HTLC in order to pass the invoice acceptance tests.
	HtlcModifier(ctx context.Context, opts ...grpc.CallOption) (Invoices_HtlcModifierClient, error)
}

type invoicesClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicesClient(cc grpc.ClientConnInterface) InvoicesClient {
	return &invoicesClient{cc}
}

func (c *invoicesClient) SubscribeSingleInvoice(ctx context.Context, in *SubscribeSingleInvoiceRequest, opts ...grpc.CallOption) (Invoices_SubscribeSingleInvoiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Invoices_ServiceDesc.Streams[0], "/invoicesrpc.Invoices/SubscribeSingleInvoice", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicesSubscribeSingleInvoiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Invoices_SubscribeSingleInvoiceClient interface {
	Recv() (*lnrpc.Invoice, error)
	grpc.ClientStream
}

type invoicesSubscribeSingleInvoiceClient struct {
	grpc.ClientStream
}

func (x *invoicesSubscribeSingleInvoiceClient) Recv() (*lnrpc.Invoice, error) {
	m := new(lnrpc.Invoice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *invoicesClient) CancelInvoice(ctx context.Context, in *CancelInvoiceMsg, opts ...grpc.CallOption) (*CancelInvoiceResp, error) {
	out := new(CancelInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/CancelInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) AddHoldInvoice(ctx context.Context, in *AddHoldInvoiceRequest, opts ...grpc.CallOption) (*AddHoldInvoiceResp, error) {
	out := new(AddHoldInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/AddHoldInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) SettleInvoice(ctx context.Context, in *SettleInvoiceMsg, opts ...grpc.CallOption) (*SettleInvoiceResp, error) {
	out := new(SettleInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/SettleInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) LookupInvoiceV2(ctx context.Context, in *LookupInvoiceMsg, opts ...grpc.CallOption) (*lnrpc.Invoice, error) {
	out := new(lnrpc.Invoice)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/LookupInvoiceV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) HtlcModifier(ctx context.Context, opts ...grpc.CallOption) (Invoices_HtlcModifierClient, error) {
	stream, err := c.cc.NewStream(ctx, &Invoices_ServiceDesc.Streams[1], "/invoicesrpc.Invoices/HtlcModifier", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicesHtlcModifierClient{stream}
	return x, nil
}

type Invoices_HtlcModifierClient interface {
	Send(*HtlcModifyResponse) error
	Recv() (*HtlcModifyRequest, error)
	grpc.ClientStream
}

type invoicesHtlcModifierClient struct {
	grpc.ClientStream
}

func (x *invoicesHtlcModifierClient) Send(m *HtlcModifyResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *invoicesHtlcModifierClient) Recv() (*HtlcModifyRequest, error) {
	m := new(HtlcModifyRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InvoicesServer is the server API for Invoices service.
// All implementations must embed UnimplementedInvoicesServer
// for forward compatibility
type InvoicesServer interface {
	// SubscribeSingleInvoice returns a uni-directional stream (server -> client)
	// to notify the client of state transitions of the specified invoice.
	// Initially the current invoice state is always sent out.
	SubscribeSingleInvoice(*SubscribeSingleInvoiceRequest, Invoices_SubscribeSingleInvoiceServer) error
	// lncli: `cancelinvoice`
	// CancelInvoice cancels a currently open invoice. If the invoice is already
	// canceled, this call will succeed. If the invoice is already settled, it will
	// fail.
	CancelInvoice(context.Context, *CancelInvoiceMsg) (*CancelInvoiceResp, error)
	// lncli: `addholdinvoice`
	// AddHoldInvoice creates a hold invoice. It ties the invoice to the hash
	// supplied in the request.
	AddHoldInvoice(context.Context, *AddHoldInvoiceRequest) (*AddHoldInvoiceResp, error)
	// lncli: `settleinvoice`
	// SettleInvoice settles an accepted invoice. If the invoice is already
	// settled, this call will succeed.
	SettleInvoice(context.Context, *SettleInvoiceMsg) (*SettleInvoiceResp, error)
	// LookupInvoiceV2 attempts to look up at invoice. An invoice can be referenced
	// using either its payment hash, payment address, or set ID.
	LookupInvoiceV2(context.Context, *LookupInvoiceMsg) (*lnrpc.Invoice, error)
	// HtlcModifier is a bidirectional streaming RPC that allows a client to
	// intercept and modify the HTLCs that attempt to settle the given invoice. The
	// server will send HTLCs of invoices to the client and the client can modify
	// some aspects of the HTLC in order to pass the invoice acceptance tests.
	HtlcModifier(Invoices_HtlcModifierServer) error
	mustEmbedUnimplementedInvoicesServer()
}

// UnimplementedInvoicesServer must be embedded to have forward compatible implementations.
type UnimplementedInvoicesServer struct {
}

func (UnimplementedInvoicesServer) SubscribeSingleInvoice(*SubscribeSingleInvoiceRequest, Invoices_SubscribeSingleInvoiceServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeSingleInvoice not implemented")
}
func (UnimplementedInvoicesServer) CancelInvoice(context.Context, *CancelInvoiceMsg) (*CancelInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelInvoice not implemented")
}
func (UnimplementedInvoicesServer) AddHoldInvoice(context.Context, *AddHoldInvoiceRequest) (*AddHoldInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHoldInvoice not implemented")
}
func (UnimplementedInvoicesServer) SettleInvoice(context.Context, *SettleInvoiceMsg) (*SettleInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettleInvoice not implemented")
}
func (UnimplementedInvoicesServer) LookupInvoiceV2(context.Context, *LookupInvoiceMsg) (*lnrpc.Invoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupInvoiceV2 not implemented")
}
func (UnimplementedInvoicesServer) HtlcModifier(Invoices_HtlcModifierServer) error {
	return status.Errorf(codes.Unimplemented, "method HtlcModifier not implemented")
}
func (UnimplementedInvoicesServer) mustEmbedUnimplementedInvoicesServer() {}

// UnsafeInvoicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicesServer will
// result in compilation errors.
type UnsafeInvoicesServer interface {
	mustEmbedUnimplementedInvoicesServer()
}

func RegisterInvoicesServer(s grpc.ServiceRegistrar, srv InvoicesServer) {
	s.RegisterService(&Invoices_ServiceDesc, srv)
}

func _Invoices_SubscribeSingleInvoice_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeSingleInvoiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InvoicesServer).SubscribeSingleInvoice(m, &invoicesSubscribeSingleInvoiceServer{stream})
}

type Invoices_SubscribeSingleInvoiceServer interface {
	Send(*lnrpc.Invoice) error
	grpc.ServerStream
}

type invoicesSubscribeSingleInvoiceServer struct {
	grpc.ServerStream
}

func (x *invoicesSubscribeSingleInvoiceServer) Send(m *lnrpc.Invoice) error {
	return x.ServerStream.SendMsg(m)
}

func _Invoices_CancelInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelInvoiceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).CancelInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/CancelInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).CancelInvoice(ctx, req.(*CancelInvoiceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_AddHoldInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHoldInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).AddHoldInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/AddHoldInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).AddHoldInvoice(ctx, req.(*AddHoldInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_SettleInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettleInvoiceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).SettleInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/SettleInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).SettleInvoice(ctx, req.(*SettleInvoiceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_LookupInvoiceV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupInvoiceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).LookupInvoiceV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/LookupInvoiceV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).LookupInvoiceV2(ctx, req.(*LookupInvoiceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_HtlcModifier_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InvoicesServer).HtlcModifier(&invoicesHtlcModifierServer{stream})
}

type Invoices_HtlcModifierServer interface {
	Send(*HtlcModifyRequest) error
	Recv() (*HtlcModifyResponse, error)
	grpc.ServerStream
}

type invoicesHtlcModifierServer struct {
	grpc.ServerStream
}

func (x *invoicesHtlcModifierServer) Send(m *HtlcModifyRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *invoicesHtlcModifierServer) Recv() (*HtlcModifyResponse, error) {
	m := new(HtlcModifyResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Invoices_ServiceDesc is the grpc.ServiceDesc for Invoices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "invoicesrpc.Invoices",
	HandlerType: (*InvoicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelInvoice",
			Handler:    _Invoices_CancelInvoice_Handler,
		},
		{
			MethodName: "AddHoldInvoice",
			Handler:    _Invoices_AddHoldInvoice_Handler,
		},
		{
			MethodName: "SettleInvoice",
			Handler:    _Invoices_SettleInvoice_Handler,
		},
		{
			MethodName: "LookupInvoiceV2",
			Handler:    _Invoices_LookupInvoiceV2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeSingleInvoice",
			Handler:       _Invoices_SubscribeSingleInvoice_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HtlcModifier",
			Handler:       _Invoices_HtlcModifier_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "invoicesrpc/invoices.proto",
}