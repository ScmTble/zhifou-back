// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package server

import (
	"context"

	"git.154896.xyz/zhifou/like/internal/logic"
	"git.154896.xyz/zhifou/like/internal/svc"
	"git.154896.xyz/zhifou/like/pb"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) GetNum(ctx context.Context, in *pb.GetNumReq) (*pb.GetNumResp, error) {
	l := logic.NewGetNumLogic(ctx, s.svcCtx)
	return l.GetNum(in)
}

func (s *LikeServer) InsertUpvote(ctx context.Context, in *pb.InsertUpvoteReq) (*pb.InsertUpvoteResp, error) {
	l := logic.NewInsertUpvoteLogic(ctx, s.svcCtx)
	return l.InsertUpvote(in)
}

func (s *LikeServer) GetStatus(ctx context.Context, in *pb.GetStatusReq) (*pb.GetStatusResp, error) {
	l := logic.NewGetStatusLogic(ctx, s.svcCtx)
	return l.GetStatus(in)
}
