package userbase

import (
	"context"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFeedBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFeedBackLogic {
	return &UserFeedBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFeedBackLogic) UserFeedBack(req *types.UserFeedbackReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	item := &model.AppUserFeedback{
		UserId:   userId,
		Content:  req.Content,
		FileUrls: req.FileUrls,
	}
	_, err = l.svcCtx.AppUserFeedbackModel.Insert(l.ctx, item)
	if err != nil {
		return &types.StatusResp{Status: 0}, nil
	}

	return &types.StatusResp{Status: 1}, nil
}
