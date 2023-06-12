package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppLanguageModel = (*customAppLanguageModel)(nil)

type (
	// AppLanguageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppLanguageModel.
	AppLanguageModel interface {
		appLanguageModel
		FindAll(ctx context.Context) ([]*AppLanguage, error)
	}

	customAppLanguageModel struct {
		*defaultAppLanguageModel
	}
)

// NewAppLanguageModel returns a model for the database table.
func NewAppLanguageModel(conn sqlx.SqlConn, c cache.CacheConf) AppLanguageModel {
	return &customAppLanguageModel{
		defaultAppLanguageModel: newAppLanguageModel(conn, c),
	}
}

func (m *customAppLanguageModel) FindAll(ctx context.Context) ([]*AppLanguage, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at is null ORDER BY id ASC", appLanguageRows, m.table)
	var resp []*AppLanguage
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
