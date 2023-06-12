package userproject

import (
	"context"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdProjectSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdProjectSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdProjectSkillsLogic {
	return &UpdProjectSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdProjectSkillsLogic) UpdProjectSkills(req *types.UpdUserProjectSkillsReq) (resp *types.StatusResp, err error) {
	entity, err := l.svcCtx.AppUserProjectModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	if entity.UserId != userId {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	entity.ProjectSkills = req.ProjectSkills
	err = l.svcCtx.AppUserProjectModel.Update(l.ctx, entity)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ChangeErrorCode)
	}

	return &types.StatusResp{Status: 1}, nil
}
