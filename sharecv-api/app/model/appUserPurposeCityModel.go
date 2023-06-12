package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserPurposeCityModel = (*customAppUserPurposeCityModel)(nil)

type (
	// AppUserPurposeCityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserPurposeCityModel.
	AppUserPurposeCityModel interface {
		appUserPurposeCityModel
	}

	customAppUserPurposeCityModel struct {
		*defaultAppUserPurposeCityModel
	}
)

// NewAppUserPurposeCityModel returns a model for the database table.
func NewAppUserPurposeCityModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserPurposeCityModel {
	return &customAppUserPurposeCityModel{
		defaultAppUserPurposeCityModel: newAppUserPurposeCityModel(conn, c),
	}
}
