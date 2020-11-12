package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Modules --database Config.DB --with-comments
//go:generate eden generate tag Modules --defaults=true
// @def primary ID
// @def unique_index U_module_id ModuleID
// @def index I_service ServiceID
type Modules struct {
	datatypes.PrimaryID
	// 业务ID
	ModuleID uint64 `json:"moduleID,string" db:"f_module_id"`
	BaseModule
	datatypes.OperateTime
}

type BaseModule struct {
	// 模块标识
	ModuleKey string `json:"moduleKey" db:"f_module_key"`
	// 模块名称
	Name string `json:"name" db:"f_name"`
	// 描述
	Comment string `json:"comment,omitempty" db:"f_comment,default=''"`
	// 所属服务
	ServiceID uint64 `json:"serviceID,string" db:"f_service_id"`
}
