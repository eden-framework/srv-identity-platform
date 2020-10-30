package dingding

import (
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/clients/client_dingding"
	"testing"
)

func TestName(t *testing.T) {
	cli := &client_dingding.ClientDingDing{
		Client: client.Client{
			Host: "oapi.dingtalk.com",
			Mode: "https",
			Port: 80,
		},

		AppKey:    "dingtn1euhifut1f3pgt",
		AppSecret: "fAn6rNkZ9HNNOHcg3b4BX3VArgm7yjfGLZO_1mirZNnfsKy4qrjU26PCVaPbouN1",

		LoginAppID:     "dingoaylmyjg0jnbuvbwwh",
		LoginAppSecret: "YXntV7OZ8rTa_yfejd-lE3nlVAfxUYgAhJ7s9qcj2jN9EBh8FanYpLDmLIBb_pHO",
	}

	resp, err := cli.GetUserInfoByCode(client_dingding.GetUserInfoByCodeRequest{
		Body: client_dingding.GetUserInfoByCodeBody{
			AuthCode: "e3908cdb8fc13bbab46fdd1976f10557",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	userIDResp, err := cli.GetUserIDByUnionID(client_dingding.GetUserIDByUnionIDRequest{
		UnionID: resp.UserInfo.UnionID,
	})
	if err != nil {
		t.Fatal(err)
	}

	user, err := cli.GetUserInfoDetail(client_dingding.GetUserInfoDetailRequest{
		UserID: userIDResp.UserID,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

//https://oapi.dingtalk.com/connect/qrconnect?appid=dingoaylmyjg0jnbuvbwwh&response_type=code&scope=snsapi_login&redirect_uri=http://localhost:8802/identity-platform/v0/auth/callback
