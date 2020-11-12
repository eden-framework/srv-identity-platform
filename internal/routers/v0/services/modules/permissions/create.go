package permissions

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
	Router.Register(courier.NewRouter(CreatePermission{}))
}

// 创建权限策略
type CreatePermission struct {
	httpx.MethodPost
	// ModuleID
	ModuleID uint64                   `name:"moduleID,string" in:"path"`
	Body     databases.BasePermission `in:"body"`
}

func (req CreatePermission) Path() string {
	return ""
}

func (req CreatePermission) Output(ctx context.Context) (result interface{}, err error) {
	_, err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).CreateModulePermission(req.ModuleID, req.Body)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return nil, err
}
