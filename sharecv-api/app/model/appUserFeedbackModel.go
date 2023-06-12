package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserFeedbackModel = (*customAppUserFeedbackModel)(nil)

type (
	// AppUserFeedbackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserFeedbackModel.
	AppUserFeedbackModel interface {
		appUserFeedbackModel
	}

	customAppUserFeedbackModel struct {
		*defaultAppUserFeedbackModel
	}
)

// NewAppUserFeedbackModel returns a model for the database table.
func NewAppUserFeedbackModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserFeedbackModel {
	return &customAppUserFeedbackModel{
		defaultAppUserFeedbackModel: newAppUserFeedbackModel(conn, c),
	}
}
