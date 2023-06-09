// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: like.proto

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
	Like_GetNum_FullMethodName       = "/user.Like/GetNum"
	Like_InsertUpvote_FullMethodName = "/user.Like/InsertUpvote"
	Like_GetStatus_FullMethodName    = "/user.Like/GetStatus"
)

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	GetNum(ctx context.Context, in *GetNumReq, opts ...grpc.CallOption) (*GetNumResp, error)
	InsertUpvote(ctx context.Context, in *InsertUpvoteReq, opts ...grpc.CallOption) (*InsertUpvoteResp, error)
	GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) GetNum(ctx context.Context, in *GetNumReq, opts ...grpc.CallOption) (*GetNumResp, error) {
	out := new(GetNumResp)
	err := c.cc.Invoke(ctx, Like_GetNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) InsertUpvote(ctx context.Context, in *InsertUpvoteReq, opts ...grpc.CallOption) (*InsertUpvoteResp, error) {
	out := new(InsertUpvoteResp)
	err := c.cc.Invoke(ctx, Like_InsertUpvote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) GetStatus(ctx context.Context, in *GetStatusReq, opts ...grpc.CallOption) (*GetStatusResp, error) {
	out := new(GetStatusResp)
	err := c.cc.Invoke(ctx, Like_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility
type LikeServer interface {
	GetNum(context.Context, *GetNumReq) (*GetNumResp, error)
	InsertUpvote(context.Context, *InsertUpvoteReq) (*InsertUpvoteResp, error)
	GetStatus(context.Context, *GetStatusReq) (*GetStatusResp, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServer struct {
}

func (UnimplementedLikeServer) GetNum(context.Context, *GetNumReq) (*GetNumResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNum not implemented")
}
func (UnimplementedLikeServer) InsertUpvote(context.Context, *InsertUpvoteReq) (*InsertUpvoteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUpvote not implemented")
}
func (UnimplementedLikeServer) GetStatus(context.Context, *GetStatusReq) (*GetStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_GetNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_GetNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetNum(ctx, req.(*GetNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_InsertUpvote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUpvoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).InsertUpvote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_InsertUpvote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).InsertUpvote(ctx, req.(*InsertUpvoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).GetStatus(ctx, req.(*GetStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNum",
			Handler:    _Like_GetNum_Handler,
		},
		{
			MethodName: "InsertUpvote",
			Handler:    _Like_InsertUpvote_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Like_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}
