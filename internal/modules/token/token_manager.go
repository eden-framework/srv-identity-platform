package token

import "github.com/eden-framework/sqlx"

type tokenManager struct {
	db sqlx.DBExecutor
}
