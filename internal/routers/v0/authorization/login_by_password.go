package authorization

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(LoginByPassword{}))
}

// 通过帐密登录，并颁发UserAccessToken
type LoginByPassword struct {
	httpx.MethodPost
}

func (req LoginByPassword) Path() string {
	return "/by-password"
}

func (req LoginByPassword) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
