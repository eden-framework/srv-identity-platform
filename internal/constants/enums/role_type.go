package enums

//go:generate eden generate enum --type-name=RoleType
// api:enum
type RoleType uint8

// 角色类型
const (
	ROLE_TYPE_UNKNOWN RoleType = iota
	ROLE_TYPE__ADMIN           // 管理员
	ROLE_TYPE__READER          // 只读
)
