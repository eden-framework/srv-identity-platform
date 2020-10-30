package client_dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/timex"
	"time"
)

type ClientDingDing struct {
	client.Client

	AppKey    string
	AppSecret string

	LoginAppID     string
	LoginAppSecret string

	accessTokenExpireDuration  time.Duration
	accessToken                string
	lastAccessTokenRequireTime time.Time
}

func (c *ClientDingDing) SetDefaults() {
	c.Client.MarshalDefaults(&c.Client)
}

func (c *ClientDingDing) Init() {
	c.SetDefaults()
}

// 通过扫码code获取用户基本信息
func (c *ClientDingDing) GetUserInfoByCode(req GetUserInfoByCodeRequest, metas ...courier.Metadata) (resp GetUserInfoByCodeResponse, err error) {
	req.AuthRequest = c.getAuthRequest()
	request := c.Request("", "POST", "/sns/getuserinfo_bycode", req, metas...)
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

// 根据UnionID获取UserID
func (c *ClientDingDing) GetUserIDByUnionID(req GetUserIDByUnionIDRequest, metas ...courier.Metadata) (resp GetUserIDByUnionIDResponse, err error) {
	req.AccessToken, err = c.getAccessToken(metas...)
	if err != nil {
		return
	}

	request := c.Request("", "GET", "/user/getUseridByUnionid", req, metas...)
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
func (c *ClientDingDing) GetUserInfoDetail(req GetUserInfoDetailRequest, metas ...courier.Metadata) (resp GetUserInfoDetailResponse, err error) {
	req.AccessToken, err = c.getAccessToken(metas...)
	if err != nil {
		return
	}

	request := c.Request("", "GET", "/user/get", req, metas...)
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
func (c *ClientDingDing) getAccessToken(metas ...courier.Metadata) (string, error) {
	if c.accessToken != "" && time.Since(c.lastAccessTokenRequireTime).Seconds() < c.accessTokenExpireDuration.Seconds() {
		return c.accessToken, nil
	}
	req := GetAccessTokenRequest{
		AppKey:    c.AppKey,
		AppSecret: c.AppSecret,
	}
	request := c.Request("", "GET", "/gettoken", req, metas...)
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

func (c *ClientDingDing) signature(timestamp uint64) string {
	mac := hmac.New(sha256.New, []byte(c.LoginAppSecret))
	mac.Write([]byte(fmt.Sprintf("%d", timestamp)))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes)
}

func (c *ClientDingDing) getAuthRequest() AuthRequest {
	timestamp := timex.UnixTimestamp()
	signature := c.signature(timestamp)
	return AuthRequest{
		AccessKey: c.LoginAppID,
		Timestamp: timestamp,
		Signature: signature,
	}
}
