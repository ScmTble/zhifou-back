package comment

import (
	"context"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/pkg/tool"
	"github.com/jinzhu/copier"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCommentLogic {
	return &InsertCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertCommentLogic) InsertComment(req *types.InsertCoomentReq) (*types.InsertCoomentResp, error) {
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	com, err := l.svcCtx.CommentRpc.InsertComment(l.ctx, &comment.InsertCommentReq{
		UserId:  uid,
		PostId:  req.PostId,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	var resp types.InsertCoomentResp
	_ = copier.Copy(&resp, com)

	return &resp, nil
}
