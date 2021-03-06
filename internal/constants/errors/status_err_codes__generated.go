package errors

import (
	"github.com/eden-framework/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("TokenExpired", 401001002, "访问令牌已过期", "", true)
	status_error.StatusErrorCodes.Register("ServiceNotFound", 404001003, "未找到服务", "", true)
	status_error.StatusErrorCodes.Register("UserBindNotFound", 404001001, "未绑定平台账户", "", true)
	status_error.StatusErrorCodes.Register("BadRequest", 400001000, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("Unauthorized", 401001000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("ModuleNotFound", 404001004, "未找到模块", "", true)
	status_error.StatusErrorCodes.Register("PermissionNotFound", 404001006, "未找到权限策略", "", true)
	status_error.StatusErrorCodes.Register("InvalidToken", 401001001, "无效的访问令牌", "", true)
	status_error.StatusErrorCodes.Register("ClientNotFound", 404001005, "未找到客户端配置", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403001000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500001000, "内部处理错误", "", true)
	status_error.StatusErrorCodes.Register("NotFound", 404001000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("Conflict", 409001000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("UserNotFound", 404001002, "用户未找到", "", true)
	status_error.StatusErrorCodes.Register("PermissionApiNotFound", 404001007, "未找到权限策略的接口配置", "", true)

}
