package entry

import (
	"github.com/eden-framework/courier"
	v0 "github.com/eden-framework/srv-identity-platform/internal/routers/entry/v0"
)

var Router = courier.NewRouter(Group{})

func init() {
	Router.Register(v0.Router)
}

type Group struct {
	courier.EmptyOperator
}

func (Group) Path() string {
	return "/entry"
}
