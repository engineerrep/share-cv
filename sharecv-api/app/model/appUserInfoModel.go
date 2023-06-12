package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserInfoModel = (*customAppUserInfoModel)(nil)

type (
	// AppUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserInfoModel.
	AppUserInfoModel interface {
		appUserInfoModel
	}

	customAppUserInfoModel struct {
		*defaultAppUserInfoModel
	}
)

// NewAppUserInfoModel returns a model for the database table.
func NewAppUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserInfoModel {
	return &customAppUserInfoModel{
		defaultAppUserInfoModel: newAppUserInfoModel(conn, c),
	}
}
