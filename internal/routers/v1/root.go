package v1

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/middleware"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v1/services"
)

var Router = courier.NewRouter(V1Group{})
var AuthMiddleware = courier.NewRouter(middleware.MustValidAccount{})

func init() {
	Router.Register(AuthMiddleware)
	AuthMiddleware.Register(services.Router)
}

type V1Group struct {
	courier.EmptyOperator
}

func (V1Group) Path() string {
	return "/v1"
}
