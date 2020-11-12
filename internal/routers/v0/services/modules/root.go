package modules

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/services/modules/clients"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/services/modules/permissions"
)

var Router = courier.NewRouter(Group{})

func init() {
	Router.Register(clients.Router)
	Router.Register(permissions.Router)
}

type Group struct {
	courier.EmptyOperator
}

func (Group) Path() string {
	return "/:serviceID/modules"
}
