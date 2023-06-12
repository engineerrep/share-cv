package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserModel = (*customAppUserModel)(nil)

type (
	// AppUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserModel.
	AppUserModel interface {
		appUserModel
		FindByEmail(ctx context.Context, email string) (*AppUser, error)
		FindByGoogleId(ctx context.Context, googleId string) (*AppUser, error)
		FindByAppleId(ctx context.Context, appleId string) (*AppUser, error)
		FindByInvitationCode(ctx context.Context, code string) (*AppUser, error)
	}

	customAppUserModel struct {
		*defaultAppUserModel
	}
)

// NewAppUserModel returns a model for the database table.
func NewAppUserModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserModel {
	return &customAppUserModel{
		defaultAppUserModel: newAppUserModel(conn, c),
	}
}

func (m *customAppUserModel) FindByEmail(ctx context.Context, email string) (*AppUser, error) {
	var resp AppUser
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", appUserRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, email)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAppUserModel) FindByGoogleId(ctx context.Context, googleId string) (*AppUser, error) {
	var resp AppUser
	query := fmt.Sprintf("select %s from %s where `google_uid` = ? limit 1", appUserRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, googleId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAppUserModel) FindByAppleId(ctx context.Context, appleId string) (*AppUser, error) {
	var resp AppUser
	query := fmt.Sprintf("select %s from %s where `apple_uid` = ? limit 1", appUserRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, appleId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customAppUserModel) FindByInvitationCode(ctx context.Context, code string) (*AppUser, error) {
	var resp AppUser
	query := fmt.Sprintf("select %s from %s where `invitation_code` = ? limit 1", appUserRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, code)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
