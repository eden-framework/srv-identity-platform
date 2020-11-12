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
	Router.Register(courier.NewRouter(GetPermissions{}))
}

// 批量获取权限策略
type GetPermissions struct {
	httpx.MethodGet
	service.ModulePermissionCondition
	// 分页偏移量
	Offset int64 `name:"offset" default:"0" in:"query"`
	// 单页数据量
	Limit int64 `name:"limit" default:"10" in:"query"`
}

func (req GetPermissions) Path() string {
	return ""
}

type GetPermissionsResult struct {
	Data  []databases.ModulePermissions `json:"data"`
	Total int                           `json:"total"`
}

func (req GetPermissions) Output(ctx context.Context) (result interface{}, err error) {
	data, count, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModulePermissions(req.ModulePermissionCondition, req.Offset, req.Limit)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}
	result = GetPermissionsResult{
		Data:  data,
		Total: count,
	}
	return result, nil
}
