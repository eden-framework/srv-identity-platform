package client_ewechat

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"time"
)

type ClientEWechat struct {
	client.Client

	CorpID     string
	CortSecret string

	accessTokenExpireDuration  time.Duration
	accessToken                string
	lastAccessTokenRequireTime time.Time
}

func (c *ClientEWechat) SetDefaults() {
	c.Client.MarshalDefaults(&c.Client)
}

func (c *ClientEWechat) Init() {
	c.SetDefaults()
}

// 通过扫码code获取用户基本信息
func (c *ClientEWechat) GetUserInfoByCode(req GetUserInfoByCodeRequest, metas ...courier.Metadata) (resp GetUserInfoByCodeResponse, err error) {
	req.AccessToken, err = c.getAccessToken(metas...)
	if err != nil {
		return
	}

	request := c.Request("", "POST", "/cgi-bin/user/getuserinfo", req, metas...)
	err = request.Do().Into(&resp)
	if err != nil {
		return
	}

	if resp.ErrCode > 0 {
		err = errors.BadRequest.StatusError().WithMsg(resp.ErrMsg)
		return
	}
	return
}

// 获取用户详情
func (c *ClientEWechat) GetUserInfoDetail(req GetUserInfoDetailRequest, metas ...courier.Metadata) (resp GetUserInfoDetailResponse, err error) {
	req.AccessToken, err = c.getAccessToken(metas...)
	if err != nil {
		return
	}

	request := c.Request("", "GET", "/cgi-bin/user/get", req, metas...)
	err = request.Do().Into(&resp)
	if err != nil {
		return
	}

	if resp.ErrCode > 0 {
		err = errors.BadRequest.StatusError().WithMsg(resp.ErrMsg)
		return
	}
	return
}

// 获取AccessToken
func (c *ClientEWechat) getAccessToken(metas ...courier.Metadata) (string, error) {
	if c.accessToken != "" && time.Since(c.lastAccessTokenRequireTime).Seconds() < c.accessTokenExpireDuration.Seconds() {
		return c.accessToken, nil
	}
	req := GetAccessTokenRequest{
		CorpID:     c.CorpID,
		CorpSecret: c.CortSecret,
	}
	request := c.Request("", "GET", "/cgi-bin/gettoken", req, metas...)
	resp := GetAccessTokenResponse{}
	err := request.Do().Into(&resp)
	if err != nil {
		return "", err
	}

	if resp.ErrCode > 0 {
		err = errors.BadRequest.StatusError().WithMsg(resp.ErrMsg)
		return "", err
	}

	c.accessToken = resp.AccessToken
	c.lastAccessTokenRequireTime = time.Now()

	if resp.ExpireIn > 0 {
		c.accessTokenExpireDuration = time.Duration(resp.ExpireIn) * time.Second
	} else {
		c.accessTokenExpireDuration = 7200 * time.Second
	}

	return resp.AccessToken, nil
}
