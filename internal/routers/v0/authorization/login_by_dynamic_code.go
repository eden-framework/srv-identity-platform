package authorization

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(LoginByDynamicCode{}))
}

// 通过动态验证码登录，支持邮箱或手机，登录后颁发UserAccessToken
type LoginByDynamicCode struct {
	httpx.MethodPost
}

func (req LoginByDynamicCode) Path() string {
	return "/by-dynamic-code"
}

func (req LoginByDynamicCode) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
