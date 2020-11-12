package service

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

func (c *Controller) CreateModulePermission(moduleID uint64, opt databases.BasePermission) (per *databases.ModulePermissions, err error) {
	id, err := client.GetUniqueID(c.clientID)
	if err != nil {
		return
	}
	per = &databases.ModulePermissions{
		PermissionsID:  id,
		ModuleID:       moduleID,
		BasePermission: opt,
	}
	err = per.Create(c.db)
	if err != nil {
		logrus.Errorf("service.Controller.CreateModulePermission err: %v, opt: %+v", err, opt)
	}
	return
}

func (c *Controller) UpdateModulePermission(id uint64, opt UpdateOption, zeroFieldNames ...string) error {
	m := &databases.ModulePermissions{
		PermissionsID: id,
	}
	err := m.FetchByPermissionsID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("service.Controller.UpdateModulePermission FetchByPermissionsID err: %v, id: %d, opt: %+v", err, id, opt)
		}
		return err
	}

	fieldValues := opt.ToUpdateFieldValues(zeroFieldNames...)
	err = m.UpdateByPermissionsIDWithMap(c.db, fieldValues)
	if err != nil {
		logrus.Errorf("service.Controller.UpdateModulePermission UpdateByPermissionsIDWithMap err: %v, id: %d, opt: %+v, fieldValues: %+v", err, id, opt, fieldValues)
	}
	return err
}

func (c *Controller) GetModulePermissions(condition Condition, offset, limit int64) ([]databases.ModulePermissions, int, error) {
	m := &databases.ModulePermissions{}
	sqlCondition := condition.ToConditions(c.db)
	list, err := m.List(c.db, sqlCondition, builder.Limit(limit).Offset(offset))
	if err != nil {
		logrus.Errorf("service.Controller.GetModulePermissions List err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	count, err := m.Count(c.db, sqlCondition)
	if err != nil {
		logrus.Errorf("service.Controller.GetModulePermissions Count err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	return list, count, nil
}

func (c *Controller) GetModulePermissionByPermissionID(id uint64) (model *databases.ModulePermissions, err error) {
	model = &databases.ModulePermissions{
		PermissionsID: id,
	}
	err = model.FetchByPermissionsID(c.db)
	if err != nil && !sqlx.DBErr(err).IsNotFound() {
		logrus.Errorf("service.Controller.GetModulePermissionByPermissionID err: %v, id: %d", err, id)
	}
	return
}

func (c *Controller) DeleteModulePermission(id uint64, soft bool) (err error) {
	model := &databases.ModulePermissions{
		PermissionsID: id,
	}
	if soft {
		err = model.SoftDeleteByPermissionsID(c.db)
	} else {
		err = model.DeleteByPermissionsID(c.db)
	}
	if err != nil {
		logrus.Errorf("service.Controller.DeleteModulePermission err: %v, id: %d, soft: %v", err, id, soft)
	}
	return
}
