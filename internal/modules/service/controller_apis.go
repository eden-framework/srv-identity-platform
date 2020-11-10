package service

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

func (c *Controller) CreateModulePermissionApi(opt databases.BaseApi) (per *databases.ModulePermissionApi, err error) {
	id, err := client.GetUniqueID(c.clientID)
	if err != nil {
		return
	}
	per = &databases.ModulePermissionApi{
		ApiID:   id,
		BaseApi: opt,
	}
	err = per.Create(c.db)
	if err != nil {
		logrus.Errorf("service.Controller.CreateModulePermissionApi err: %v, opt: %+v", err, opt)
	}
	return
}

func (c *Controller) UpdateModulePermissionApi(id uint64, opt UpdateOption, zeroFieldNames ...string) error {
	m := &databases.ModulePermissionApi{
		ApiID: id,
	}
	err := m.FetchByApiID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("service.Controller.UpdateModulePermissionApi FetchByApiID err: %v, id: %d, opt: %+v", err, id, opt)
		}
		return err
	}

	fieldValues := opt.ToUpdateFieldValues(zeroFieldNames...)
	err = m.UpdateByApiIDWithMap(c.db, fieldValues)
	if err != nil {
		logrus.Errorf("service.Controller.UpdateModulePermissionApi UpdateByApiIDWithMap err: %v, id: %d, opt: %+v, fieldValues: %+v", err, id, opt, fieldValues)
	}
	return err
}

func (c *Controller) GetModulePermissionApi(condition Condition, offset, limit int64) ([]databases.ModulePermissionApi, int, error) {
	m := &databases.ModulePermissionApi{}
	sqlCondition := condition.ToConditions(c.db)
	list, err := m.List(c.db, sqlCondition, builder.Limit(limit).Offset(offset))
	if err != nil {
		logrus.Errorf("service.Controller.GetModulePermissionApi List err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	count, err := m.Count(c.db, sqlCondition)
	if err != nil {
		logrus.Errorf("service.Controller.GetModulePermissionApi Count err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	return list, count, nil
}

func (c *Controller) GetModulePermissionApiByApiID(id uint64) (model *databases.ModulePermissionApi, err error) {
	model = &databases.ModulePermissionApi{
		ApiID: id,
	}
	err = model.FetchByApiID(c.db)
	if err != nil && !sqlx.DBErr(err).IsNotFound() {
		logrus.Errorf("service.Controller.GetModulePermissionApiByApiID err: %v, id: %d", err, id)
	}
	return
}
