package account

import (
	"context"
	"fmt"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeEmailWithCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEmailWithCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailWithCodeLogic {
	return &ChangeEmailWithCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEmailWithCodeLogic) ChangeEmailWithCode(req *types.ChangeEmailWithCodeReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	user, err := l.svcCtx.AppUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.NewEmail)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	l.svcCtx.Redis.Del(cacheCodeKey)

	findUser, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.NewEmail)
	if findUser != nil {
		return nil, errorx.NewCodeError(errorx.EmailExistCode)
	}

	// 更改当前邮箱为新邮箱
	user.Email = req.NewEmail
	err = l.svcCtx.AppUserModel.Update(l.ctx, user)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ChangeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
