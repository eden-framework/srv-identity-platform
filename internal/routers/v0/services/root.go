package services

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/services/modules"
)

var Router = courier.NewRouter(Group{})

func init() {
	Router.Register(modules.Router)
}

type Group struct {
	courier.EmptyOperator
}

func (Group) Path() string {
	return "/services"
}
