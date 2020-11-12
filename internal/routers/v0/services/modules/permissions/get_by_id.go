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
	Router.Register(courier.NewRouter(GetPermissionByPermissionID{}))
}

// 根据PermissionID获取Permission
type GetPermissionByPermissionID struct {
	httpx.MethodGet

	PermissionID uint64 `name:"permissionID,string" in:"path"`
}

func (req GetPermissionByPermissionID) Path() string {
	return "/:permissionID"
}

func (req GetPermissionByPermissionID) Output(ctx context.Context) (result interface{}, err error) {
	data, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModulePermissionByPermissionID(req.PermissionID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.PermissionNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return data, err
}
