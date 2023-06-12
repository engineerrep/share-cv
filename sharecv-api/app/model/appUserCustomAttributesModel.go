package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserCustomAttributesModel = (*customAppUserCustomAttributesModel)(nil)

type (
	// AppUserCustomAttributesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserCustomAttributesModel.
	AppUserCustomAttributesModel interface {
		appUserCustomAttributesModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserCustomAttributes, error)
	}

	customAppUserCustomAttributesModel struct {
		*defaultAppUserCustomAttributesModel
	}
)

// NewAppUserCustomAttributesModel returns a model for the database table.
func NewAppUserCustomAttributesModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserCustomAttributesModel {
	return &customAppUserCustomAttributesModel{
		defaultAppUserCustomAttributesModel: newAppUserCustomAttributesModel(conn, c),
	}
}

func (m *customAppUserCustomAttributesModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserCustomAttributes, error) {
	var resp []*AppUserCustomAttributes
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc, id asc", appUserCustomAttributesRows, m.table)
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
