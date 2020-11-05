package token

import (
	"context"
	"fmt"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/plugin-cache/cache"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"strconv"
	"strings"
	"time"
)

var Manager *controller

func Initializer() error {
	Manager = newController(global.Config.JwtIssuer,
		global.CacheConfig.Cache,
		global.Config.MasterDB,
		global.ClientConfig.ID,
		time.Duration(global.Config.SecureCodeDefaultExpire),
		time.Duration(global.Config.AccessTokenDefaultExpire),
		global.Config.JwtPrivateKey.String())
	return nil
}

type controller struct {
	tokenManger *tokenManager
	codeManager *codeManager
}

func newController(issuer string, cache *cache.Cache, db sqlx.DBExecutor, clientID client_srv_id.ClientSrvIDInterface, codeExpireIn time.Duration, tokenExpireIn time.Duration, privateKey string) *controller {
	jwtPK := newRSA(privateKey)
	return &controller{
		codeManager: &codeManager{cache: cache, clientID: clientID, defaultExpireIn: codeExpireIn},
		tokenManger: newTokenManager(issuer, db, clientID, jwtPK, tokenExpireIn),
	}
}

func (c *controller) GenerateSecureCode(ctx context.Context, user *databases.Users, clientID uint64) (SecureCode, error) {
	state := fmt.Sprintf("%d|||%d", user.UserID, clientID)
	code, err := c.codeManager.GenerateSecureCode(ctx, state)
	if err != nil {
		return code, errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return code, nil
}

func (c *controller) ExchangeAccessToken(ctx context.Context, subject enums.TokenSubject, code string) (token AccessToken, err error) {
	state, err := c.codeManager.ExchangeSecureCode(ctx, code)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}

	states := strings.Split(state, "|||")
	if len(states) != 2 {
		err = errors.Forbidden.StatusError().WithMsg("Code 所包含的信息非法")
		return
	}

	_, err = strconv.ParseUint(states[0], 10, 64)
	if err != nil {
		err = errors.Forbidden.StatusError().WithMsg("Code 所包含的用户标识非法")
		return
	}

	token, err = c.tokenManger.ExchangeAccessToken(subject, states[0])
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}

	return token, nil
}
