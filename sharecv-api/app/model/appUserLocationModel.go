package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserLocationModel = (*customAppUserLocationModel)(nil)

type (
	// AppUserLocationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserLocationModel.
	AppUserLocationModel interface {
		appUserLocationModel
	}

	customAppUserLocationModel struct {
		*defaultAppUserLocationModel
	}
)

// NewAppUserLocationModel returns a model for the database table.
func NewAppUserLocationModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserLocationModel {
	return &customAppUserLocationModel{
		defaultAppUserLocationModel: newAppUserLocationModel(conn, c),
	}
}
