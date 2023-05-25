package logic

import (
	"context"
	"errors"
	"git.154896.xyz/zhifou/user/internal/svc"
	"git.154896.xyz/zhifou/user/model"
	"git.154896.xyz/zhifou/user/pb"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var user model.User
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"username": in.Username,
	}).One(&user); err != nil {
		return nil, err
	}
	passwdMd5 := cryptor.Md5String(in.Password)
	if passwdMd5 != user.Password {
		// 密码不正确
		return nil, errors.New("the password is incorrect")
	}

	var resp pb.LoginResp
	_ = copier.Copy(&resp, &user)

	return &resp, nil
}
