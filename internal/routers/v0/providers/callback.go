package providers

import (
	"context"
	"github.com/eden-framework/client"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/common"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
	"github.com/eden-framework/srv-identity-platform/internal/modules/users"
	"net/url"
	"strings"
)

func init() {
	Router.Register(courier.NewRouter(Callback{}))
}

// 第三方回调验证
type Callback struct {
	httpx.MethodGet
	// Code for provider verify
	Code string `name:"code" in:"query" default:""`
	// State for provider verify
	State string `name:"state" in:"query"`

	// ClientID for inner verify
	ClientID uint64 `name:"clientID,string" in:"query"`
	// RedirectURI for inner verify
	RedirectURI string `name:"redirectURI" in:"query"`
}

func (req Callback) Path() string {
	return "/callback"
}

func (req Callback) Output(ctx context.Context) (result interface{}, err error) {
	if req.Code == "" {
		// 用户取消授权，跳转至登录前的页面，现在临时报错
		err = errors.Unauthorized
		return
	}

	// 认证流程
	// 根据State标识获取provider
	states := strings.Split(req.State, "_")
	if len(states) < 2 {
		err = errors.BadRequest.StatusError().WithDesc("state has wrong format")
		return
	}
	providerType, err := enums.ParseBindTypeFromString(states[0])
	if err != nil {
		err = errors.BadRequest.StatusError().WithDesc(err.Error())
		return
	}
	provider, exist := providers.P.GetProvider(providerType)
	if !exist {
		err = errors.InternalError.StatusError().WithDesc("dingding provider not found")
	}
	userID, err := provider.GetUserID(req.Code)

	c := users.NewController(global.Config.SlaveDB.Get())
	user, err := c.GetUserByBindID(providerType, userID)
	if err != nil {
		if err == errors.UserBindNotFound {
			// 绑定不存在，直接创建并绑定用户
			userInfo, err := provider.GetUserInfo(userID)
			if err != nil {
				err = errors.InternalError.StatusError().WithDesc(err.Error())
				return nil, err
			}
			user, _, err = createUserOrBind(userInfo, providerType, c)
			if err != nil {
				err = errors.InternalError.StatusError().WithDesc(err.Error())
				return nil, err
			}
		} else {
			err = errors.InternalError.StatusError().WithDesc(err.Error())
			return
		}
	}

	// 生成登录凭证
	code, err := token.Manager.GenerateSecureCode(ctx, user, req.ClientID)
	if err != nil {
		return
	}

	uri, err := url.ParseRequestURI(req.RedirectURI)
	if err != nil {
		return
	}

	query := uri.Query()
	query.Add("code", code.Code)
	uri.RawQuery = query.Encode()
	return httpx.RedirectWithStatusFound(uri.String()), nil
}

func createUserOrBind(userInfo common.UserInfo, typ enums.BindType, controller *users.Controller) (*databases.Users, *databases.UserBinds, error) {
	// 通过手机号查询用户是否存在，若已存在则直接绑定
	if userInfo.Mobile != "" {
		user, err := controller.GetUserByMobile(userInfo.Mobile)
		if err == nil {
			bind, err := controller.CreateBind(user.UserID, userInfo.UserID, typ)
			if err != nil {
				return nil, nil, err
			}
			return user, bind, err
		} else if err != errors.UserNotFound {
			return nil, nil, err
		}
	}

	// 不存在则创建用户并建立绑定
	userID, err := client.GetUniqueID(global.ClientConfig.ID)
	if err != nil {
		return nil, nil, err
	}

	return controller.CreateUserAndBind(userID, userInfo.UserID, typ,
		users.WithMobile(userInfo.Mobile),
		users.WithEmail(userInfo.Email))
}
