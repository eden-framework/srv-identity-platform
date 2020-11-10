package oauth

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/middleware"
)

func init() {
	Router.Register(courier.NewRouter(Authorize{}))
}

// 第三方登录授权，获取SecureCode
type Authorize struct {
	httpx.MethodGet
	// ClientID
	ClientID uint64 `name:"clientID,string" in:"query"`
	// RedirectURI
	RedirectURI string `name:"redirectURI" in:"query"`
	// State
	State string `name:"state" in:"query" default:""`
	// ResponseType
	ResponseType string `name:"responseType" in:"query" default:""`
}

func (req Authorize) Path() string {
	return "/authorize"
}

func (req Authorize) Output(ctx context.Context) (result interface{}, err error) {
	user := middleware.UserFromContext(ctx)
	userModel := &databases.Users{
		UserID: user.UserID,
	}
	err = userModel.FetchByUserID(global.Config.SlaveDB.Get())
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.UserNotFound
		}
		return nil, errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return token.Manager.GenerateSecureCode(ctx, userModel, req.ClientID)
}
