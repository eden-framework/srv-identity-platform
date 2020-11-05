package enums

//go:generate eden generate enum --type-name=TokenSubject
// api:enum
type TokenSubject uint8

// token用途
const (
	TOKEN_SUBJECT_UNKNOWN    TokenSubject = iota
	TOKEN_SUBJECT__REFRESHER              // 刷新
	TOKEN_SUBJECT__USER                   // 用户
	TOKEN_SUBJECT__CLIENT                 // 调用方
)
