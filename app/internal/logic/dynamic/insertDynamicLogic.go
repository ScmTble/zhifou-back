package dynamic

import (
	"context"
	"strconv"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"git.154896.xyz/zhifou/dynamic/dynamic"
	"git.154896.xyz/zhifou/pkg/tool"
	"git.154896.xyz/zhifou/tag/tag"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type InsertDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertDynamicLogic {
	return &InsertDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertDynamicLogic) getFullTag(tags []string) ([]*dynamic.Tag, error) {
	result, err := mr.MapReduce(func(source chan<- string) {
		for _, v := range tags {
			source <- v
		}
	}, func(item string, writer mr.Writer[*dynamic.Tag], cancel func(error)) {
		one, err := l.svcCtx.TagRpc.FindOne(l.ctx, &tag.FindOneReq{Id: item})
		if err != nil {
			cancel(err)
		}
		writer.Write(&dynamic.Tag{
			Label: one.Tag,
			Value: one.Id,
		})
	}, func(pipe <-chan *dynamic.Tag, writer mr.Writer[[]*dynamic.Tag], cancel func(error)) {
		var tags []*dynamic.Tag
		for v := range pipe {
			tags = append(tags, v)
		}
		writer.Write(tags)
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *InsertDynamicLogic) getTagString(tags []int64) []string {
	var result []string
	for _, v := range tags {
		result = append(result, strconv.FormatInt(v, 10))
	}
	return result
}

// InsertDynamic 新增动态
func (l *InsertDynamicLogic) InsertDynamic(req *types.InsertPostReq) (*types.InsertPostResp, error) {
	// 获取uid
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 根据Tag id获取完整的tag
	fullTag, err := l.getFullTag(req.Tags)
	if err != nil {
		return nil, err
	}

	// 构建插入grpc参数
	insertPostReq := &dynamic.InsertPostReq{}
	_ = copier.Copy(insertPostReq, req)
	insertPostReq.UserId = uid
	insertPostReq.Tags = fullTag

	// 插入动态
	post, err := l.svcCtx.DynamicRpc.InsertPost(l.ctx, insertPostReq)
	if err != nil {
		return nil, err
	}

	// 构建返回参数
	var resp types.InsertPostResp
	_ = copier.Copy(&resp, post)

	return &resp, nil
}
