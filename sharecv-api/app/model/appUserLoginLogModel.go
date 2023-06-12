package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserLoginLogModel = (*customAppUserLoginLogModel)(nil)

type (
	// AppUserLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserLoginLogModel.
	AppUserLoginLogModel interface {
		appUserLoginLogModel
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserLoginLog, error)
		FindByPlatform(ctx context.Context, userId int64, platform int64) (*AppUserLoginLog, error)
	}

	customAppUserLoginLogModel struct {
		*defaultAppUserLoginLogModel
	}
)

// NewAppUserLoginLogModel returns a model for the database table.
func NewAppUserLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserLoginLogModel {
	return &customAppUserLoginLogModel{
		defaultAppUserLoginLogModel: newAppUserLoginLogModel(conn, c),
	}
}

func (m *customAppUserLoginLogModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserLoginLog, error) {
	var resp []*AppUserLoginLog
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", appUserLoginLogRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAppUserLoginLogModel) FindByPlatform(ctx context.Context, userId int64, platform int64) (*AppUserLoginLog, error) {
	var resp AppUserLoginLog
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `platform` = ? limit 1", appUserLoginLogRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, platform)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
