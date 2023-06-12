package account

import (
	"context"
	"fmt"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"
	"sharecvapi/common/verify"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailCodeLogic {
	return &CheckEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckEmailCodeLogic) CheckEmailCode(req *types.CheckCodeEmailReq) (resp *types.StatusResp, err error) {
	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.Email)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
