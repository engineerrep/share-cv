package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserExperienceModel = (*customAppUserExperienceModel)(nil)

type (
	// AppUserExperienceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserExperienceModel.
	AppUserExperienceModel interface {
		appUserExperienceModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserExperience, error)
	}

	customAppUserExperienceModel struct {
		*defaultAppUserExperienceModel
	}
)

// NewAppUserExperienceModel returns a model for the database table.
func NewAppUserExperienceModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserExperienceModel {
	return &customAppUserExperienceModel{
		defaultAppUserExperienceModel: newAppUserExperienceModel(conn, c),
	}
}

func (m *customAppUserExperienceModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserExperience, error) {
	var resp []*AppUserExperience
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc, id asc", appUserExperienceRows, m.table)
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
