// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"git.154896.xyz/zhifou/comment/internal/logic"
	"git.154896.xyz/zhifou/comment/internal/svc"
	"git.154896.xyz/zhifou/comment/pb"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) InsertComment(ctx context.Context, in *pb.InsertCommentReq) (*pb.InsertCommentResp, error) {
	l := logic.NewInsertCommentLogic(ctx, s.svcCtx)
	return l.InsertComment(in)
}

func (s *CommentServer) GetComment(ctx context.Context, in *pb.GetCommentReq) (*pb.GetCommentResp, error) {
	l := logic.NewGetCommentLogic(ctx, s.svcCtx)
	return l.GetComment(in)
}

func (s *CommentServer) GetCommentCount(ctx context.Context, in *pb.GetCommentCountReq) (*pb.GetCommentCountResp, error) {
	l := logic.NewGetCommentCountLogic(ctx, s.svcCtx)
	return l.GetCommentCount(in)
}
