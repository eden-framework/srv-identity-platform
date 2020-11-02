package dingding

import (
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_dingding"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/common"
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
