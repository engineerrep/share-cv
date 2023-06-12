package userbase

import (
	"context"
	"fmt"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPhoneLogic {
	return &UpdateUserPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPhoneLogic) UpdateUserPhone(req *types.UpdateUserPhoneReq) (resp *types.StatusResp, err error) {

	phone := fmt.Sprintf("+%s%s", req.PhoneNum, req.Phone)
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, phone)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	userID := ctxdata.GetUserIdFromCtx(l.ctx)
	user, err := l.svcCtx.AppUserModel.FindOne(l.ctx, userID)
	if err != nil {
		return &types.StatusResp{Status: 0}, nil
	}

	user.MobileCountryNum = req.PhoneNum
	user.Mobile = req.Phone
	user.MobileCountryId = req.CountryId
	err = l.svcCtx.AppUserModel.Update(l.ctx, user)
	if err != nil {
		return &types.StatusResp{Status: 0}, nil
	}

	return &types.StatusResp{Status: 1}, nil
}
