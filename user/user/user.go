// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"git.154896.xyz/zhifou/user/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetOneByNameReq  = pb.GetOneByNameReq
	GetOneByNameResp = pb.GetOneByNameResp
	GetOneReq        = pb.GetOneReq
	GetOneResp       = pb.GetOneResp
	LoginReq         = pb.LoginReq
	LoginResp        = pb.LoginResp
	RegisterReq      = pb.RegisterReq
	RegisterResp     = pb.RegisterResp

	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneResp, error)
		GetOneByName(ctx context.Context, in *GetOneByNameReq, opts ...grpc.CallOption) (*GetOneByNameResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) GetOne(ctx context.Context, in *GetOneReq, opts ...grpc.CallOption) (*GetOneResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetOne(ctx, in, opts...)
}

func (m *defaultUser) GetOneByName(ctx context.Context, in *GetOneByNameReq, opts ...grpc.CallOption) (*GetOneByNameResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetOneByName(ctx, in, opts...)
}
