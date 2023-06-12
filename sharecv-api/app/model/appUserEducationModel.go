package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserEducationModel = (*customAppUserEducationModel)(nil)

type (
	// AppUserEducationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserEducationModel.
	AppUserEducationModel interface {
		appUserEducationModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserEducation, error)
	}

	customAppUserEducationModel struct {
		*defaultAppUserEducationModel
	}
)

// NewAppUserEducationModel returns a model for the database table.
func NewAppUserEducationModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserEducationModel {
	return &customAppUserEducationModel{
		defaultAppUserEducationModel: newAppUserEducationModel(conn, c),
	}
}

func (m *customAppUserEducationModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserEducation, error) {
	var resp []*AppUserEducation
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc, id asc", appUserEducationRows, m.table)
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
