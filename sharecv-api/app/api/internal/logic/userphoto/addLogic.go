package userphoto

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

func (l *AddLogic) Add(req *types.AddUserPhotoReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	if len(req.List) == 0 {
		return nil, errorx.NewCodeError(errorx.ParamErrorCode)
	}

	list, err := l.svcCtx.AppUserPhotoModel.FindByUserId(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	if len(list) >= l.svcCtx.Config.User.MaxPhoto {
		return nil, errorx.NewCodeError(errorx.MaxCountErrorCode)
	}

	var entitys []*model.AppUserPhoto
	for _, item := range req.List {
		var val model.AppUserPhoto
		err = copier.Copy(&val, &item)
		if err != nil {
			continue
		}
		val.UserId = userId
		entitys = append(entitys, &val)
	}
	_, err = l.svcCtx.AppUserPhotoModel.InsertByList(l.ctx, entitys)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.SaveErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
