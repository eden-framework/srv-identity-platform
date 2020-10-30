package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidRoleType = errors.New("invalid RoleType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("RoleType", map[string]string{
		"READER": "只读",
		"ADMIN":  "管理员",
	})
}

func ParseRoleTypeFromString(s string) (RoleType, error) {
	switch s {
	case "":
		return ROLE_TYPE_UNKNOWN, nil
	case "READER":
		return ROLE_TYPE__READER, nil
	case "ADMIN":
		return ROLE_TYPE__ADMIN, nil
	}
	return ROLE_TYPE_UNKNOWN, InvalidRoleType
}

func ParseRoleTypeFromLabelString(s string) (RoleType, error) {
	switch s {
	case "":
		return ROLE_TYPE_UNKNOWN, nil
	case "只读":
		return ROLE_TYPE__READER, nil
	case "管理员":
		return ROLE_TYPE__ADMIN, nil
	}
	return ROLE_TYPE_UNKNOWN, InvalidRoleType
}

func (RoleType) EnumType() string {
	return "RoleType"
}

func (RoleType) Enums() map[int][]string {
	return map[int][]string{
		int(ROLE_TYPE__READER): {"READER", "只读"},
		int(ROLE_TYPE__ADMIN):  {"ADMIN", "管理员"},
	}
}

func (v RoleType) String() string {
	switch v {
	case ROLE_TYPE_UNKNOWN:
		return ""
	case ROLE_TYPE__READER:
		return "READER"
	case ROLE_TYPE__ADMIN:
		return "ADMIN"
	}
	return "UNKNOWN"
}

func (v RoleType) Label() string {
	switch v {
	case ROLE_TYPE_UNKNOWN:
		return ""
	case ROLE_TYPE__READER:
		return "只读"
	case ROLE_TYPE__ADMIN:
		return "管理员"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RoleType)(nil)

func (v RoleType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRoleType
	}
	return []byte(str), nil
}

func (v *RoleType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRoleTypeFromString(string(bytes.ToUpper(data)))
	return
}
