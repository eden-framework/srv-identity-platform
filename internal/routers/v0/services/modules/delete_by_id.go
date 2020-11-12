package modules

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(DeleteModuleByModuleID{}))
}

// 根据ModuleID删除Module
type DeleteModuleByModuleID struct {
	httpx.MethodDelete
	ModuleID uint64 `name:"moduleID,string" in:"path"`
}

func (req DeleteModuleByModuleID) Path() string {
	return "/:moduleID"
}

func (req DeleteModuleByModuleID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).DeleteModule(req.ModuleID, true)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
