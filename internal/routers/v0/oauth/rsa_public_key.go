package oauth

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
)

func init() {
	Router.Register(courier.NewRouter(RSAPublicKey{}))
}

// 获取token公钥
type RSAPublicKey struct {
	httpx.MethodGet
}

func (req RSAPublicKey) Path() string {
	return "/public-key"
}

func (req RSAPublicKey) Output(ctx context.Context) (result interface{}, err error) {
	return token.Manager.PublicKey(), nil
}
