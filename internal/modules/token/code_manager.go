package token

import (
	"context"
	"fmt"
	"github.com/eden-framework/client"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/plugin-cache/cache"
	"github.com/profzone/envconfig"
	"strconv"
	"time"
)

type codeManager struct {
	cache    *cache.Cache
	clientID client_srv_id.ClientSrvIDInterface

	defaultExpireIn time.Duration
}

func (m *codeManager) GenerateSecureCode(ctx context.Context, state string) (code SecureCode, err error) {
	c, err := m.generateCode(ctx, "secureCode", State(state), m.defaultExpireIn)
	if err != nil {
		return
	}

	return SecureCode{
		Code:     c,
		ExpireIn: envconfig.Duration(m.defaultExpireIn),
	}, nil
}

func (m *codeManager) ExchangeSecureCode(ctx context.Context, code string) (string, error) {
	state, err := m.getStateFromCode(ctx, "secureCode", code, false)
	if err != nil {
		return "", err
	}
	return string(state), nil
}

func (m *codeManager) generateCode(ctx context.Context, subject string, state State, expireIn time.Duration) (code string, err error) {
	id, err := client.GetUniqueID(m.clientID)
	if err != nil {
		return
	}

	code = strconv.FormatUint(id, 16)
	err = m.cache.Set(ctx, m.keyFor(subject, code), state, expireIn)
	return
}

func (m *codeManager) getStateFromCode(ctx context.Context, subject string, code string, withDel bool) (state State, err error) {
	key := m.keyFor(subject, code)
	err = m.cache.Get(ctx, key, &state)
	if err != nil {
		return
	}

	if withDel {
		_ = m.cache.Del(ctx, key)
	}

	return
}

func (m *codeManager) keyFor(subject, code string) string {
	return fmt.Sprintf("%s:code:%s", subject, code)
}
