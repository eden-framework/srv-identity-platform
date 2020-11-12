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
	Router.Register(courier.NewRouter(GetClients{}))
}

// 批量获取客户端
type GetClients struct {
	httpx.MethodGet
	service.ModuleClientCondition
	// 分页偏移量
	Offset int64 `name:"offset" default:"0" in:"query"`
	// 单页数据量
	Limit int64 `name:"limit" default:"10" in:"query"`
}

func (req GetClients) Path() string {
	return ""
}

type GetClientsResult struct {
	Data  []databases.ModuleClients `json:"data"`
	Total int                       `json:"total"`
}

func (req GetClients) Output(ctx context.Context) (result interface{}, err error) {
	data, count, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModuleClients(req.ModuleClientCondition, req.Offset, req.Limit)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}
	result = GetClientsResult{
		Data:  data,
		Total: count,
	}
	return result, nil
}
