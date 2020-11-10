package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model ModulePermissionApi --database Config.DB --with-comments
//go:generate eden generate tag ModulePermissionApi --defaults=true
// @def primary ID
// @def unique_index U_api_id ApiID
// @def index I_permission PermissionID
type ModulePermissionApi struct {
	datatypes.PrimaryID
	// 业务ID
	ApiID uint64 `json:"apiID,string" db:"f_module_permission_api_id"`
	BaseApi
	datatypes.OperateTime
}

type BaseApi struct {
	// 名称
	Name string `json:"name" db:"f_name"`
	// 请求标识
	RequestKey string `json:"requestKey" db:"f_request_key,default=''"`
	// 请求路径
	RequestPath string `json:"requestPath" db:"f_request_path,default=''"`
	// 所属权限策略
	PermissionID uint64 `json:"permissionsID,string" db:"f_permissions_id"`
}
