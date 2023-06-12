package usercompany

import (
	"context"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdJobSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdJobSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdJobSkillsLogic {
	return &UpdJobSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdJobSkillsLogic) UpdJobSkills(req *types.UpdUserCompanySkillsReq) (resp *types.StatusResp, err error) {
	entity, err := l.svcCtx.AppUserCompanyModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	if entity.UserId != userId {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	entity.JobSkills = req.JobSkills
	err = l.svcCtx.AppUserCompanyModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ChangeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
