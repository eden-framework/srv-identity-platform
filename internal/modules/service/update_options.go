package service

import (
	"github.com/eden-framework/sqlx/builder"
)

type UpdateOption interface {
	ToUpdateFieldValues(zeroFields ...string) builder.FieldValues
}

type ServiceUpdateOption struct {
	// 服务标识
	ServiceKey string `json:"serviceKey,omitempty"`
	// 服务名称
	Name string `json:"name,omitempty"`
	// 介绍
	Comment string `json:"comment,omitempty"`
}

func (s ServiceUpdateOption) ToUpdateFieldValues(zeroFields ...string) builder.FieldValues {
	return builder.FieldValuesFromStructByNonZero(s, zeroFields...)
}

type ModuleUpdateOption struct {
	// 模块标识
	ModuleKey string `json:"moduleKey,omitempty"`
	// 模块名称
	Name string `json:"name,omitempty"`
	// 描述
	Comment string `json:"comment,omitempty"`
	// 所属服务
	ServiceID uint64 `json:"serviceID,string,omitempty"`
}

func (s ModuleUpdateOption) ToUpdateFieldValues(zeroFields ...string) builder.FieldValues {
	return builder.FieldValuesFromStructByNonZero(s, zeroFields...)
}

type ModulePermissionUpdateOption struct {
	// 权限策略名称
	Name string `json:"name,omitempty"`
	// 权限标识
	PermissionKey string `json:"permissionKey,omitempty"`
	// 所属模块
	ModuleID uint64 `json:"moduleID,string,omitempty"`
}

func (s ModulePermissionUpdateOption) ToUpdateFieldValues(zeroFields ...string) builder.FieldValues {
	return builder.FieldValuesFromStructByNonZero(s, zeroFields...)
}

type ModuleClientUpdateOption struct {
	// Endpoint
	Endpoint string `json:"endpoint,omitempty"`
	// AccessKey
	AccessKey string `json:"accessKey,omitempty"`
	// AccessSecret
	AccessSecret string `json:"accessSecret,omitempty"`
	// 所属模块
	ModuleID uint64 `json:"moduleID,string,omitempty"`
}

func (s ModuleClientUpdateOption) ToUpdateFieldValues(zeroFields ...string) builder.FieldValues {
	return builder.FieldValuesFromStructByNonZero(s, zeroFields...)
}

type PermissionApiUpdateOption struct {
	// 名称
	Name string `json:"name,omitempty"`
	// 请求标识
	RequestKey string `json:"requestKey,omitempty"`
	// 请求路径
	RequestPath string `json:"requestPath,omitempty"`
	// 所属权限策略
	PermissionID uint64 `json:"permissionsID,string,omitempty"`
}

func (p PermissionApiUpdateOption) ToUpdateFieldValues(zeroFields ...string) builder.FieldValues {
	return builder.FieldValuesFromStructByNonZero(p, zeroFields...)
}
