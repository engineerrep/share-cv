// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	appUserInvitationFieldNames          = builder.RawFieldNames(&AppUserInvitation{})
	appUserInvitationRows                = strings.Join(appUserInvitationFieldNames, ",")
	appUserInvitationRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserInvitationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	appUserInvitationRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserInvitationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAppUserInvitationIdPrefix = "cache:appUserInvitation:id:"
)

type (
	appUserInvitationModel interface {
		Insert(ctx context.Context, data *AppUserInvitation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AppUserInvitation, error)
		Update(ctx context.Context, data *AppUserInvitation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAppUserInvitationModel struct {
		sqlc.CachedConn
		table string
	}

	AppUserInvitation struct {
		Id             int64        `db:"id"`
		UserId         int64        `db:"user_id"`         // 登录用户id(app_user表id)
		SuperUserId    int64        `db:"super_user_id"`   // 上级用户id(app_user表id)
		InvitationCode string       `db:"invitation_code"` // 邀请码
		ChannelName    string       `db:"channel_name"`    // 渠道名称
		CreateAt       time.Time    `db:"create_at"`       // 创建时间
		UpdateAt       time.Time    `db:"update_at"`       // 更新时间
		DeletedAt      sql.NullTime `db:"deleted_at"`      // 删除时间
	}
)

func newAppUserInvitationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAppUserInvitationModel {
	return &defaultAppUserInvitationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`app_user_invitation`",
	}
}

func (m *defaultAppUserInvitationModel) Delete(ctx context.Context, id int64) error {
	appUserInvitationIdKey := fmt.Sprintf("%s%v", cacheAppUserInvitationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, appUserInvitationIdKey)
	return err
}

func (m *defaultAppUserInvitationModel) FindOne(ctx context.Context, id int64) (*AppUserInvitation, error) {
	appUserInvitationIdKey := fmt.Sprintf("%s%v", cacheAppUserInvitationIdPrefix, id)
	var resp AppUserInvitation
	err := m.QueryRowCtx(ctx, &resp, appUserInvitationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserInvitationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppUserInvitationModel) Insert(ctx context.Context, data *AppUserInvitation) (sql.Result, error) {
	appUserInvitationIdKey := fmt.Sprintf("%s%v", cacheAppUserInvitationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, appUserInvitationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.SuperUserId, data.InvitationCode, data.ChannelName, data.DeletedAt)
	}, appUserInvitationIdKey)
	return ret, err
}

func (m *defaultAppUserInvitationModel) Update(ctx context.Context, data *AppUserInvitation) error {
	appUserInvitationIdKey := fmt.Sprintf("%s%v", cacheAppUserInvitationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserInvitationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.SuperUserId, data.InvitationCode, data.ChannelName, data.DeletedAt, data.Id)
	}, appUserInvitationIdKey)
	return err
}

func (m *defaultAppUserInvitationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAppUserInvitationIdPrefix, primary)
}

func (m *defaultAppUserInvitationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserInvitationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAppUserInvitationModel) tableName() string {
	return m.table
}