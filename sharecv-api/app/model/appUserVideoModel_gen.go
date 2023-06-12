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
	appUserVideoFieldNames          = builder.RawFieldNames(&AppUserVideo{})
	appUserVideoRows                = strings.Join(appUserVideoFieldNames, ",")
	appUserVideoRowsExpectAutoSet   = strings.Join(stringx.Remove(appUserVideoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	appUserVideoRowsWithPlaceHolder = strings.Join(stringx.Remove(appUserVideoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAppUserVideoIdPrefix = "cache:appUserVideo:id:"
)

type (
	appUserVideoModel interface {
		Insert(ctx context.Context, data *AppUserVideo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AppUserVideo, error)
		Update(ctx context.Context, data *AppUserVideo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAppUserVideoModel struct {
		sqlc.CachedConn
		table string
	}

	AppUserVideo struct {
		Id        int64        `db:"id"`
		UserId    int64        `db:"user_id"`    // 用户id
		Title     string       `db:"title"`      // 视频标题
		CoverUrl  string       `db:"cover_url"`  // 视频封面URL
		FileUrl   string       `db:"file_url"`   // 视频URL
		PlayTime  float64      `db:"play_time"`  // 播放时长
		VideoSize float64      `db:"video_size"` // 视频大小kb
		OrderBy   int64        `db:"order_by"`   // 排序从小到大
		CreateAt  time.Time    `db:"create_at"`  // 创建时间
		UpdateAt  time.Time    `db:"update_at"`  // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func newAppUserVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAppUserVideoModel {
	return &defaultAppUserVideoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`app_user_video`",
	}
}

func (m *defaultAppUserVideoModel) Delete(ctx context.Context, id int64) error {
	appUserVideoIdKey := fmt.Sprintf("%s%v", cacheAppUserVideoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, appUserVideoIdKey)
	return err
}

func (m *defaultAppUserVideoModel) FindOne(ctx context.Context, id int64) (*AppUserVideo, error) {
	appUserVideoIdKey := fmt.Sprintf("%s%v", cacheAppUserVideoIdPrefix, id)
	var resp AppUserVideo
	err := m.QueryRowCtx(ctx, &resp, appUserVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserVideoRows, m.table)
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

func (m *defaultAppUserVideoModel) Insert(ctx context.Context, data *AppUserVideo) (sql.Result, error) {
	appUserVideoIdKey := fmt.Sprintf("%s%v", cacheAppUserVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, appUserVideoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Title, data.CoverUrl, data.FileUrl, data.PlayTime, data.VideoSize, data.OrderBy, data.DeletedAt)
	}, appUserVideoIdKey)
	return ret, err
}

func (m *defaultAppUserVideoModel) Update(ctx context.Context, data *AppUserVideo) error {
	appUserVideoIdKey := fmt.Sprintf("%s%v", cacheAppUserVideoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserVideoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Title, data.CoverUrl, data.FileUrl, data.PlayTime, data.VideoSize, data.OrderBy, data.DeletedAt, data.Id)
	}, appUserVideoIdKey)
	return err
}

func (m *defaultAppUserVideoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAppUserVideoIdPrefix, primary)
}

func (m *defaultAppUserVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appUserVideoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAppUserVideoModel) tableName() string {
	return m.table
}
