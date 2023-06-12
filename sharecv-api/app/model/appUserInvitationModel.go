package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserInvitationModel = (*customAppUserInvitationModel)(nil)

type (
	// AppUserInvitationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserInvitationModel.
	AppUserInvitationModel interface {
		appUserInvitationModel
	}

	customAppUserInvitationModel struct {
		*defaultAppUserInvitationModel
	}
)

// NewAppUserInvitationModel returns a model for the database table.
func NewAppUserInvitationModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserInvitationModel {
	return &customAppUserInvitationModel{
		defaultAppUserInvitationModel: newAppUserInvitationModel(conn, c),
	}
}
