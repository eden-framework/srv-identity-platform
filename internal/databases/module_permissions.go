package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model ModulePermissions --database Config.DB --with-comments
//go:generate eden generate tag ModulePermissions --defaults=true
// @def primary ID
// @def unique_index U_permissions_id PermissionsID
// @def index I_module ModuleID
type ModulePermissions struct {
	datatypes.PrimaryID
	// 业务ID
	PermissionsID uint64 `json:"permissionsID,string" db:"f_permissions_id"`
	BasePermission
	datatypes.OperateTime
}

type BasePermission struct {
	// 权限策略名称
	Name string `json:"name" db:"f_name"`
	// 权限标识
	PermissionKey string `json:"permissionKey" db:"f_permission_key"`
	// 所属模块
	ModuleID uint64 `json:"moduleID,string" db:"f_module_id"`
}
