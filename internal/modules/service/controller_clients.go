package service

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

func (c *Controller) CreateModuleClient(opt databases.BaseClient) (cli *databases.ModuleClients, err error) {
	id, err := client.GetUniqueID(c.clientID)
	if err != nil {
		return
	}
	cli = &databases.ModuleClients{
		ClientID:   id,
		BaseClient: opt,
	}
	err = cli.Create(c.db)
	if err != nil {
		logrus.Errorf("service.Controller.CreateModuleClient err: %v, opt: %+v", err, opt)
	}
	return
}

func (c *Controller) UpdateModuleClient(id uint64, opt UpdateOption, zeroFieldNames ...string) error {
	m := &databases.ModuleClients{
		ClientID: id,
	}
	err := m.FetchByClientID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("service.Controller.UpdateModuleClient FetchByClientID err: %v, id: %d, opt: %+v", err, id, opt)
		}
		return err
	}

	fieldValues := opt.ToUpdateFieldValues(zeroFieldNames...)
	err = m.UpdateByClientIDWithMap(c.db, fieldValues)
	if err != nil {
		logrus.Errorf("service.Controller.UpdateModuleClient UpdateByClientIDWithMap err: %v, id: %d, opt: %+v, fieldValues: %+v", err, id, opt, fieldValues)
	}
	return err
}

func (c *Controller) GetModuleClients(condition Condition, offset, limit int64) ([]databases.ModuleClients, int, error) {
	m := &databases.ModuleClients{}
	sqlCondition := condition.ToConditions(c.db)
	list, err := m.List(c.db, sqlCondition, builder.Limit(limit).Offset(offset))
	if err != nil {
		logrus.Errorf("service.Controller.GetModuleClients List err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	count, err := m.Count(c.db, sqlCondition)
	if err != nil {
		logrus.Errorf("service.Controller.GetModuleClients Count err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	return list, count, nil
}

func (c *Controller) GetModuleClientByClientID(id uint64) (model *databases.ModuleClients, err error) {
	model = &databases.ModuleClients{
		ClientID: id,
	}
	err = model.FetchByClientID(c.db)
	if err != nil && !sqlx.DBErr(err).IsNotFound() {
		logrus.Errorf("service.Controller.GetModuleClientByClientID err: %v, id: %d", err, id)
	}
	return
}
