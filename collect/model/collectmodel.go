package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollectModel = (*customCollectModel)(nil)

type (
	// CollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectModel.
	CollectModel interface {
		collectModel
		GetMany(ctx context.Context, userId int64) ([]int64, error)
		GetOne(ctx context.Context, userId int64, postId int64) error
	}

	customCollectModel struct {
		*defaultCollectModel
	}
)

func (m *customCollectModel) GetOne(ctx context.Context, userId int64, postId int64) error {
	var id int64
	query, args, err := squirrel.Select("id").From(m.tableName()).Where(squirrel.Eq{
		"post_id": postId,
		"user_id": userId,
	}).ToSql()
	if err != nil {
		return err
	}
	err = m.conn.QueryRowCtx(ctx, &id, query, args...)
	return err
}

func (m *customCollectModel) GetMany(ctx context.Context, userId int64) ([]int64, error) {
	query, args, err := squirrel.Select("post_id").From(m.tableName()).Where(squirrel.Eq{
		"deleted_time": nil,
		"user_id":      userId,
	}).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []int64
	if err = m.conn.QueryRowsCtx(ctx, &resp, query, args...); err != nil {
		return nil, err
	}

	return resp, nil
}

// NewCollectModel returns a model for the database table.
func NewCollectModel(conn sqlx.SqlConn) CollectModel {
	return &customCollectModel{
		defaultCollectModel: newCollectModel(conn),
	}
}
