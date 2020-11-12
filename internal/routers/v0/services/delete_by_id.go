package services

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(DeleteServiceByServiceID{}))
}

// 根据ServiceID删除Service
type DeleteServiceByServiceID struct {
	httpx.MethodDelete
	ServiceID uint64 `name:"serviceID,string" in:"path"`
}

func (req DeleteServiceByServiceID) Path() string {
	return "/:serviceID"
}

func (req DeleteServiceByServiceID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).DeleteService(req.ServiceID, true)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
