package v0

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/auth"
)

var Router = courier.NewRouter(V0Router{})

func init() {
	Router.Register(auth.Router)
}

type V0Router struct {
	courier.EmptyOperator
}

func (V0Router) Path() string {
	return "/v0"
}
