package comment

import (
	"context"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/user/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsLogic) GetComments(req *types.GetCommentReq) (*types.GetCommentResp, error) {
	// 获取评论
	resp, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &comment.GetCommentReq{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}

	comments, err := mr.MapReduce(func(source chan<- *comment.CommonCommentResp) {
		for _, v := range resp.Comments {
			source <- v
		}
	}, func(item *comment.CommonCommentResp, writer mr.Writer[*types.CommentResp], cancel func(error)) {
		var resp types.CommentResp
		_ = copier.Copy(&resp, item)
		// 注入userinfo
		mr.FinishVoid(func() {
			u, err := l.svcCtx.UserRpc.GetOne(l.ctx, &user.GetOneReq{Id: item.UserId})
			if err != nil {
				return
			}
			var uu types.CommonUserResp
			_ = copier.Copy(&uu, u)
			resp.User = &uu
		})
		writer.Write(&resp)
	}, func(pipe <-chan *types.CommentResp, writer mr.Writer[[]*types.CommentResp], cancel func(error)) {
		var result []*types.CommentResp
		for item := range pipe {
			result = append(result, item)
		}
		writer.Write(result)
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCommentResp{Comments: comments}, nil
}
