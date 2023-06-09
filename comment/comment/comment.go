// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package comment

import (
	"context"

	"git.154896.xyz/zhifou/comment/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonCommentResp   = pb.CommonCommentResp
	GetCommentCountReq  = pb.GetCommentCountReq
	GetCommentCountResp = pb.GetCommentCountResp
	GetCommentReq       = pb.GetCommentReq
	GetCommentResp      = pb.GetCommentResp
	InsertCommentReq    = pb.InsertCommentReq
	InsertCommentResp   = pb.InsertCommentResp

	Comment interface {
		InsertComment(ctx context.Context, in *InsertCommentReq, opts ...grpc.CallOption) (*InsertCommentResp, error)
		GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error)
		GetCommentCount(ctx context.Context, in *GetCommentCountReq, opts ...grpc.CallOption) (*GetCommentCountResp, error)
	}

	defaultComment struct {
		cli zrpc.Client
	}
)

func NewComment(cli zrpc.Client) Comment {
	return &defaultComment{
		cli: cli,
	}
}

func (m *defaultComment) InsertComment(ctx context.Context, in *InsertCommentReq, opts ...grpc.CallOption) (*InsertCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.InsertComment(ctx, in, opts...)
}

func (m *defaultComment) GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetComment(ctx, in, opts...)
}

func (m *defaultComment) GetCommentCount(ctx context.Context, in *GetCommentCountReq, opts ...grpc.CallOption) (*GetCommentCountResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetCommentCount(ctx, in, opts...)
}
