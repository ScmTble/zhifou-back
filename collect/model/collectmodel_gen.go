// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	collectFieldNames          = builder.RawFieldNames(&Collect{})
	collectRows                = strings.Join(collectFieldNames, ",")
	collectRowsExpectAutoSet   = strings.Join(stringx.Remove(collectFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	collectRowsWithPlaceHolder = strings.Join(stringx.Remove(collectFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	collectModel interface {
		Insert(ctx context.Context, data *Collect) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Collect, error)
		Update(ctx context.Context, data *Collect) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCollectModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Collect struct {
		Id          int64        `db:"id"`
		PostId      int64        `db:"post_id"`
		UserId      int64        `db:"user_id"`
		CreatedTime time.Time    `db:"created_time"`
		UpdatedTime sql.NullTime `db:"updated_time"`
		DeletedTime sql.NullTime `db:"deleted_time"`
	}
)

func newCollectModel(conn sqlx.SqlConn) *defaultCollectModel {
	return &defaultCollectModel{
		conn:  conn,
		table: "`collect`",
	}
}

func (m *defaultCollectModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCollectModel) FindOne(ctx context.Context, id int64) (*Collect, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", collectRows, m.table)
	var resp Collect
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCollectModel) Insert(ctx context.Context, data *Collect) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, collectRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.PostId, data.UserId, data.CreatedTime, data.UpdatedTime, data.DeletedTime)
	return ret, err
}

func (m *defaultCollectModel) Update(ctx context.Context, data *Collect) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, collectRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PostId, data.UserId, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.Id)
	return err
}

func (m *defaultCollectModel) tableName() string {
	return m.table
}
