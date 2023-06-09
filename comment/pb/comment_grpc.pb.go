// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: comment.proto

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
	Comment_InsertComment_FullMethodName   = "/comment.Comment/InsertComment"
	Comment_GetComment_FullMethodName      = "/comment.Comment/GetComment"
	Comment_GetCommentCount_FullMethodName = "/comment.Comment/GetCommentCount"
)

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	InsertComment(ctx context.Context, in *InsertCommentReq, opts ...grpc.CallOption) (*InsertCommentResp, error)
	GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error)
	GetCommentCount(ctx context.Context, in *GetCommentCountReq, opts ...grpc.CallOption) (*GetCommentCountResp, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) InsertComment(ctx context.Context, in *InsertCommentReq, opts ...grpc.CallOption) (*InsertCommentResp, error) {
	out := new(InsertCommentResp)
	err := c.cc.Invoke(ctx, Comment_InsertComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error) {
	out := new(GetCommentResp)
	err := c.cc.Invoke(ctx, Comment_GetComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentCount(ctx context.Context, in *GetCommentCountReq, opts ...grpc.CallOption) (*GetCommentCountResp, error) {
	out := new(GetCommentCountResp)
	err := c.cc.Invoke(ctx, Comment_GetCommentCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	InsertComment(context.Context, *InsertCommentReq) (*InsertCommentResp, error)
	GetComment(context.Context, *GetCommentReq) (*GetCommentResp, error)
	GetCommentCount(context.Context, *GetCommentCountReq) (*GetCommentCountResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) InsertComment(context.Context, *InsertCommentReq) (*InsertCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertComment not implemented")
}
func (UnimplementedCommentServer) GetComment(context.Context, *GetCommentReq) (*GetCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentServer) GetCommentCount(context.Context, *GetCommentCountReq) (*GetCommentCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentCount not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_InsertComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).InsertComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_InsertComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).InsertComment(ctx, req.(*InsertCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetComment(ctx, req.(*GetCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_GetCommentCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentCount(ctx, req.(*GetCommentCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertComment",
			Handler:    _Comment_InsertComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _Comment_GetComment_Handler,
		},
		{
			MethodName: "GetCommentCount",
			Handler:    _Comment_GetCommentCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
