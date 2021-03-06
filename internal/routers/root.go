package routers

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/swagger"
	"github.com/eden-framework/srv-identity-platform/internal/routers/entry"
	v0 "github.com/eden-framework/srv-identity-platform/internal/routers/v0"
)

var Router = courier.NewRouter(RootRouter{})

func init() {
	if !context.IsOnline() {
		Router.Register(swagger.SwaggerRouter)
	}
	Router.Register(v0.Router)
	Router.Register(entry.Router)
}

type RootRouter struct {
	courier.EmptyOperator
}

func (RootRouter) Path() string {
	return "/identity-platform"
}
