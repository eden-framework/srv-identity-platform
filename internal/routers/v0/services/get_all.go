package services

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
	Router.Register(courier.NewRouter(GetServices{}))
}

// 批量获取Service
type GetServices struct {
	httpx.MethodGet
	service.ServiceCondition
	// 分页偏移量
	Offset int64 `name:"offset" default:"0" in:"query"`
	// 单页数据量
	Limit int64 `name:"limit" default:"10" in:"query"`
}

func (req GetServices) Path() string {
	return ""
}

type FetchServicesResult struct {
	Data  []databases.Services `json:"data"`
	Total int                  `json:"total"`
}

func (req GetServices) Output(ctx context.Context) (result interface{}, err error) {
	svcs, count, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetServices(req.ServiceCondition, req.Offset, req.Limit)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}
	result = FetchServicesResult{
		Data:  svcs,
		Total: count,
	}
	return result, nil
}
