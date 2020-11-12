package modules

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(CreateModule{}))
}

// 创建模块
type CreateModule struct {
	httpx.MethodPost
	// ServiceID
	ServiceID uint64               `name:"serviceID,string" in:"path"`
	Body      databases.BaseModule `in:"body"`
}

func (req CreateModule) Path() string {
	return ""
}

func (req CreateModule) Output(ctx context.Context) (result interface{}, err error) {
	_, err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).CreateModule(req.ServiceID, req.Body)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return nil, err
}
