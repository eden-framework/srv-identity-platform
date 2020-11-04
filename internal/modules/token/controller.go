package token

import (
	"context"
	"fmt"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/plugin-cache/cache"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"time"
)

var Manager *controller

func Initializer() error {
	Manager = newController(global.CacheConfig.Cache,
		global.Config.MasterDB,
		global.ClientConfig.ID,
		time.Duration(global.Config.SecureCodeDefaultExpire))
	return nil
}

type controller struct {
	tokenManger *tokenManager
	codeManager *codeManager
}

func newController(cache *cache.Cache, db sqlx.DBExecutor, clientID client_srv_id.ClientSrvIDInterface, codeExpireIn time.Duration) *controller {
	return &controller{
		codeManager: &codeManager{cache: cache, clientID: clientID, defaultExpireIn: codeExpireIn},
		tokenManger: &tokenManager{db: db},
	}
}

func (c *controller) GenerateSecureCode(ctx context.Context, user *databases.Users, clientID uint64) (SecureCode, error) {
	state := fmt.Sprintf("%d|||%d", user.ID, clientID)
	return c.codeManager.GenerateSecureCode(ctx, state)
}

func (c *controller) ExchangeAccessToken(code string) (string, error) {
	return "", nil
}
