package userbase

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOtherUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOtherUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOtherUserInfoLogic {
	return &GetOtherUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOtherUserInfoLogic) GetOtherUserInfo(req *types.GetOtherUserInfoReq) (resp *types.OtherUserInfo, err error) {

	user, err := l.svcCtx.AppUserModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	userInfo, err := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, user.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	// 已经隐藏个人资料
	if userInfo.ModeType != 1 {
		return nil, errorx.NewCodeError(errorx.AccessErrorCode)
	}

	var val types.OtherUserInfo
	err = copier.Copy(&val, &userInfo)
	if err != nil {
		return nil, errorx.NewSystemError()
	}

	return &val, nil
}
