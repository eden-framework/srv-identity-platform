package clients

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
	Router.Register(courier.NewRouter(GetClientByClientID{}))
}

// 根据ClientID获取Client
type GetClientByClientID struct {
	httpx.MethodGet

	ClientID uint64 `name:"clientID,string" in:"path"`
}

func (req GetClientByClientID) Path() string {
	return "/:clientID"
}

func (req GetClientByClientID) Output(ctx context.Context) (result interface{}, err error) {
	data, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModuleClientByClientID(req.ClientID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.ClientNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return data, err
}
