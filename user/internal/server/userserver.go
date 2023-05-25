// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"git.154896.xyz/zhifou/user/internal/logic"
	"git.154896.xyz/zhifou/user/internal/svc"
	"git.154896.xyz/zhifou/user/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) GetOne(ctx context.Context, in *pb.GetOneReq) (*pb.GetOneResp, error) {
	l := logic.NewGetOneLogic(ctx, s.svcCtx)
	return l.GetOne(in)
}

func (s *UserServer) GetOneByName(ctx context.Context, in *pb.GetOneByNameReq) (*pb.GetOneByNameResp, error) {
	l := logic.NewGetOneByNameLogic(ctx, s.svcCtx)
	return l.GetOneByName(in)
}