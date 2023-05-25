package tag

import (
	"context"
	"errors"

	"git.154896.xyz/zhifou/tag/pb"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagLogic {
	return &ListTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListTag 获取Tag列表
func (l *ListTagLogic) ListTag(req *types.ListTagReq) (resp *types.ListTagResp, err error) {
	var topics []*types.CommonTagResp

	switch req.Type {
	case "hot":
		// 获取热点Tag
		tags, err := l.svcCtx.TagRpc.ListHotTags(l.ctx, &pb.ListHotTagReq{
			Num: req.Num,
		})
		if err != nil {
			return nil, err
		}

		for _, t := range tags.Topics {
			topics = append(topics, &types.CommonTagResp{
				Id:       t.GetId(),
				Tag:      t.GetTag(),
				Avatar:   t.GetAvatar(),
				QuoteNum: t.GetQuoteNum(),
			})
		}

		return &types.ListTagResp{Topics: topics}, nil
	case "new":
		// 获取最新Tag
		tags, err := l.svcCtx.TagRpc.ListNewTags(l.ctx, &pb.ListNewTagReq{
			Num: req.Num,
		})
		if err != nil {
			return nil, err
		}

		for _, t := range tags.Topics {
			topics = append(topics, &types.CommonTagResp{
				Id:       t.GetId(),
				Tag:      t.GetTag(),
				Avatar:   t.GetAvatar(),
				QuoteNum: t.GetQuoteNum(),
			})
		}

		return &types.ListTagResp{Topics: topics}, nil
	default:
		return nil, errors.New("type is error")
	}
}
