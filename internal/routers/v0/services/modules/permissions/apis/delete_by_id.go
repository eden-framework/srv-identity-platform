package apis

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(DeletePermissionApiByApiID{}))
}

// 根据ApiID删除Api
type DeletePermissionApiByApiID struct {
	httpx.MethodDelete
	ApiID uint64 `name:"apiID,string" in:"path"`
}

func (req DeletePermissionApiByApiID) Path() string {
	return "/:apiID"
}

func (req DeletePermissionApiByApiID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).DeleteModulePermissionApi(req.ApiID, true)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
