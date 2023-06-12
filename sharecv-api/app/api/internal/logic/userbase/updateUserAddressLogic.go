package userbase

import (
	"context"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(req *types.UpdateUserAddressReq) (resp *types.StatusResp, err error) {
	userID := ctxdata.GetUserIdFromCtx(l.ctx)
	userInfo, err := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, userID)
	if userInfo == nil {
		userInfo = &model.AppUserInfo{UserId: userID, Address: req.Address}
		_, err = l.svcCtx.AppUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	} else {
		userInfo.Address = req.Address
		err = l.svcCtx.AppUserInfoModel.Update(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	}

	return &types.StatusResp{Status: 1}, nil
}
