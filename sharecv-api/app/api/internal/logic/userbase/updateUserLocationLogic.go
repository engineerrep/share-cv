package userbase

import (
	"context"
	"sharecvapi/app/api/internal/svc"
	"sharecvapi/app/api/internal/types"
	"sharecvapi/app/model"
	"sharecvapi/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLocationLogic {
	return &UpdateUserLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLocationLogic) UpdateUserLocation(req *types.UpdateUserLocationReq) (resp *types.StatusResp, err error) {
	userId := ctxdata.GetUserIdFromCtx(l.ctx)
	var isAdd = false
	userLocation, err := l.svcCtx.AppUserLocationModel.FindOne(l.ctx, userId)
	if userLocation == nil {
		isAdd = true
		userLocation = &model.AppUserLocation{UserId: userId}
	}
	if len(req.TimeZone) > 0 {
		userLocation.TimeZone = req.TimeZone
	}
	if len(req.Country) > 0 {
		userLocation.Country = req.Country
	}
	if len(req.CountryCode) > 0 {
		userLocation.CountryCode = req.CountryCode
	}
	if len(req.Province) > 0 {
		userLocation.Province = req.Province
	}
	if len(req.City) > 0 {
		userLocation.City = req.City
	}
	if len(req.State) > 0 {
		userLocation.State = req.State
	}
	if len(req.Area) > 0 {
		userLocation.Area = req.Area
	}
	if len(req.Street) > 0 {
		userLocation.Street = req.Street
	}
	if req.Latitude != 0 {
		userLocation.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		userLocation.Longitude = req.Longitude
	}
	if len(req.PostalCode) > 0 {
		userLocation.PostalCode = req.PostalCode
	}
	if len(req.Address) > 0 {
		userLocation.Address = req.Address
	}
	if isAdd {
		_, err = l.svcCtx.AppUserLocationModel.Insert(l.ctx, userLocation)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	} else {
		err = l.svcCtx.AppUserLocationModel.Update(l.ctx, userLocation)
		if err != nil {
			return &types.StatusResp{Status: 0}, nil
		}
	}

	return &types.StatusResp{Status: 1}, nil
}
