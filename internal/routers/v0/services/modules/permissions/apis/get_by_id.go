package apis

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(GetPermissionApiByApiID{}))
}

// 根据ApiID获取PermissionApi
type GetPermissionApiByApiID struct {
	httpx.MethodGet

	ApiID uint64 `name:"apiID,string" in:"path"`
}

func (req GetPermissionApiByApiID) Path() string {
	return "/:apiID"
}

func (req GetPermissionApiByApiID) Output(ctx context.Context) (result interface{}, err error) {
	data, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModulePermissionApiByApiID(req.ApiID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.PermissionApiNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return data, err
}
