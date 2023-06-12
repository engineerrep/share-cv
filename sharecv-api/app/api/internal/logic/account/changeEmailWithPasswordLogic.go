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

type ChangeEmailWithPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEmailWithPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailWithPasswordLogic {
	return &ChangeEmailWithPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEmailWithPasswordLogic) ChangeEmailWithPassword(req *types.ChangeEmailWithPasswordReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	user, err := l.svcCtx.AppUserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	salt := fmt.Sprintf("%s%s", req.OldPassword, l.svcCtx.Config.Salt)
	pwd := utils.MD5(salt)
	if user.Password != pwd {
		return nil, errorx.NewCodeError(errorx.PasswordErrorCode)
	}

	newSalt := fmt.Sprintf("%s%s", req.NewPassword, l.svcCtx.Config.Salt)
	newPwd := utils.MD5(newSalt)
	user.Password = newPwd

	err = l.svcCtx.AppUserModel.Update(l.ctx, user)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ChangeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
