package apis

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
	Router.Register(courier.NewRouter(CreatePermissionApi{}))
}

// 创建权限策略
type CreatePermissionApi struct {
	httpx.MethodPost
	// PermissionID
	PermissionID uint64            `name:"permissionID,string" in:"path"`
	Body         databases.BaseApi `in:"body"`
}

func (req CreatePermissionApi) Path() string {
	return ""
}

func (req CreatePermissionApi) Output(ctx context.Context) (result interface{}, err error) {
	_, err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).CreateModulePermissionApi(req.PermissionID, req.Body)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return nil, err
}
