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
	Router.Register(courier.NewRouter(GetServiceByServiceID{}))
}

// 根据ServiceID获取Service
type GetServiceByServiceID struct {
	httpx.MethodGet

	ServiceID uint64 `name:"serviceID,string" in:"path"`
}

func (req GetServiceByServiceID) Path() string {
	return "/:serviceID"
}

func (req GetServiceByServiceID) Output(ctx context.Context) (result interface{}, err error) {
	svc, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetServiceByServiceID(req.ServiceID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.ServiceNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return svc, err
}
