// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chainrpc

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

// ChainKitClient is the client API for ChainKit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainKitClient interface {
	// lncli: `chain getblock`
	// GetBlock returns a block given the corresponding block hash.
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	// lncli: `chain getblockheader`
	// GetBlockHeader returns a block header with a particular block hash.
	GetBlockHeader(ctx context.Context, in *GetBlockHeaderRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error)
	// lncli: `chain getbestblock`
	// GetBestBlock returns the block hash and current height from the valid
	// most-work chain.
	GetBestBlock(ctx context.Context, in *GetBestBlockRequest, opts ...grpc.CallOption) (*GetBestBlockResponse, error)
	// lncli: `chain getblockhash`
	// GetBlockHash returns the hash of the block in the best blockchain
	// at the given height.
	GetBlockHash(ctx context.Context, in *GetBlockHashRequest, opts ...grpc.CallOption) (*GetBlockHashResponse, error)
}

type chainKitClient struct {
	cc grpc.ClientConnInterface
}

func NewChainKitClient(cc grpc.ClientConnInterface) ChainKitClient {
	return &chainKitClient{cc}
}

func (c *chainKitClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/chainrpc.ChainKit/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainKitClient) GetBlockHeader(ctx context.Context, in *GetBlockHeaderRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error) {
	out := new(GetBlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/chainrpc.ChainKit/GetBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainKitClient) GetBestBlock(ctx context.Context, in *GetBestBlockRequest, opts ...grpc.CallOption) (*GetBestBlockResponse, error) {
	out := new(GetBestBlockResponse)
	err := c.cc.Invoke(ctx, "/chainrpc.ChainKit/GetBestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainKitClient) GetBlockHash(ctx context.Context, in *GetBlockHashRequest, opts ...grpc.CallOption) (*GetBlockHashResponse, error) {
	out := new(GetBlockHashResponse)
	err := c.cc.Invoke(ctx, "/chainrpc.ChainKit/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainKitServer is the server API for ChainKit service.
// All implementations must embed UnimplementedChainKitServer
// for forward compatibility
type ChainKitServer interface {
	// lncli: `chain getblock`
	// GetBlock returns a block given the corresponding block hash.
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	// lncli: `chain getblockheader`
	// GetBlockHeader returns a block header with a particular block hash.
	GetBlockHeader(context.Context, *GetBlockHeaderRequest) (*GetBlockHeaderResponse, error)
	// lncli: `chain getbestblock`
	// GetBestBlock returns the block hash and current height from the valid
	// most-work chain.
	GetBestBlock(context.Context, *GetBestBlockRequest) (*GetBestBlockResponse, error)
	// lncli: `chain getblockhash`
	// GetBlockHash returns the hash of the block in the best blockchain
	// at the given height.
	GetBlockHash(context.Context, *GetBlockHashRequest) (*GetBlockHashResponse, error)
	mustEmbedUnimplementedChainKitServer()
}

// UnimplementedChainKitServer must be embedded to have forward compatible implementations.
type UnimplementedChainKitServer struct {
}

func (UnimplementedChainKitServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedChainKitServer) GetBlockHeader(context.Context, *GetBlockHeaderRequest) (*GetBlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeader not implemented")
}
func (UnimplementedChainKitServer) GetBestBlock(context.Context, *GetBestBlockRequest) (*GetBestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBestBlock not implemented")
}
func (UnimplementedChainKitServer) GetBlockHash(context.Context, *GetBlockHashRequest) (*GetBlockHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (UnimplementedChainKitServer) mustEmbedUnimplementedChainKitServer() {}

// UnsafeChainKitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainKitServer will
// result in compilation errors.
type UnsafeChainKitServer interface {
	mustEmbedUnimplementedChainKitServer()
}

func RegisterChainKitServer(s grpc.ServiceRegistrar, srv ChainKitServer) {
	s.RegisterService(&ChainKit_ServiceDesc, srv)
}

func _ChainKit_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainKitServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainrpc.ChainKit/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainKitServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainKit_GetBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainKitServer).GetBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainrpc.ChainKit/GetBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainKitServer).GetBlockHeader(ctx, req.(*GetBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainKit_GetBestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainKitServer).GetBestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainrpc.ChainKit/GetBestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainKitServer).GetBestBlock(ctx, req.(*GetBestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainKit_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainKitServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainrpc.ChainKit/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainKitServer).GetBlockHash(ctx, req.(*GetBlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChainKit_ServiceDesc is the grpc.ServiceDesc for ChainKit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainKit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainrpc.ChainKit",
	HandlerType: (*ChainKitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _ChainKit_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockHeader",
			Handler:    _ChainKit_GetBlockHeader_Handler,
		},
		{
			MethodName: "GetBestBlock",
			Handler:    _ChainKit_GetBestBlock_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _ChainKit_GetBlockHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chainrpc/chainkit.proto",
}