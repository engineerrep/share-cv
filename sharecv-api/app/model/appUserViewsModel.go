package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserViewsModel = (*customAppUserViewsModel)(nil)

type (
	// AppUserViewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserViewsModel.
	AppUserViewsModel interface {
		appUserViewsModel
	}

	customAppUserViewsModel struct {
		*defaultAppUserViewsModel
	}
)

// NewAppUserViewsModel returns a model for the database table.
func NewAppUserViewsModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserViewsModel {
	return &customAppUserViewsModel{
		defaultAppUserViewsModel: newAppUserViewsModel(conn, c),
	}
}
