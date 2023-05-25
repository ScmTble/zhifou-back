// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package server

import (
	"context"

	"git.154896.xyz/zhifou/tag/internal/logic"
	"git.154896.xyz/zhifou/tag/internal/svc"
	"git.154896.xyz/zhifou/tag/pb"
)

type TagServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTagServer
}

func NewTagServer(svcCtx *svc.ServiceContext) *TagServer {
	return &TagServer{
		svcCtx: svcCtx,
	}
}

func (s *TagServer) InsertTag(ctx context.Context, in *pb.InsertTagReq) (*pb.InsertTagResp, error) {
	l := logic.NewInsertTagLogic(ctx, s.svcCtx)
	return l.InsertTag(in)
}

func (s *TagServer) ListTags(ctx context.Context, in *pb.ListTagReq) (*pb.ListTagResp, error) {
	l := logic.NewListTagsLogic(ctx, s.svcCtx)
	return l.ListTags(in)
}

func (s *TagServer) ListHotTags(ctx context.Context, in *pb.ListHotTagReq) (*pb.ListHotTagResp, error) {
	l := logic.NewListHotTagsLogic(ctx, s.svcCtx)
	return l.ListHotTags(in)
}

func (s *TagServer) ListNewTags(ctx context.Context, in *pb.ListNewTagReq) (*pb.ListNewTagResp, error) {
	l := logic.NewListNewTagsLogic(ctx, s.svcCtx)
	return l.ListNewTags(in)
}

func (s *TagServer) FindOne(ctx context.Context, in *pb.FindOneReq) (*pb.FindOneResp, error) {
	l := logic.NewFindOneLogic(ctx, s.svcCtx)
	return l.FindOne(in)
}