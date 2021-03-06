package service

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

func (c *Controller) CreateModule(serviceID uint64, opt databases.BaseModule) (module *databases.Modules, err error) {
	id, err := client.GetUniqueID(c.clientID)
	if err != nil {
		return
	}
	module = &databases.Modules{
		ModuleID:   id,
		ServiceID:  serviceID,
		BaseModule: opt,
	}
	err = module.Create(c.db)
	if err != nil {
		logrus.Errorf("service.Controller.CreateModule err: %v, opt: %+v", err, opt)
	}
	return
}

func (c *Controller) UpdateModule(id uint64, opt UpdateOption, zeroFieldNames ...string) error {
	module := &databases.Modules{
		ModuleID: id,
	}
	err := module.FetchByModuleID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("service.Controller.UpdateModule FetchByModulesID err: %v, id: %d, opt: %+v", err, id, opt)
		}
		return err
	}

	fieldValues := opt.ToUpdateFieldValues(zeroFieldNames...)
	err = module.UpdateByModuleIDWithMap(c.db, fieldValues)
	if err != nil {
		logrus.Errorf("service.Controller.UpdateModule UpdateByModulesIDWithMap err: %v, id: %d, opt: %+v, fieldValues: %+v", err, id, opt, fieldValues)
	}
	return err
}

func (c *Controller) GetModules(condition Condition, offset, limit int64) ([]databases.Modules, int, error) {
	m := &databases.Modules{}
	sqlCondition := condition.ToConditions(c.db)
	list, err := m.List(c.db, sqlCondition, builder.Limit(limit).Offset(offset))
	if err != nil {
		logrus.Errorf("service.Controller.GetModules List err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	count, err := m.Count(c.db, sqlCondition)
	if err != nil {
		logrus.Errorf("service.Controller.GetModules Count err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	return list, count, nil
}

func (c *Controller) GetModuleByModuleID(id uint64) (model *databases.Modules, err error) {
	model = &databases.Modules{
		ModuleID: id,
	}
	err = model.FetchByModuleID(c.db)
	if err != nil && !sqlx.DBErr(err).IsNotFound() {
		logrus.Errorf("service.Controller.GetModuleByModuleID err: %v, id: %d", err, id)
	}
	return
}

func (c *Controller) DeleteModule(id uint64, soft bool) (err error) {
	model := &databases.Modules{
		ModuleID: id,
	}
	if soft {
		err = model.SoftDeleteByModuleID(c.db)
	} else {
		err = model.DeleteByModuleID(c.db)
	}
	if err != nil {
		logrus.Errorf("service.Controller.DeleteModulePermission err: %v, id: %d, soft: %v", err, id, soft)
	}
	return
}
