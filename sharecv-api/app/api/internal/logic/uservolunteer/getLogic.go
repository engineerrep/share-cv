package uservolunteer

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.GetUserVolunteerReq) (resp *types.UserVolunteer, err error) {
	entity, err := l.svcCtx.AppUserVolunteerModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	var val types.UserVolunteer
	err = copier.Copy(&val, &entity)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	return &val, nil
}
