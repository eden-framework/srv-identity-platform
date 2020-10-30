package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-framework/srv-identity-platform/internal/constants/enums"
)

//go:generate eden generate model UserGroup --database Config.DB --with-comments
//go:generate eden generate tag UserGroup --defaults=true
// @def primary ID
// @def unique_index U_user_group_id UserID GroupID
type UserGroup struct {
	datatypes.PrimaryID
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 用户组ID
	GroupID uint64 `json:"groupID,string" db:"f_group_id"`
	// 角色
	Role enums.RoleType `json:"role" db:"f_role"`

	datatypes.OperateTime
}
