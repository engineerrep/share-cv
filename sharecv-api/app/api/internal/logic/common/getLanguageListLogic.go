package common

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageListLogic {
	return &GetLanguageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageListLogic) GetLanguageList() (resp *types.LanguageListResp, err error) {
	list, err := l.svcCtx.AppLanguageModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	var array []types.Language
	for _, item := range list {
		var val types.Language
		err = copier.Copy(&val, &item)
		if err != nil {
			return nil, errorx.NewSystemError()
		}

		array = append(array, val)
	}

	return &types.LanguageListResp{
		List: array,
	}, nil
}
