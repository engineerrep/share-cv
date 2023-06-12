package userbase

import (
	"context"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserNameLogic {
	return &UpdateUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserNameLogic) UpdateUserName(req *types.UpdateUserNameReq) (resp *types.StatusResp, err error) {
	userID := ctxdata.GetUserIdFromCtx(l.ctx)
	userInfo, err := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, userID)
	if userInfo == nil {
		userInfo = &model.AppUserInfo{UserId: userID, FirstName: req.FirstName, LastName: req.LastName}
		_, err = l.svcCtx.AppUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	} else {
		userInfo.FirstName = req.FirstName
		userInfo.LastName = req.LastName
		err = l.svcCtx.AppUserInfoModel.Update(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	}

	return &types.StatusResp{Status: 1}, nil
}
