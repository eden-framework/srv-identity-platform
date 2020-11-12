package permissions

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
	Router.Register(courier.NewRouter(UpdatePermissionByPermissionID{}))
}

// 根据PermissionID更新Permission
type UpdatePermissionByPermissionID struct {
	httpx.MethodPatch

	PermissionID uint64                               `name:"permissionID,string" in:"path"`
	Data         service.ModulePermissionUpdateOption `in:"body"`
}

func (req UpdatePermissionByPermissionID) Path() string {
	return "/:permissionID"
}

func (req UpdatePermissionByPermissionID) Output(ctx context.Context) (result interface{}, err error) {
	err = service.NewController(global.Config.MasterDB.Get(), global.ClientConfig.ID).UpdateModulePermission(req.PermissionID, req.Data)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.PermissionNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
