package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Users --database Config.DB --with-comments
//go:generate eden generate tag Users --defaults=true
// @def primary ID
// @def unique_index U_user_id UserID
type Users struct {
	datatypes.PrimaryID
	// 业务ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 用户名
	UserName string `json:"userName" db:"f_user_name,default=''"`
	// 密码
	Password string `json:"-" db:"f_password,default=''"`
	// 盐值
	Salt string `json:"-" db:"f_salt,default=''"`
	// 手机号
	Mobile string `json:"mobile" db:"f_mobile,default=''"`
	// 邮箱
	Email string `json:"email" db:"f_email,default=''"`

	datatypes.OperateTime
}
