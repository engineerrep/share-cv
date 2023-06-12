package usercustomattribute

import (
	"context"
	"github.com/jinzhu/copier"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/common/ctxdata"
	"sharecvapi/common/errorx"

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

func (l *ListLogic) List() (resp *types.UserCustomAttributeResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	list, err := l.svcCtx.AppUserCustomAttributesModel.FindByUserId(l.ctx, userId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.NotFountErrorCode)
	}

	var array []types.UserCustomAttribute
	for _, item := range list {
		var val types.UserCustomAttribute
		err = copier.Copy(&val, &item)
		if err != nil {
			continue
		}
		array = append(array, val)
	}

	return &types.UserCustomAttributeResp{List: array}, nil
}
