package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidUserType = errors.New("invalid UserType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("UserType", map[string]string{
		"CLIENT": "client",
		"USER":   "user",
	})
}

func ParseUserTypeFromString(s string) (UserType, error) {
	switch s {
	case "":
		return USER_TYPE_UNKNOWN, nil
	case "CLIENT":
		return USER_TYPE__CLIENT, nil
	case "USER":
		return USER_TYPE__USER, nil
	}
	return USER_TYPE_UNKNOWN, InvalidUserType
}

func ParseUserTypeFromLabelString(s string) (UserType, error) {
	switch s {
	case "":
		return USER_TYPE_UNKNOWN, nil
	case "client":
		return USER_TYPE__CLIENT, nil
	case "user":
		return USER_TYPE__USER, nil
	}
	return USER_TYPE_UNKNOWN, InvalidUserType
}

func (UserType) EnumType() string {
	return "UserType"
}

func (UserType) Enums() map[int][]string {
	return map[int][]string{
		int(USER_TYPE__CLIENT): {"CLIENT", "client"},
		int(USER_TYPE__USER):   {"USER", "user"},
	}
}

func (v UserType) String() string {
	switch v {
	case USER_TYPE_UNKNOWN:
		return ""
	case USER_TYPE__CLIENT:
		return "CLIENT"
	case USER_TYPE__USER:
		return "USER"
	}
	return "UNKNOWN"
}

func (v UserType) Label() string {
	switch v {
	case USER_TYPE_UNKNOWN:
		return ""
	case USER_TYPE__CLIENT:
		return "client"
	case USER_TYPE__USER:
		return "user"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*UserType)(nil)

func (v UserType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidUserType
	}
	return []byte(str), nil
}

func (v *UserType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseUserTypeFromString(string(bytes.ToUpper(data)))
	return
}
