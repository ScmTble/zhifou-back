package user

import (
	"context"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"git.154896.xyz/zhifou/pkg/tool"
	"git.154896.xyz/zhifou/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info() (resp *types.CommonUserResp, err error) {
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	u, err := l.svcCtx.UserRpc.GetOne(l.ctx, &user.GetOneReq{
		Id: uid,
	})

	if err != nil {
		return nil, err
	}

	return &types.CommonUserResp{
		Id:       u.Id,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		IsAdmin:  u.IsAdmin,
	}, err
}
