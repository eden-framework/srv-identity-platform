package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
)

//go:generate eden generate model UserBinds --database Config.DB --with-comments
//go:generate eden generate tag UserBinds --defaults=true
// @def primary ID
// @def unique_index U_user_binds_id UserID BindID
// @def index I_user UserID
// @def unique_index U_bind Type BindID
type UserBinds struct {
	datatypes.PrimaryID
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 第三方绑定系统唯一ID
	BindID string `json:"bindID" db:"f_bind_id"`
	// 第三方来源类型
	Type enums.BindType `json:"type" db:"f_type"`
	// 过期时间
	ExpireAt datatypes.MySQLTimestamp `db:"f_expire_at,default='0'" json:"expireAt" `

	datatypes.OperateTime
}
