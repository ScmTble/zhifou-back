// Code generated by goctl. DO NOT EDIT.
// Source: collect.proto

package server

import (
	"context"

	"git.154896.xyz/zhifou/collect/internal/logic"
	"git.154896.xyz/zhifou/collect/internal/svc"
	"git.154896.xyz/zhifou/collect/pb"
)

type CollectServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCollectServer
}

func NewCollectServer(svcCtx *svc.ServiceContext) *CollectServer {
	return &CollectServer{
		svcCtx: svcCtx,
	}
}

func (s *CollectServer) InsertCollect(ctx context.Context, in *pb.InsertCollectReq) (*pb.InsertCollectResp, error) {
	l := logic.NewInsertCollectLogic(ctx, s.svcCtx)
	return l.InsertCollect(in)
}

func (s *CollectServer) GetCollect(ctx context.Context, in *pb.GetCollectReq) (*pb.GetCollectResp, error) {
	l := logic.NewGetCollectLogic(ctx, s.svcCtx)
	return l.GetCollect(in)
}

func (s *CollectServer) GetStatus(ctx context.Context, in *pb.GetStatusReq) (*pb.GetStatusResp, error) {
	l := logic.NewGetStatusLogic(ctx, s.svcCtx)
	return l.GetStatus(in)
}

func (s *CollectServer) GetNum(ctx context.Context, in *pb.GetNumReq) (*pb.GetNumResp, error) {
	l := logic.NewGetNumLogic(ctx, s.svcCtx)
	return l.GetNum(in)
}
