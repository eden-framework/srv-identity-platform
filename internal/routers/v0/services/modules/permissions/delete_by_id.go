package permissions

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(DeletePermissionByPermissionID{}))
}

// 根据PermissionID删除Permission
type DeletePermissionByPermissionID struct {
	httpx.MethodDelete
	PermissionID uint64 `name:"permissionID,string" in:"path"`
}

func (req DeletePermissionByPermissionID) Path() string {
	return "/:permissionID"
}

func (req DeletePermissionByPermissionID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).DeleteModulePermission(req.PermissionID, true)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
