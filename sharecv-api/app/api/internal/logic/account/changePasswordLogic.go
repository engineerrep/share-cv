package account

import (
	"context"
	"fmt"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"
	"sharecvapi/common/utils"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	user, err := l.svcCtx.AppUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	salt := fmt.Sprintf("%s%s", req.OldPassword, l.svcCtx.Config.Salt)
	oldPwd := utils.MD5(salt)
	if user.Password != oldPwd {
		return nil, errorx.NewCodeError(errorx.PasswordErrorCode)
	}

	saltNew := fmt.Sprintf("%s%s", req.NewPassword, l.svcCtx.Config.Salt)
	newPwd := utils.MD5(saltNew)
	user.Password = newPwd
	err = l.svcCtx.AppUserModel.Update(l.ctx, user)
	if err != nil {
		return &types.StatusResp{Status: 0}, nil
	}

	return &types.StatusResp{Status: 1}, nil
}
