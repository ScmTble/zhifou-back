package tag

import (
	"context"
	"git.154896.xyz/zhifou/tag/tag"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTagLogic {
	return &QueryTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTagLogic) QueryTag(req *types.QueryTagReq) (resp *types.QueryTagResp, err error) {
	one, err := l.svcCtx.TagRpc.FindOne(l.ctx, &tag.FindOneReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryTagResp{
		CommonTagResp: &types.CommonTagResp{
			Id:       one.Id,
			Tag:      one.Tag,
			Avatar:   one.Avatar,
			QuoteNum: one.QuoteNum,
		},
	}, nil
}
