package userexperience

import (
	"context"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteUserExperienceReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	entity, err := l.svcCtx.AppUserExperienceModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	if entity.UserId != userId {
		return &types.StatusResp{Status: 0}, nil
	}

	err = l.svcCtx.AppUserExperienceModel.Delete(l.ctx, req.Id)
	if err != nil {
		return &types.StatusResp{Status: 0}, nil
	}

	return &types.StatusResp{Status: 1}, nil
}
