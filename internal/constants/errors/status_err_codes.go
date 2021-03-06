package errors

import (
	"net/http"

	"github.com/eden-framework/courier/status_error"
)

//go:generate eden generate error
const ServiceStatusErrorCode = 1 * 1e3 // todo rename this

const (
	// 请求参数错误
	BadRequest status_error.StatusErrorCode = http.StatusBadRequest*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 未找到
	NotFound status_error.StatusErrorCode = http.StatusNotFound*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 未绑定平台账户
	UserBindNotFound
	// @errTalk 用户未找到
	UserNotFound
	// @errTalk 未找到服务
	ServiceNotFound
	// @errTalk 未找到模块
	ModuleNotFound
	// @errTalk 未找到客户端配置
	ClientNotFound
	// @errTalk 未找到权限策略
	PermissionNotFound
	// @errTalk 未找到权限策略的接口配置
	PermissionApiNotFound
)

const (
	// @errTalk 未授权
	Unauthorized status_error.StatusErrorCode = http.StatusUnauthorized*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 无效的访问令牌
	InvalidToken
	// @errTalk 访问令牌已过期
	TokenExpired
)

const (
	// @errTalk 操作冲突
	Conflict status_error.StatusErrorCode = http.StatusConflict*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 不允许操作
	Forbidden status_error.StatusErrorCode = http.StatusForbidden*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 内部处理错误
	InternalError status_error.StatusErrorCode = http.StatusInternalServerError*1e6 + ServiceStatusErrorCode + iota
)
