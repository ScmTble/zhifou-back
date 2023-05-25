// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: dynamic.proto

package pb

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

const (
	Dynamic_InsertPost_FullMethodName = "/dynamic.Dynamic/InsertPost"
	Dynamic_GetDynamic_FullMethodName = "/dynamic.Dynamic/GetDynamic"
	Dynamic_FindPost_FullMethodName   = "/dynamic.Dynamic/FindPost"
	Dynamic_FindMany_FullMethodName   = "/dynamic.Dynamic/FindMany"
	Dynamic_QueryTag_FullMethodName   = "/dynamic.Dynamic/QueryTag"
	Dynamic_QueryUser_FullMethodName  = "/dynamic.Dynamic/QueryUser"
)

// DynamicClient is the client API for Dynamic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DynamicClient interface {
	InsertPost(ctx context.Context, in *InsertPostReq, opts ...grpc.CallOption) (*InsertPostResp, error)
	GetDynamic(ctx context.Context, in *GetDynamicReq, opts ...grpc.CallOption) (*GetDynamicResp, error)
	FindPost(ctx context.Context, in *FindOneReq, opts ...grpc.CallOption) (*FindOneResp, error)
	FindMany(ctx context.Context, in *FindManyReq, opts ...grpc.CallOption) (*FindManyReqResp, error)
	QueryTag(ctx context.Context, in *QueryTagReq, opts ...grpc.CallOption) (*QueryTagResp, error)
	QueryUser(ctx context.Context, in *QueryUserReq, opts ...grpc.CallOption) (*QueryUserResp, error)
}

type dynamicClient struct {
	cc grpc.ClientConnInterface
}

func NewDynamicClient(cc grpc.ClientConnInterface) DynamicClient {
	return &dynamicClient{cc}
}

func (c *dynamicClient) InsertPost(ctx context.Context, in *InsertPostReq, opts ...grpc.CallOption) (*InsertPostResp, error) {
	out := new(InsertPostResp)
	err := c.cc.Invoke(ctx, Dynamic_InsertPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicClient) GetDynamic(ctx context.Context, in *GetDynamicReq, opts ...grpc.CallOption) (*GetDynamicResp, error) {
	out := new(GetDynamicResp)
	err := c.cc.Invoke(ctx, Dynamic_GetDynamic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicClient) FindPost(ctx context.Context, in *FindOneReq, opts ...grpc.CallOption) (*FindOneResp, error) {
	out := new(FindOneResp)
	err := c.cc.Invoke(ctx, Dynamic_FindPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicClient) FindMany(ctx context.Context, in *FindManyReq, opts ...grpc.CallOption) (*FindManyReqResp, error) {
	out := new(FindManyReqResp)
	err := c.cc.Invoke(ctx, Dynamic_FindMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicClient) QueryTag(ctx context.Context, in *QueryTagReq, opts ...grpc.CallOption) (*QueryTagResp, error) {
	out := new(QueryTagResp)
	err := c.cc.Invoke(ctx, Dynamic_QueryTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dynamicClient) QueryUser(ctx context.Context, in *QueryUserReq, opts ...grpc.CallOption) (*QueryUserResp, error) {
	out := new(QueryUserResp)
	err := c.cc.Invoke(ctx, Dynamic_QueryUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DynamicServer is the server API for Dynamic service.
// All implementations must embed UnimplementedDynamicServer
// for forward compatibility
type DynamicServer interface {
	InsertPost(context.Context, *InsertPostReq) (*InsertPostResp, error)
	GetDynamic(context.Context, *GetDynamicReq) (*GetDynamicResp, error)
	FindPost(context.Context, *FindOneReq) (*FindOneResp, error)
	FindMany(context.Context, *FindManyReq) (*FindManyReqResp, error)
	QueryTag(context.Context, *QueryTagReq) (*QueryTagResp, error)
	QueryUser(context.Context, *QueryUserReq) (*QueryUserResp, error)
	mustEmbedUnimplementedDynamicServer()
}

// UnimplementedDynamicServer must be embedded to have forward compatible implementations.
type UnimplementedDynamicServer struct {
}

func (UnimplementedDynamicServer) InsertPost(context.Context, *InsertPostReq) (*InsertPostResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPost not implemented")
}
func (UnimplementedDynamicServer) GetDynamic(context.Context, *GetDynamicReq) (*GetDynamicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDynamic not implemented")
}
func (UnimplementedDynamicServer) FindPost(context.Context, *FindOneReq) (*FindOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPost not implemented")
}
func (UnimplementedDynamicServer) FindMany(context.Context, *FindManyReq) (*FindManyReqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMany not implemented")
}
func (UnimplementedDynamicServer) QueryTag(context.Context, *QueryTagReq) (*QueryTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}
func (UnimplementedDynamicServer) QueryUser(context.Context, *QueryUserReq) (*QueryUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}
func (UnimplementedDynamicServer) mustEmbedUnimplementedDynamicServer() {}

// UnsafeDynamicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DynamicServer will
// result in compilation errors.
type UnsafeDynamicServer interface {
	mustEmbedUnimplementedDynamicServer()
}

func RegisterDynamicServer(s grpc.ServiceRegistrar, srv DynamicServer) {
	s.RegisterService(&Dynamic_ServiceDesc, srv)
}

func _Dynamic_InsertPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).InsertPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_InsertPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).InsertPost(ctx, req.(*InsertPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dynamic_GetDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDynamicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).GetDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_GetDynamic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).GetDynamic(ctx, req.(*GetDynamicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dynamic_FindPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).FindPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_FindPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).FindPost(ctx, req.(*FindOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dynamic_FindMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindManyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).FindMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_FindMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).FindMany(ctx, req.(*FindManyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dynamic_QueryTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).QueryTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_QueryTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).QueryTag(ctx, req.(*QueryTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dynamic_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DynamicServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dynamic_QueryUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DynamicServer).QueryUser(ctx, req.(*QueryUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Dynamic_ServiceDesc is the grpc.ServiceDesc for Dynamic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dynamic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dynamic.Dynamic",
	HandlerType: (*DynamicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertPost",
			Handler:    _Dynamic_InsertPost_Handler,
		},
		{
			MethodName: "GetDynamic",
			Handler:    _Dynamic_GetDynamic_Handler,
		},
		{
			MethodName: "FindPost",
			Handler:    _Dynamic_FindPost_Handler,
		},
		{
			MethodName: "FindMany",
			Handler:    _Dynamic_FindMany_Handler,
		},
		{
			MethodName: "QueryTag",
			Handler:    _Dynamic_QueryTag_Handler,
		},
		{
			MethodName: "QueryUser",
			Handler:    _Dynamic_QueryUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dynamic.proto",
}
