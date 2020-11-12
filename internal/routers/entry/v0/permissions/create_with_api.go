package permissions

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(CreatePermissionWithApi{}))
}

type CreatePermissionRequest struct {
	ModuleID uint64 `json:"moduleID,string"`
	databases.BasePermission
	Apis []databases.BaseApi `json:"apis,omitempty"`
}

// 创建权限策略与对应的API
type CreatePermissionWithApi struct {
	httpx.MethodPost
	Body CreatePermissionRequest `in:"body"`
}

func (req CreatePermissionWithApi) Path() string {
	return ""
}

func (req CreatePermissionWithApi) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB.Get())
	err = tx.With(func(db sqlx.DBExecutor) error {
		controller := service.NewController(db, global.ClientConfig.ID)
		permission, err := controller.CreateModulePermission(req.Body.ModuleID, req.Body.BasePermission)
		if err != nil {
			return err
		}

		for _, api := range req.Body.Apis {
			_, err := controller.CreateModulePermissionApi(permission.PermissionID, api)
			if err != nil {
				return err
			}
		}
		return nil
	}).Do()

	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return
}
