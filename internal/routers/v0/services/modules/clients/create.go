package clients

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
	Router.Register(courier.NewRouter(CreateClient{}))
}

// 创建客户端
type CreateClient struct {
	httpx.MethodPost
	// 模块ID
	ModuleID uint64               `name:"moduleID,string" in:"path"`
	Body     databases.BaseClient `in:"body"`
}

func (req CreateClient) Path() string {
	return ""
}

func (req CreateClient) Output(ctx context.Context) (result interface{}, err error) {
	_, err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).CreateModuleClient(req.ModuleID, req.Body)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return nil, err
}
