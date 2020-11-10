package oauth

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/users"
	"github.com/eden-framework/srv-identity-platform/internal/routers/v0/middleware"
)

func init() {
	Router.Register(courier.NewRouter(GetUserByToken{}))
}

// 通过Token获取用户信息
type GetUserByToken struct {
	httpx.MethodGet
}

func (req GetUserByToken) Path() string {
	return "/user"
}

func (req GetUserByToken) Output(ctx context.Context) (result interface{}, err error) {
	userType := middleware.UserFromContext(ctx)
	controller := users.NewController(global.Config.SlaveDB.Get())
	user, err := controller.GetUserByUserID(userType.UserID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.UserNotFound.StatusError().WithMsg("根据用户ID没有找到用户")
			return
		}
		return nil, errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return user, nil
}
