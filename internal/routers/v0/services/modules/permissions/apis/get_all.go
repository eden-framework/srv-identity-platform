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
	Router.Register(courier.NewRouter(GetPermissionApis{}))
}

// 批量获取权限策略的API
type GetPermissionApis struct {
	httpx.MethodGet
	service.PermissionApiCondition
	// 分页偏移量
	Offset int64 `name:"offset" default:"0" in:"query"`
	// 单页数据量
	Limit int64 `name:"limit" default:"10" in:"query"`
}

func (req GetPermissionApis) Path() string {
	return ""
}

type GetPermissionApisResult struct {
	Data  []databases.ModulePermissionApi `json:"data"`
	Total int                             `json:"total"`
}

func (req GetPermissionApis) Output(ctx context.Context) (result interface{}, err error) {
	data, count, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModulePermissionApi(req.PermissionApiCondition, req.Offset, req.Limit)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}
	result = GetPermissionApisResult{
		Data:  data,
		Total: count,
	}
	return result, nil
}
