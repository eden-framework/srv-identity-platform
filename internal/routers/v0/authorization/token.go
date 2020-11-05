package authorization

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
)

func init() {
	Router.Register(courier.NewRouter(ExchangeToken{}))
}

// 使用SecureCode换取AccessToken
type ExchangeToken struct {
	httpx.MethodGet
	// Code
	Code string `name:"code" in:"query"`
}

func (req ExchangeToken) Path() string {
	return "/token"
}

func (req ExchangeToken) Output(ctx context.Context) (result interface{}, err error) {
	result, err = token.Manager.ExchangeAccessToken(ctx, req.Code)
	return
}
