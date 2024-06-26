// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: shard/v2/shard.proto

package shardv2

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

// TransactionHandlerClient is the client API for TransactionHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionHandlerClient interface {
	// RegisterGameShard registers a game shard to be used in the Router system.
	RegisterGameShard(ctx context.Context, in *RegisterGameShardRequest, opts ...grpc.CallOption) (*RegisterGameShardResponse, error)
	// SubmitCardinalBatch handles receiving transactions from a game shard and persisting them to the chain.
	Submit(ctx context.Context, in *SubmitTransactionsRequest, opts ...grpc.CallOption) (*SubmitTransactionsResponse, error)
	// QueryTransactions queries the base shard for sequenced transactions.
	QueryTransactions(ctx context.Context, in *QueryTransactionsRequest, opts ...grpc.CallOption) (*QueryTransactionsResponse, error)
}

type transactionHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionHandlerClient(cc grpc.ClientConnInterface) TransactionHandlerClient {
	return &transactionHandlerClient{cc}
}

func (c *transactionHandlerClient) RegisterGameShard(ctx context.Context, in *RegisterGameShardRequest, opts ...grpc.CallOption) (*RegisterGameShardResponse, error) {
	out := new(RegisterGameShardResponse)
	err := c.cc.Invoke(ctx, "/world.engine.shard.v2.TransactionHandler/RegisterGameShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionHandlerClient) Submit(ctx context.Context, in *SubmitTransactionsRequest, opts ...grpc.CallOption) (*SubmitTransactionsResponse, error) {
	out := new(SubmitTransactionsResponse)
	err := c.cc.Invoke(ctx, "/world.engine.shard.v2.TransactionHandler/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionHandlerClient) QueryTransactions(ctx context.Context, in *QueryTransactionsRequest, opts ...grpc.CallOption) (*QueryTransactionsResponse, error) {
	out := new(QueryTransactionsResponse)
	err := c.cc.Invoke(ctx, "/world.engine.shard.v2.TransactionHandler/QueryTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionHandlerServer is the server API for TransactionHandler service.
// All implementations must embed UnimplementedTransactionHandlerServer
// for forward compatibility
type TransactionHandlerServer interface {
	// RegisterGameShard registers a game shard to be used in the Router system.
	RegisterGameShard(context.Context, *RegisterGameShardRequest) (*RegisterGameShardResponse, error)
	// SubmitCardinalBatch handles receiving transactions from a game shard and persisting them to the chain.
	Submit(context.Context, *SubmitTransactionsRequest) (*SubmitTransactionsResponse, error)
	// QueryTransactions queries the base shard for sequenced transactions.
	QueryTransactions(context.Context, *QueryTransactionsRequest) (*QueryTransactionsResponse, error)
	mustEmbedUnimplementedTransactionHandlerServer()
}

// UnimplementedTransactionHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionHandlerServer struct {
}

func (UnimplementedTransactionHandlerServer) RegisterGameShard(context.Context, *RegisterGameShardRequest) (*RegisterGameShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterGameShard not implemented")
}
func (UnimplementedTransactionHandlerServer) Submit(context.Context, *SubmitTransactionsRequest) (*SubmitTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedTransactionHandlerServer) QueryTransactions(context.Context, *QueryTransactionsRequest) (*QueryTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTransactions not implemented")
}
func (UnimplementedTransactionHandlerServer) mustEmbedUnimplementedTransactionHandlerServer() {}

// UnsafeTransactionHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionHandlerServer will
// result in compilation errors.
type UnsafeTransactionHandlerServer interface {
	mustEmbedUnimplementedTransactionHandlerServer()
}

func RegisterTransactionHandlerServer(s grpc.ServiceRegistrar, srv TransactionHandlerServer) {
	s.RegisterService(&TransactionHandler_ServiceDesc, srv)
}

func _TransactionHandler_RegisterGameShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterGameShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionHandlerServer).RegisterGameShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/world.engine.shard.v2.TransactionHandler/RegisterGameShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionHandlerServer).RegisterGameShard(ctx, req.(*RegisterGameShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionHandler_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionHandlerServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/world.engine.shard.v2.TransactionHandler/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionHandlerServer).Submit(ctx, req.(*SubmitTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionHandler_QueryTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionHandlerServer).QueryTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/world.engine.shard.v2.TransactionHandler/QueryTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionHandlerServer).QueryTransactions(ctx, req.(*QueryTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionHandler_ServiceDesc is the grpc.ServiceDesc for TransactionHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "world.engine.shard.v2.TransactionHandler",
	HandlerType: (*TransactionHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterGameShard",
			Handler:    _TransactionHandler_RegisterGameShard_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _TransactionHandler_Submit_Handler,
		},
		{
			MethodName: "QueryTransactions",
			Handler:    _TransactionHandler_QueryTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shard/v2/shard.proto",
}
