package logic

import (
	"context"
	"git.154896.xyz/zhifou/idgen/idgen"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/jinzhu/copier"
	"time"

	"git.154896.xyz/zhifou/pkg/tool"
	"git.154896.xyz/zhifou/user/model"

	"git.154896.xyz/zhifou/user/internal/svc"
	"git.154896.xyz/zhifou/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	result, err := l.svcCtx.IgRpc.IdGenServer(l.ctx, &idgen.Empty{})
	if err != nil {
		return nil, err
	}

	createdTime := time.Now().Unix()
	passwdMd5 := cryptor.Md5String(in.Password)
	user := model.User{
		Id:          result.Id,
		Nickname:    in.Username,
		Username:    in.Username,
		Password:    passwdMd5,
		Avatar:      tool.GetRandomAvatar(),
		IsAdmin:     false,
		CreatedTime: createdTime,
		UpdatedTime: 0,
		DeletedTime: 0,
	}

	_, err = l.svcCtx.MongoDB.InsertOne(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	var resp pb.RegisterResp
	_ = copier.Copy(&resp, &user)

	return &resp, nil
}
