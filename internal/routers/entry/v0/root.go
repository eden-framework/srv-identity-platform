package v0

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/srv-identity-platform/internal/routers/entry/v0/permissions"
)

var Router = courier.NewRouter(Group{})

func init() {
	Router.Register(permissions.Router)
}

type Group struct {
	courier.EmptyOperator
}

func (Group) Path() string {
	return "/v0"
}
