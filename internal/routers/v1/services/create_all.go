package services

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(CreateServiceWithAll{}))
}

// 创建所有服务需要的资源
type CreateServiceWithAll struct {
	httpx.MethodPost
}

func (req CreateServiceWithAll) Path() string {
	return ""
}

func (req CreateServiceWithAll) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
