package enums

//go:generate eden generate enum --type-name=UserType
// api:enum
type UserType uint8

// 账号类型
const (
	USER_TYPE_UNKNOWN UserType = iota
	USER_TYPE__USER            // user
	USER_TYPE__CLIENT          // client
)
