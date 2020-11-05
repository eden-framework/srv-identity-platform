package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidTokenSubject = errors.New("invalid TokenSubject")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("TokenSubject", map[string]string{
		"CLIENT":    "调用方",
		"USER":      "用户",
		"REFRESHER": "刷新",
	})
}

func ParseTokenSubjectFromString(s string) (TokenSubject, error) {
	switch s {
	case "":
		return TOKEN_SUBJECT_UNKNOWN, nil
	case "CLIENT":
		return TOKEN_SUBJECT__CLIENT, nil
	case "USER":
		return TOKEN_SUBJECT__USER, nil
	case "REFRESHER":
		return TOKEN_SUBJECT__REFRESHER, nil
	}
	return TOKEN_SUBJECT_UNKNOWN, InvalidTokenSubject
}

func ParseTokenSubjectFromLabelString(s string) (TokenSubject, error) {
	switch s {
	case "":
		return TOKEN_SUBJECT_UNKNOWN, nil
	case "调用方":
		return TOKEN_SUBJECT__CLIENT, nil
	case "用户":
		return TOKEN_SUBJECT__USER, nil
	case "刷新":
		return TOKEN_SUBJECT__REFRESHER, nil
	}
	return TOKEN_SUBJECT_UNKNOWN, InvalidTokenSubject
}

func (TokenSubject) EnumType() string {
	return "TokenSubject"
}

func (TokenSubject) Enums() map[int][]string {
	return map[int][]string{
		int(TOKEN_SUBJECT__CLIENT):    {"CLIENT", "调用方"},
		int(TOKEN_SUBJECT__USER):      {"USER", "用户"},
		int(TOKEN_SUBJECT__REFRESHER): {"REFRESHER", "刷新"},
	}
}

func (v TokenSubject) String() string {
	switch v {
	case TOKEN_SUBJECT_UNKNOWN:
		return ""
	case TOKEN_SUBJECT__CLIENT:
		return "CLIENT"
	case TOKEN_SUBJECT__USER:
		return "USER"
	case TOKEN_SUBJECT__REFRESHER:
		return "REFRESHER"
	}
	return "UNKNOWN"
}

func (v TokenSubject) Label() string {
	switch v {
	case TOKEN_SUBJECT_UNKNOWN:
		return ""
	case TOKEN_SUBJECT__CLIENT:
		return "调用方"
	case TOKEN_SUBJECT__USER:
		return "用户"
	case TOKEN_SUBJECT__REFRESHER:
		return "刷新"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*TokenSubject)(nil)

func (v TokenSubject) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidTokenSubject
	}
	return []byte(str), nil
}

func (v *TokenSubject) UnmarshalText(data []byte) (err error) {
	*v, err = ParseTokenSubjectFromString(string(bytes.ToUpper(data)))
	return
}
