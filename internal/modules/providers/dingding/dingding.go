package dingding

import (
	"fmt"
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_dingding"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/common"
)

type DingDing struct {
	client *client_dingding.ClientDingDing
}

func NewDingDingProvider(config global.DingDingConfig) *DingDing {
	p := &DingDing{
		client: &client_dingding.ClientDingDing{
			Client: client.Client{
				Host: "oapi.dingtalk.com",
				Mode: "https",
			},
			AppKey:         config.AppKey,
			AppSecret:      config.AppSecret.String(),
			LoginAppID:     config.LoginAppID,
			LoginAppSecret: config.LoginAppSecret.String(),
		},
	}
	p.client.Init()
	return p
}

func (d DingDing) ProviderID() enums.BindType {
	return enums.BIND_TYPE__DINGDING
}

func (d DingDing) ProviderName() string {
	return "钉钉"
}

func (d DingDing) ProviderIcon() string {
	return "data:image/png;base64,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"
}

func (d DingDing) ProviderEntry() string {
	return fmt.Sprintf("https://oapi.dingtalk.com/connect/qrconnect?appid=%s&response_type=code&scope=snsapi_login&state=%s_{{RANDOM_NUMBER}}&redirect_uri={{REDIRECT_URI}}", global.ProviderConfig.DingDing.LoginAppID, d.ProviderID().String())
}

func (d *DingDing) GetUserID(token string) (string, error) {
	resp, err := d.client.GetUserInfoByCode(client_dingding.GetUserInfoByCodeRequest{
		Body: client_dingding.GetUserInfoByCodeBody{
			AuthCode: token,
		},
	})
	if err != nil {
		return "", err
	}

	return resp.UserInfo.UnionID, nil
}

func (d *DingDing) GetUserInfo(userID string) (user common.UserInfo, err error) {
	userIDResp, err := d.client.GetUserIDByUnionID(client_dingding.GetUserIDByUnionIDRequest{
		UnionID: userID,
	})
	if err != nil {
		return
	}

	userInfo, err := d.client.GetUserInfoDetail(client_dingding.GetUserInfoDetailRequest{
		UserID: userIDResp.UserID,
	})
	if err != nil {
		return
	}

	return common.UserInfo{
		UserID: userInfo.UnionID,
		Name:   userInfo.Name,
		Mobile: userInfo.Mobile,
		Email:  userInfo.Email,
	}, nil
}
