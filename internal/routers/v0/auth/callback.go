package auth

import (
	"context"
	"github.com/eden-framework/client"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/common"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/dingding"
	"github.com/eden-framework/srv-identity-platform/internal/modules/users"
)

func init() {
	Router.Register(courier.NewRouter(Callback{}))
}

// 第三方回调
type Callback struct {
	httpx.MethodGet
	// Code
	Code string `name:"code" in:"query"`
	// state
	State string `name:"state" in:"query"`
}

func (req Callback) Path() string {
	return "/callback"
}

func (req Callback) Output(ctx context.Context) (result interface{}, err error) {
	// 认证流程
	providerType := (dingding.DingDing{}).ProviderID()
	provider, exist := providers.P.GetProvider(providerType)
	if !exist {
		err = errors.InternalError.StatusError().WithDesc("dingding provider not found")
	}
	userID, err := provider.GetUserID(req.Code)

	c := users.NewUserController(global.Config.SlaveDB.Get())
	user, err := c.GetUserByBindID(providerType, userID)
	if err != nil {
		if err == errors.UserBindNotFound {
			// TODO 绑定不存在，跳转新建用户界面，现在临时做法是直接创建用户
			userInfo, err := provider.GetUserInfo(userID)
			if err != nil {
				err = errors.InternalError.StatusError().WithDesc(err.Error())
				return nil, err
			}
			user, _, err = createUser(userInfo, providerType, c)
			if err != nil {
				err = errors.InternalError.StatusError().WithDesc(err.Error())
				return nil, err
			}
		} else {
			err = errors.InternalError.StatusError().WithDesc(err.Error())
			return
		}
	}

	// 生成登录凭证，并回调
	return user, nil
}

func createUser(userInfo common.UserInfo, typ enums.BindType, controller *users.UserController) (*databases.Users, *databases.UserBinds, error) {
	userID, err := client.GetUniqueID(global.ClientConfig.ID)
	if err != nil {
		return nil, nil, err
	}

	return controller.CreateUserAndBind(userID, userInfo.UserID, typ,
		users.WithMobile(userInfo.Mobile),
		users.WithEmail(userInfo.Email))
}
