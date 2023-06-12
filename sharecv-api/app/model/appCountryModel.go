package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppCountryModel = (*customAppCountryModel)(nil)

type (
	// AppCountryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppCountryModel.
	AppCountryModel interface {
		appCountryModel
		FindAll(ctx context.Context) ([]*AppCountry, error)
	}

	customAppCountryModel struct {
		*defaultAppCountryModel
	}
)

// NewAppCountryModel returns a model for the database table.
func NewAppCountryModel(conn sqlx.SqlConn, c cache.CacheConf) AppCountryModel {
	return &customAppCountryModel{
		defaultAppCountryModel: newAppCountryModel(conn, c),
	}
}

func (m *customAppCountryModel) FindAll(ctx context.Context) ([]*AppCountry, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at is null ORDER BY id ASC", appCountryRows, m.table)
	var resp []*AppCountry
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
