package global

import "github.com/profzone/envconfig"

type DingDingConfig struct {
	Enabled        bool
	AppKey         string
	AppSecret      envconfig.Password
	LoginAppID     string
	LoginAppSecret envconfig.Password
}

type EWechatConfig struct {
	Enabled    bool
	CorpID     string
	CorpSecret envconfig.Password
	AgentID    string
}

var ProviderConfig = struct {
	DingDing DingDingConfig
	EWechat  EWechatConfig
}{}
