package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppConfigModel = (*customAppConfigModel)(nil)

type (
	// AppConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppConfigModel.
	AppConfigModel interface {
		appConfigModel
	}

	customAppConfigModel struct {
		*defaultAppConfigModel
	}
)

// NewAppConfigModel returns a model for the database table.
func NewAppConfigModel(conn sqlx.SqlConn, c cache.CacheConf) AppConfigModel {
	return &customAppConfigModel{
		defaultAppConfigModel: newAppConfigModel(conn, c),
	}
}
