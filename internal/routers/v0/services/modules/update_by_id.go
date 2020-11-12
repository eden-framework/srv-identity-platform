package modules

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
	Router.Register(courier.NewRouter(UpdateModuleByModuleID{}))
}

// 根据ModuleID更新Module
type UpdateModuleByModuleID struct {
	httpx.MethodPatch

	ModuleID uint64                     `name:"moduleID,string" in:"path"`
	Data     service.ModuleUpdateOption `in:"body"`
}

func (req UpdateModuleByModuleID) Path() string {
	return "/:moduleID"
}

func (req UpdateModuleByModuleID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).UpdateModule(req.ModuleID, req.Data)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.ServiceNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
