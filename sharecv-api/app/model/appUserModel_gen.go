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
	appUserFieldNames          = builder.RawFieldNames(&AppUser{})
	appUserRows                = strings.Join(appUserFieldNames, ",")
	appUserRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	appUserRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAppUserIdPrefix   = "cache:appUser:id:"
	cacheAppUserUuidPrefix = "cache:appUser:uuid:"
)

type (
	appUserModel interface {
		Insert(ctx context.Context, data *AppUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AppUser, error)
		FindOneByUuid(ctx context.Context, uuid string) (*AppUser, error)
		Update(ctx context.Context, data *AppUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAppUserModel struct {
		sqlc.CachedConn
		table string
	}

	AppUser struct {
		Id               int64        `db:"id"`                 // 编号
		Uuid             string       `db:"uuid"`               // uuid
		AppleUid         string       `db:"apple_uid"`          // 苹果授权ID
		GoogleUid        string       `db:"google_uid"`         // 谷歌授权ID
		Email            string       `db:"email"`              // 邮件
		Mobile           string       `db:"mobile"`             // 手机号
		MobileCountryId  int64        `db:"mobile_country_id"`  // 手机国家ID
		MobileCountryNum string       `db:"mobile_country_num"` // 手机国家区号
		Locale           string       `db:"locale"`             // 本地语言
		UserType         int64        `db:"user_type"`          // 账户类型 1=邮箱 2=手机 3=谷歌 4=苹果
		Password         string       `db:"password"`           // 密码
		InvitationCode   string       `db:"invitation_code"`    // 邀请码
		CreateAt         time.Time    `db:"create_at"`          // 创建时间
		UpdateAt         time.Time    `db:"update_at"`          // 更新时间
		DeletedAt        sql.NullTime `db:"deleted_at"`         // 删除时间
	}
)

func newAppUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAppUserModel {
	return &defaultAppUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`app_user`",
	}
}

func (m *defaultAppUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, id)
	appUserUuidKey := fmt.Sprintf("%s%v", cacheAppUserUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, appUserIdKey, appUserUuidKey)
	return err
}

func (m *defaultAppUserModel) FindOne(ctx context.Context, id int64) (*AppUser, error) {
	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, id)
	var resp AppUser
	err := m.QueryRowCtx(ctx, &resp, appUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserRows, m.table)
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

func (m *defaultAppUserModel) FindOneByUuid(ctx context.Context, uuid string) (*AppUser, error) {
	appUserUuidKey := fmt.Sprintf("%s%v", cacheAppUserUuidPrefix, uuid)
	var resp AppUser
	err := m.QueryRowIndexCtx(ctx, &resp, appUserUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", appUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uuid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppUserModel) Insert(ctx context.Context, data *AppUser) (sql.Result, error) {
	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, data.Id)
	appUserUuidKey := fmt.Sprintf("%s%v", cacheAppUserUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.AppleUid, data.GoogleUid, data.Email, data.Mobile, data.MobileCountryId, data.MobileCountryNum, data.Locale, data.UserType, data.Password, data.InvitationCode, data.DeletedAt)
	}, appUserIdKey, appUserUuidKey)
	return ret, err
}

func (m *defaultAppUserModel) Update(ctx context.Context, newData *AppUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	appUserIdKey := fmt.Sprintf("%s%v", cacheAppUserIdPrefix, data.Id)
	appUserUuidKey := fmt.Sprintf("%s%v", cacheAppUserUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Uuid, newData.AppleUid, newData.GoogleUid, newData.Email, newData.Mobile, newData.MobileCountryId, newData.MobileCountryNum, newData.Locale, newData.UserType, newData.Password, newData.InvitationCode, newData.DeletedAt, newData.Id)
	}, appUserIdKey, appUserUuidKey)
	return err
}

func (m *defaultAppUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAppUserIdPrefix, primary)
}

func (m *defaultAppUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAppUserModel) tableName() string {
	return m.table
}
