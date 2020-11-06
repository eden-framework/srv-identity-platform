package middleware

import (
	"context"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
	"github.com/eden-framework/srv-identity-platform/internal/modules/users"
	"strconv"
)

// 校验accessToken，并获取账户
type MustValidAccount struct {
	Authorization string `name:"Authorization" in:"header"`
}

func (req MustValidAccount) ContextKey() string {
	return "auth.Account"
}

func (req MustValidAccount) Output(ctx context.Context) (result interface{}, err error) {
	auth := token.ParseAuthorizationToken(req.Authorization)
	accessToken := auth.Get("Bearer")

	tok, err := token.Manager.ValidateToken(accessToken)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseUint(tok.Audience, 10, 64)
	if err != nil {
		return nil, errors.InvalidToken
	}

	resp := &users.UserIDWithUserType{
		UserID: userID,
	}

	switch tok.Subject {
	case enums.TOKEN_SUBJECT__USER:
		resp.UserType = enums.USER_TYPE__USER
	case enums.TOKEN_SUBJECT__CLIENT:
		resp.UserType = enums.USER_TYPE__CLIENT
	}

	return resp, nil
}

func UserFromContext(ctx context.Context) *users.UserIDWithUserType {
	a, ok := ctx.Value((&MustValidAccount{}).ContextKey()).(*users.UserIDWithUserType)
	if ok {
		return a
	}
	return nil
}
