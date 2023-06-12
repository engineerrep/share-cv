package userproject

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.UserProjectResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	list, err := l.svcCtx.AppUserProjectModel.FindByUserId(l.ctx, userId)

	var array []types.UserProject
	if err != nil {
		return &types.UserProjectResp{List: array}, nil
	}

	for _, item := range list {
		var val types.UserProject
		err = copier.Copy(&val, &item)
		if err != nil {
			continue
		}
		array = append(array, val)
	}

	return &types.UserProjectResp{List: array}, nil
}
