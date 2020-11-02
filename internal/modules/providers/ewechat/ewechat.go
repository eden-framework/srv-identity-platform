package ewechat

import (
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_ewechat"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers/common"
)

type EWechat struct {
	client *client_ewechat.ClientEWechat
}

func NewEWechatProvider(config global.EWechatConfig) *EWechat {
	p := &EWechat{
		client: &client_ewechat.ClientEWechat{
			Client: client.Client{
				Host: "qyapi.weixin.qq.com",
				Mode: "https",
			},

			CorpID:     config.CorpID,
			CortSecret: config.CorpSecret.String(),
		},
	}
	p.client.Init()
	return p
}

func (d EWechat) ProviderID() enums.BindType {
	return enums.BIND_TYPE__EWECHAT
}

func (d *EWechat) GetUserID(token string) (string, error) {
	resp, err := d.client.GetUserInfoByCode(client_ewechat.GetUserInfoByCodeRequest{
		Code: token,
	})
	if err != nil {
		return "", err
	}

	return resp.UserID, nil
}

func (d *EWechat) GetUserInfo(userID string) (user common.UserInfo, err error) {
	userInfo, err := d.client.GetUserInfoDetail(client_ewechat.GetUserInfoDetailRequest{
		UserID: userID,
	})
	if err != nil {
		return
	}

	return common.UserInfo{
		UserID: userInfo.UserID,
		Name:   userInfo.Name,
		Mobile: userInfo.Mobile,
		Email:  userInfo.Email,
	}, nil
}
