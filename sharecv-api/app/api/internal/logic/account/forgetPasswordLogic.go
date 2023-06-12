package account

import (
	"context"
	"fmt"
	"sharecvapi/common/errorx"
	"sharecvapi/common/globalkey"
	"sharecvapi/common/utils"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(req *types.ForgetPasswordReq) (resp *types.StatusResp, err error) {
	cacheCodeKey := fmt.Sprintf(globalkey.CacheCodeKey, req.Email)
	verifyCode, _ := l.svcCtx.Redis.Get(cacheCodeKey)
	if verifyCode != req.Code {
		return nil, errorx.NewCodeError(errorx.CodeErrorCode)
	}

	user, err := l.svcCtx.AppUserModel.FindByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	// 密码处理
	salt := fmt.Sprintf("%s%s", req.Password, l.svcCtx.Config.Salt)
	pwd := utils.MD5(salt)
	user.Password = pwd
	// 更改用户密码
	err = l.svcCtx.AppUserModel.Update(l.ctx, user)
	var status int64 = 0
	if err == nil {
		// 删除验证码缓存
		l.svcCtx.Redis.DelCtx(l.ctx, cacheCodeKey)
		status = 1
	}
	return &types.StatusResp{
		Status: status,
	}, nil
}
