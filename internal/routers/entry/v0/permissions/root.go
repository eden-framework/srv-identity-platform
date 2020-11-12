package permissions

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(Group{})

type Group struct {
	courier.EmptyOperator
}

func (Group) Path() string {
	return "/permissions"
}
