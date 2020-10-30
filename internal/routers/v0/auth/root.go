package auth

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(AuthGroup{})

type AuthGroup struct {
	courier.EmptyOperator
}

func (AuthGroup) Path() string {
	return "/auth"
}
