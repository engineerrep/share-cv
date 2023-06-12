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

type CheckPhoneCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckPhoneCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPhoneCodeLogic {
	return &CheckPhoneCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckPhoneCodeLogic) CheckPhoneCode(req *types.CheckCodePhoneReq) (resp *types.StatusResp, err error) {
	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	phone := fmt.Sprintf("+%s%s", req.PhoneNum, req.Phone)
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, phone)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
