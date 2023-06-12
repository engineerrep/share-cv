package userbase

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfo, err error) {
	userID := ctxdata.GetUserIdFromCtx(l.ctx)
	user, err := l.svcCtx.AppUserModel.FindOne(l.ctx, userID)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	userInfo, err := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, userID)
	if err != nil {
		return &types.UserInfo{
			UserId: user.Uuid,
		}, nil
	}

	var val types.UserInfo
	err = copier.Copy(&val, &userInfo)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	val.UserId = user.Uuid

	return &val, nil
}
