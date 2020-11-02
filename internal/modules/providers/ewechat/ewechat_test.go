package ewechat

import (
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_ewechat"
	"testing"
)

func TestName(t *testing.T) {

	cli := &client_ewechat.ClientEWechat{
		Client: client.Client{
			Host: "qyapi.weixin.qq.com",
			Mode: "https",
		},

		CorpID:     "ww9304d15fc3ac4a74",
		CortSecret: "6gnrFikMJ4-iBAfy9IDAeZnXfmE8Se0ykhnQW7d_k_g",
	}

	resp, err := cli.GetUserInfoByCode(client_ewechat.GetUserInfoByCodeRequest{
		Code: "47vUE9wpRX64fHldRtL1HJ_a84j566wpqsai0tnHcgk",
	})
	if err != nil {
		t.Fatal(err)
	}

	user, err := cli.GetUserInfoDetail(client_ewechat.GetUserInfoDetailRequest{
		UserID: resp.UserID,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

//https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=ww9304d15fc3ac4a74&agentid=1000002&state=EWECHAT_1234567890&redirect_uri=http://www.profzone.net:8802/identity-platform/v0/auth/callback
