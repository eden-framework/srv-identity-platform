package providers

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers"
)

func init() {
	Router.Register(courier.NewRouter(AvailableList{}))
}

// 获取可用的第三方登录组件
type AvailableList struct {
	httpx.MethodGet
}

func (req AvailableList) Path() string {
	return ""
}

func (req AvailableList) Output(ctx context.Context) (result interface{}, err error) {
	return providers.P.GetAvailableProviderInfo(), nil
}
