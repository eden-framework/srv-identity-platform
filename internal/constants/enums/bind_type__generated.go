package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidBindType = errors.New("invalid BindType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("BindType", map[string]string{
		"FEISHU":   "飞书",
		"DINGDING": "钉钉",
		"EWECHAT":  "企业微信",
		"WECHAT":   "微信",
		"LDAP":     "LDAP",
	})
}

func ParseBindTypeFromString(s string) (BindType, error) {
	switch s {
	case "":
		return BIND_TYPE_UNKNOWN, nil
	case "FEISHU":
		return BIND_TYPE__FEISHU, nil
	case "DINGDING":
		return BIND_TYPE__DINGDING, nil
	case "EWECHAT":
		return BIND_TYPE__EWECHAT, nil
	case "WECHAT":
		return BIND_TYPE__WECHAT, nil
	case "LDAP":
		return BIND_TYPE__LDAP, nil
	}
	return BIND_TYPE_UNKNOWN, InvalidBindType
}

func ParseBindTypeFromLabelString(s string) (BindType, error) {
	switch s {
	case "":
		return BIND_TYPE_UNKNOWN, nil
	case "飞书":
		return BIND_TYPE__FEISHU, nil
	case "钉钉":
		return BIND_TYPE__DINGDING, nil
	case "企业微信":
		return BIND_TYPE__EWECHAT, nil
	case "微信":
		return BIND_TYPE__WECHAT, nil
	case "LDAP":
		return BIND_TYPE__LDAP, nil
	}
	return BIND_TYPE_UNKNOWN, InvalidBindType
}

func (BindType) EnumType() string {
	return "BindType"
}

func (BindType) Enums() map[int][]string {
	return map[int][]string{
		int(BIND_TYPE__FEISHU):   {"FEISHU", "飞书"},
		int(BIND_TYPE__DINGDING): {"DINGDING", "钉钉"},
		int(BIND_TYPE__EWECHAT):  {"EWECHAT", "企业微信"},
		int(BIND_TYPE__WECHAT):   {"WECHAT", "微信"},
		int(BIND_TYPE__LDAP):     {"LDAP", "LDAP"},
	}
}

func (v BindType) String() string {
	switch v {
	case BIND_TYPE_UNKNOWN:
		return ""
	case BIND_TYPE__FEISHU:
		return "FEISHU"
	case BIND_TYPE__DINGDING:
		return "DINGDING"
	case BIND_TYPE__EWECHAT:
		return "EWECHAT"
	case BIND_TYPE__WECHAT:
		return "WECHAT"
	case BIND_TYPE__LDAP:
		return "LDAP"
	}
	return "UNKNOWN"
}

func (v BindType) Label() string {
	switch v {
	case BIND_TYPE_UNKNOWN:
		return ""
	case BIND_TYPE__FEISHU:
		return "飞书"
	case BIND_TYPE__DINGDING:
		return "钉钉"
	case BIND_TYPE__EWECHAT:
		return "企业微信"
	case BIND_TYPE__WECHAT:
		return "微信"
	case BIND_TYPE__LDAP:
		return "LDAP"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BindType)(nil)

func (v BindType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBindType
	}
	return []byte(str), nil
}

func (v *BindType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBindTypeFromString(string(bytes.ToUpper(data)))
	return
}
