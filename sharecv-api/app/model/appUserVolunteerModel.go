package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserVolunteerModel = (*customAppUserVolunteerModel)(nil)

type (
	// AppUserVolunteerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserVolunteerModel.
	AppUserVolunteerModel interface {
		appUserVolunteerModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserVolunteer, error)
	}

	customAppUserVolunteerModel struct {
		*defaultAppUserVolunteerModel
	}
)

// NewAppUserVolunteerModel returns a model for the database table.
func NewAppUserVolunteerModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserVolunteerModel {
	return &customAppUserVolunteerModel{
		defaultAppUserVolunteerModel: newAppUserVolunteerModel(conn, c),
	}
}

func (m *customAppUserVolunteerModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserVolunteer, error) {
	var resp []*AppUserVolunteer
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc, id asc", appUserVolunteerRows, m.table)
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
