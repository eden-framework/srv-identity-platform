package v0

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/authorization"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/oauth"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/providers"
)

var Router = courier.NewRouter(V0Router{})

func init() {
	Router.Register(providers.Router)
	Router.Register(authorization.Router)
	Router.Register(oauth.Router)
}

type V0Router struct {
	courier.EmptyOperator
}

func (V0Router) Path() string {
	return "/v0"
}
