package register

import (
	"context"
	"sharecvapi/common/errorx"
	"sharecvapi/common/verify"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailLogic {
	return &CheckEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckEmailLogic) CheckEmail(req *types.CheckEmailReq) (resp *types.StatusResp, err error) {

	// 谷歌人机验证
	err = verify.GoogleVerifyReCaptchaV3(req.GoogleCaptcha, l.svcCtx.Config.Google.Recaptcha)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.TokenErrorCode)
	}

	// 邮箱是否存在
	user, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.Email)
	if user != nil {
		return nil, errorx.NewCodeError(errorx.EmailExistCode)
	}

	return &types.StatusResp{
		Status: 1,
	}, nil
}
