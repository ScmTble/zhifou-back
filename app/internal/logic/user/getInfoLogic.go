package user

import (
	"context"
	"git.154896.xyz/zhifou/user/user"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo(req *types.GetInfoReq) (resp *types.GetInfoResp, err error) {
	u, err := l.svcCtx.UserRpc.GetOneByName(l.ctx, &user.GetOneByNameReq{
		UserName: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetInfoResp{
		Id:       u.Id,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}, nil
}
