package oauth

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/authorization"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/middleware"
)

func init() {
	authorization.Router.Register(courier.NewRouter(middleware.MustValidAccount{}, ExchangeToken{}))
}

// 使用SecureCode颁发ClientAccessToken
type ExchangeToken struct {
	httpx.MethodGet
	// Code
	Code string `name:"code" in:"query"`
}

func (req ExchangeToken) Path() string {
	return "/token"
}

func (req ExchangeToken) Output(ctx context.Context) (result interface{}, err error) {
	result, err = token.Manager.ExchangeAccessToken(ctx, enums.TOKEN_SUBJECT__CLIENT, req.Code)
	return
}
