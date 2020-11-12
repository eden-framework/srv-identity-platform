package clients

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(DeleteClientByClientID{}))
}

// 根据ClientID删除Client
type DeleteClientByClientID struct {
	httpx.MethodDelete
	ClientID uint64 `name:"clientID,string" in:"path"`
}

func (req DeleteClientByClientID) Path() string {
	return "/:clientID"
}

func (req DeleteClientByClientID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).DeleteModuleClient(req.ClientID, true)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
