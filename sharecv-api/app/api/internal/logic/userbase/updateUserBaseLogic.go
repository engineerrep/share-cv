package userbase

import (
	"context"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"

	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBaseLogic {
	return &UpdateUserBaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBaseLogic) UpdateUserBase(req *types.UpdateUserBaseReq) (resp *types.StatusResp, err error) {
	userID := ctxdata.GetUserIdFromCtx(l.ctx)
	userInfo, err := l.svcCtx.AppUserInfoModel.FindOne(l.ctx, userID)
	if userInfo == nil {
		userInfo = &model.AppUserInfo{UserId: userID,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Birthday:    req.Birthday,
			CountryId:   req.CountryId,
			CountryName: req.CountryName,
			Province:    req.Province,
			City:        req.City,
			Address:     req.Address,
		}
		_, err = l.svcCtx.AppUserInfoModel.Insert(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	} else {
		if len(req.FirstName) > 0 {
			userInfo.FirstName = req.FirstName
		}
		if len(req.LastName) > 0 {
			userInfo.LastName = req.LastName
		}
		if len(req.Birthday) > 0 {
			userInfo.Birthday = req.Birthday
		}
		if req.CountryId > 0 {
			userInfo.CountryId = req.CountryId
		}
		if len(req.CountryName) > 0 {
			userInfo.CountryName = req.CountryName
		}
		if len(req.Province) > 0 {
			userInfo.Province = req.Province
		}
		if len(req.City) > 0 {
			userInfo.City = req.City
		}
		if len(req.Address) > 0 {
			userInfo.Address = req.Address
		}
		err = l.svcCtx.AppUserInfoModel.Update(l.ctx, userInfo)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	}

	return &types.StatusResp{Status: 1}, nil
}
