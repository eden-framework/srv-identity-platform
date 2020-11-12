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
	Router.Register(courier.NewRouter(UpdatePermissionApiByApiID{}))
}

// 根据ApiID更新PermissionApi
type UpdatePermissionApiByApiID struct {
	httpx.MethodPatch

	ApiID uint64                            `name:"apiID,string" in:"path"`
	Data  service.PermissionApiUpdateOption `in:"body"`
}

func (req UpdatePermissionApiByApiID) Path() string {
	return "/:apiID"
}

func (req UpdatePermissionApiByApiID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).UpdateModulePermissionApi(req.ApiID, req.Data)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.PermissionApiNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
