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
	appUserLocationFieldNames          = builder.RawFieldNames(&AppUserLocation{})
	appUserLocationRows                = strings.Join(appUserLocationFieldNames, ",")
	appUserLocationRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserLocationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	appUserLocationRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserLocationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAppUserLocationIdPrefix = "cache:appUserLocation:id:"
)

type (
	appUserLocationModel interface {
		Insert(ctx context.Context, data *AppUserLocation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AppUserLocation, error)
		Update(ctx context.Context, data *AppUserLocation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAppUserLocationModel struct {
		sqlc.CachedConn
		table string
	}

	AppUserLocation struct {
		Id          int64        `db:"id"`
		UserId      int64        `db:"user_id"`      // app_user表id
		TimeZone    string       `db:"time_zone"`    // 时区
		Country     string       `db:"country"`      // 国家
		CountryCode string       `db:"country_code"` // 国家代码
		Province    string       `db:"province"`     // 省份
		City        string       `db:"city"`         // 城市
		State       string       `db:"state"`        // 县
		Area        string       `db:"area"`         // 区域
		Street      string       `db:"street"`       // 街道
		Latitude    float64      `db:"latitude"`     // 纬度
		Longitude   float64      `db:"longitude"`    // 经度
		PostalCode  string       `db:"postal_code"`  // 邮政编号
		Address     string       `db:"address"`      // 详细地点
		CreateAt    time.Time    `db:"create_at"`    // 创建时间
		UpdateAt    time.Time    `db:"update_at"`    // 更新时间
		DeletedAt   sql.NullTime `db:"deleted_at"`   // 删除时间
	}
)

func newAppUserLocationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAppUserLocationModel {
	return &defaultAppUserLocationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`app_user_location`",
	}
}

func (m *defaultAppUserLocationModel) Delete(ctx context.Context, id int64) error {
	appUserLocationIdKey := fmt.Sprintf("%s%v", cacheAppUserLocationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, appUserLocationIdKey)
	return err
}

func (m *defaultAppUserLocationModel) FindOne(ctx context.Context, id int64) (*AppUserLocation, error) {
	appUserLocationIdKey := fmt.Sprintf("%s%v", cacheAppUserLocationIdPrefix, id)
	var resp AppUserLocation
	err := m.QueryRowCtx(ctx, &resp, appUserLocationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserLocationRows, m.table)
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

func (m *defaultAppUserLocationModel) Insert(ctx context.Context, data *AppUserLocation) (sql.Result, error) {
	appUserLocationIdKey := fmt.Sprintf("%s%v", cacheAppUserLocationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appUserLocationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.TimeZone, data.Country, data.CountryCode, data.Province, data.City, data.State, data.Area, data.Street, data.Latitude, data.Longitude, data.PostalCode, data.Address, data.DeletedAt)
	}, appUserLocationIdKey)
	return ret, err
}

func (m *defaultAppUserLocationModel) Update(ctx context.Context, data *AppUserLocation) error {
	appUserLocationIdKey := fmt.Sprintf("%s%v", cacheAppUserLocationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserLocationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.TimeZone, data.Country, data.CountryCode, data.Province, data.City, data.State, data.Area, data.Street, data.Latitude, data.Longitude, data.PostalCode, data.Address, data.DeletedAt, data.Id)
	}, appUserLocationIdKey)
	return err
}

func (m *defaultAppUserLocationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAppUserLocationIdPrefix, primary)
}

func (m *defaultAppUserLocationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserLocationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAppUserLocationModel) tableName() string {
	return m.table
}
