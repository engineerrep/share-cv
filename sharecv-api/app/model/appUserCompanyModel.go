package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserCompanyModel = (*customAppUserCompanyModel)(nil)

type (
	// AppUserCompanyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserCompanyModel.
	AppUserCompanyModel interface {
		appUserCompanyModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserCompany, error)
	}

	customAppUserCompanyModel struct {
		*defaultAppUserCompanyModel
	}
)

// NewAppUserCompanyModel returns a model for the database table.
func NewAppUserCompanyModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserCompanyModel {
	return &customAppUserCompanyModel{
		defaultAppUserCompanyModel: newAppUserCompanyModel(conn, c),
	}
}

func (m *customAppUserCompanyModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserCompany, error) {
	var resp []*AppUserCompany
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc", appUserCompanyRows, m.table)
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
