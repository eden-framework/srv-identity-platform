package global

import "github.com/profzone/envconfig"

type DingDingConfig struct {
	AppKey         string
	AppSecret      envconfig.Password
	LoginAppID     string
	LoginAppSecret envconfig.Password
}

type EWechatConfig struct {
	CorpID     string
	CorpSecret envconfig.Password
}

var ProviderConfig = struct {
	DingDing DingDingConfig
	EWechat  EWechatConfig
}{}
