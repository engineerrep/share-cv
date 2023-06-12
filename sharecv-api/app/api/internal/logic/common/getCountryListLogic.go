package common

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCountryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountryListLogic {
	return &GetCountryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCountryListLogic) GetCountryList() (resp *types.CountryListResp, err error) {
	list, err := l.svcCtx.AppCountryModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	var array []types.Country
	for _, item := range list {
		var val types.Country
		err = copier.Copy(&val, &item)
		if err != nil {
			return nil, errorx.NewSystemError()
		}

		array = append(array, val)
	}
	return &types.CountryListResp{
		List: array,
	}, nil
}
