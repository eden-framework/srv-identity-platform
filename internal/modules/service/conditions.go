package service

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
)

type Condition interface {
	ToConditions(db sqlx.DBExecutor) builder.SqlCondition
}

type ServiceCondition struct {
	// 服务标识
	ServiceKey string `json:"serviceKey,omitempty" name:"serviceKey,omitempty"`
	// 服务名称
	Name string `json:"name,omitempty" name:"name,omitempty"`
}

func (c ServiceCondition) ToConditions(db sqlx.DBExecutor) builder.SqlCondition {
	m := &databases.Services{}
	t := db.T(m)
	var condition builder.SqlCondition
	if c.ServiceKey != "" {
		condition = builder.And(condition, t.F(m.FieldKeyServiceKey()).Eq(c.ServiceKey))
	}
	if c.Name != "" {
		condition = builder.And(condition, t.F(m.FieldKeyName()).Eq(c.Name))
	}
	return condition
}

type ModuleCondition struct {
	// 模块标识
	ModuleKey string `json:"moduleKey,omitempty" name:"moduleKey,omitempty"`
	// 模块名称
	Name string `json:"name,omitempty" name:"name,omitempty"`
	// 所属服务
	ServiceID uint64 `json:"serviceID,string,omitempty" name:"serviceID,string,omitempty"`
}

func (c ModuleCondition) ToConditions(db sqlx.DBExecutor) builder.SqlCondition {
	m := &databases.Modules{}
	t := db.T(m)
	var condition builder.SqlCondition
	if c.ModuleKey != "" {
		condition = builder.And(condition, t.F(m.FieldKeyModuleKey()).Eq(c.ModuleKey))
	}
	if c.Name != "" {
		condition = builder.And(condition, t.F(m.FieldKeyName()).Eq(c.Name))
	}
	if c.ServiceID != 0 {
		condition = builder.And(condition, t.F(m.FieldKeyServiceID()).Eq(c.ServiceID))
	}
	return condition
}

type ModulePermissionCondition struct {
	// 权限策略名称
	Name string `json:"name,omitempty" name:"name,omitempty"`
	// 权限标识
	PermissionKey string `json:"permissionKey,omitempty" name:"permissionKey,omitempty"`
	// 所属模块
	ModuleID uint64 `json:"moduleID,string,omitempty" name:"moduleID,string,omitempty"`
}

func (c ModulePermissionCondition) ToConditions(db sqlx.DBExecutor) builder.SqlCondition {
	m := &databases.ModulePermissions{}
	t := db.T(m)
	var condition builder.SqlCondition
	if c.Name != "" {
		condition = builder.And(condition, t.F(m.FieldKeyName()).Eq(c.Name))
	}
	if c.PermissionKey != "" {
		condition = builder.And(condition, t.F(m.FieldKeyPermissionKey()).Eq(c.PermissionKey))
	}
	if c.ModuleID != 0 {
		condition = builder.And(condition, t.F(m.FieldKeyModuleID()).Eq(c.ModuleID))
	}
	return condition
}

type ModuleClientCondition struct {
	// 所属模块
	ModuleID uint64 `json:"moduleID,string" name:"moduleID,string"`
}

func (c ModuleClientCondition) ToConditions(db sqlx.DBExecutor) builder.SqlCondition {
	m := &databases.ModuleClients{}
	t := db.T(m)
	var condition = builder.And(t.F(m.FieldKeyModuleID()).Eq(c.ModuleID))
	return condition
}
