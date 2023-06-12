package usercompany

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddUserCompanyReq) (resp *types.IdResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)

	var val model.AppUserCompany
	err = copier.Copy(&val, &req)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ParamErrorCode)
	}

	val.UserId = userId

	if val.Id > 0 {
		err = l.svcCtx.AppUserCompanyModel.Update(l.ctx, &val)
		if err != nil {
			return &types.IdResp{Id: 0}, nil
		}

		return &types.IdResp{Id: val.Id}, nil
	}

	res, err := l.svcCtx.AppUserCompanyModel.Insert(l.ctx, &val)
	if err != nil {
		return &types.IdResp{Id: 0}, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return &types.IdResp{Id: 0}, nil
	}

	return &types.IdResp{Id: id}, nil
}
