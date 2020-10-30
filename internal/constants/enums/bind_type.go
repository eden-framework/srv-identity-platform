package enums

//go:generate eden generate enum --type-name=BindType
// api:enum
type BindType uint8

// 第三方账户绑定类型
const (
	BIND_TYPE_UNKNOWN   BindType = iota
	BIND_TYPE__LDAP              // LDAP
	BIND_TYPE__WECHAT            // 微信
	BIND_TYPE__EWECHAT           // 企业微信
	BIND_TYPE__DINGDING          // 钉钉
	BIND_TYPE__FEISHU            // 飞书
)
