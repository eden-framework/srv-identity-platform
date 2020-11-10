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
	Router.Register(courier.NewRouter(CreateService{}))
}

// 创建服务
type CreateService struct {
	httpx.MethodPost
	Body databases.BaseService `in:"body"`
}

func (req CreateService) Path() string {
	return ""
}

func (req CreateService) Output(ctx context.Context) (result interface{}, err error) {
	_, err = service.NewController(global.Config.MasterDB, global.ClientConfig.ID).CreateService(req.Body)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return nil, err
}
