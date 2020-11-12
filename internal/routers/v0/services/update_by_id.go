package services

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
	Router.Register(courier.NewRouter(UpdateServiceByServiceID{}))
}

// 根据ServiceID更新Service
type UpdateServiceByServiceID struct {
	httpx.MethodPatch

	ServiceID uint64                      `name:"serviceID,string" in:"path"`
	Data      service.ServiceUpdateOption `in:"body"`
}

func (req UpdateServiceByServiceID) Path() string {
	return "/:serviceID"
}

func (req UpdateServiceByServiceID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).UpdateService(req.ServiceID, req.Data)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.ServiceNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
