package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppUserPhotoModel = (*customAppUserPhotoModel)(nil)

type (
	// AppUserPhotoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserPhotoModel.
	AppUserPhotoModel interface {
		appUserPhotoModel
		InsertByList(ctx context.Context, list []*AppUserPhoto) (sql.Result, error)
		FindByUserId(ctx context.Context, userId int64) ([]*AppUserPhoto, error)
	}

	customAppUserPhotoModel struct {
		*defaultAppUserPhotoModel
	}
)

// NewAppUserPhotoModel returns a model for the database table.
func NewAppUserPhotoModel(conn sqlx.SqlConn, c cache.CacheConf) AppUserPhotoModel {
	return &customAppUserPhotoModel{
		defaultAppUserPhotoModel: newAppUserPhotoModel(conn, c),
	}
}

func (m *customAppUserPhotoModel) InsertByList(ctx context.Context, list []*AppUserPhoto) (sql.Result, error) {

	res, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		err := conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			for _, item := range list {
				if item.Id > 0 {
					query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appUserPhotoRowsWithPlaceHolder)
					_, err := session.ExecCtx(ctx, query, item.UserId, item.Title, item.FileUrl, item.OrderBy, item.DeletedAt, item.Id)
					if err != nil {
						return err
					}
				} else {
					query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, appUserPhotoRowsExpectAutoSet)
					_, err := session.ExecCtx(ctx, query, item.UserId, item.Title, item.FileUrl, item.OrderBy, item.DeletedAt)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		return nil, err
	})
	return res, err
}

func (m *customAppUserPhotoModel) FindByUserId(ctx context.Context, userId int64) ([]*AppUserPhoto, error) {
	var resp []*AppUserPhoto
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by order_by asc", appUserPhotoRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
