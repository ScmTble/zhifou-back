package user

import (
	"context"
	"time"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"git.154896.xyz/zhifou/pkg/jwtx"
	"git.154896.xyz/zhifou/user/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	u, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	// 生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, err := jwtx.GenerateToken(accessSecret, now, accessExpire, u.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{Token: token}, nil
}
