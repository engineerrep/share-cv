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
	appCountryFieldNames          = builder.RawFieldNames(&AppCountry{})
	appCountryRows                = strings.Join(appCountryFieldNames, ",")
	appCountryRowsExpectAutoSet   = strings.Join(stringx.Remove(appCountryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	appCountryRowsWithPlaceHolder = strings.Join(stringx.Remove(appCountryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAppCountryIdPrefix = "cache:appCountry:id:"
)

type (
	appCountryModel interface {
		Insert(ctx context.Context, data *AppCountry) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AppCountry, error)
		Update(ctx context.Context, data *AppCountry) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAppCountryModel struct {
		sqlc.CachedConn
		table string
	}

	AppCountry struct {
		Id        int64        `db:"id"`         // 编号
		Num       string       `db:"num"`        // 国家区号
		Code      string       `db:"code"`       // 国家图标
		Name      string       `db:"name"`       // 国家英文名称
		NameCn    string       `db:"name_cn"`    // 国家中文名称
		CreateAt  time.Time    `db:"create_at"`  // 创建时间
		UpdateAt  time.Time    `db:"update_at"`  // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func newAppCountryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAppCountryModel {
	return &defaultAppCountryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`app_country`",
	}
}

func (m *defaultAppCountryModel) Delete(ctx context.Context, id int64) error {
	appCountryIdKey := fmt.Sprintf("%s%v", cacheAppCountryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, appCountryIdKey)
	return err
}

func (m *defaultAppCountryModel) FindOne(ctx context.Context, id int64) (*AppCountry, error) {
	appCountryIdKey := fmt.Sprintf("%s%v", cacheAppCountryIdPrefix, id)
	var resp AppCountry
	err := m.QueryRowCtx(ctx, &resp, appCountryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appCountryRows, m.table)
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

func (m *defaultAppCountryModel) Insert(ctx context.Context, data *AppCountry) (sql.Result, error) {
	appCountryIdKey := fmt.Sprintf("%s%v", cacheAppCountryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, appCountryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Num, data.Code, data.Name, data.NameCn, data.DeletedAt)
	}, appCountryIdKey)
	return ret, err
}

func (m *defaultAppCountryModel) Update(ctx context.Context, data *AppCountry) error {
	appCountryIdKey := fmt.Sprintf("%s%v", cacheAppCountryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appCountryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Num, data.Code, data.Name, data.NameCn, data.DeletedAt, data.Id)
	}, appCountryIdKey)
	return err
}

func (m *defaultAppCountryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAppCountryIdPrefix, primary)
}

func (m *defaultAppCountryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appCountryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAppCountryModel) tableName() string {
	return m.table
}