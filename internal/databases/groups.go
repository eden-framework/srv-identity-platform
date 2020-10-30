package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Groups --database Config.DB --with-comments
//go:generate eden generate tag Groups --defaults=true
// @def primary ID
// @def unique_index U_groups_id GroupID
type Groups struct {
	datatypes.PrimaryID
	// 业务ID
	GroupID uint64 `json:"groupID,string" db:"f_group_id"`
	// 用户组名称
	Name string `json:"name" db:"f_name"`
	// 描述
	Comment string `json:"comment" db:"f_comment,default=''"`

	datatypes.OperateTime
}
