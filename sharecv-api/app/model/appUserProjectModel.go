package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserProjectModel = (*customAppUserProjectModel)(nil)

type (
	// AppUserProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserProjectModel.
	AppUserProjectModel interface {
		appUserProjectModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserProject, error)
	}

	customAppUserProjectModel struct {
		*defaultAppUserProjectModel
	}
)

// NewAppUserProjectModel returns a model for the database table.
func NewAppUserProjectModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserProjectModel {
	return &customAppUserProjectModel{
		defaultAppUserProjectModel: newAppUserProjectModel(conn, c),
	}
}

func (m *customAppUserProjectModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserProject, error) {
	var resp []*AppUserProject
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc", appUserProjectRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
