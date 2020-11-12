package service

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/sirupsen/logrus"
)

func (c *Controller) CreateService(opt databases.BaseService) (service *databases.Services, err error) {
	id, err := client.GetUniqueID(c.clientID)
	if err != nil {
		return
	}
	service = &databases.Services{
		ServiceID:   id,
		BaseService: opt,
	}
	err = service.Create(c.db)
	if err != nil {
		logrus.Errorf("service.Controller.CreateService err: %v, opt: %+v", err, opt)
	}
	return
}

func (c *Controller) UpdateService(id uint64, opt UpdateOption, zeroFieldNames ...string) error {
	m := &databases.Services{
		ServiceID: id,
	}
	err := m.FetchByServiceID(c.db)
	if err != nil {
		if !sqlx.DBErr(err).IsNotFound() {
			logrus.Errorf("service.Controller.UpdateService FetchByServicesID err: %v, id: %d, opt: %+v", err, id, opt)
		}
		return err
	}

	fieldValues := opt.ToUpdateFieldValues(zeroFieldNames...)
	err = m.UpdateByServiceIDWithMap(c.db, fieldValues)
	if err != nil {
		logrus.Errorf("service.Controller.UpdateService UpdateByServicesIDWithMap err: %v, id: %d, opt: %+v, fieldValues: %+v", err, id, opt, fieldValues)
	}
	return err
}

func (c *Controller) GetServices(condition Condition, offset, limit int64) ([]databases.Services, int, error) {
	m := &databases.Services{}
	sqlCondition := condition.ToConditions(c.db)
	list, err := m.List(c.db, sqlCondition, builder.Limit(limit).Offset(offset))
	if err != nil {
		logrus.Errorf("service.Controller.GetServices List err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	count, err := m.Count(c.db, sqlCondition)
	if err != nil {
		logrus.Errorf("service.Controller.GetServices Count err: %v, condition: %+v", err, condition)
		return nil, 0, err
	}
	return list, count, nil
}

func (c *Controller) GetServiceByServiceID(id uint64) (model *databases.Services, err error) {
	model = &databases.Services{
		ServiceID: id,
	}
	err = model.FetchByServiceID(c.db)
	if err != nil && !sqlx.DBErr(err).IsNotFound() {
		logrus.Errorf("service.Controller.GetServiceByServiceID err: %v, id: %d", err, id)
	}
	return
}

func (c *Controller) DeleteService(id uint64, soft bool) (err error) {
	model := &databases.Services{
		ServiceID: id,
	}
	if soft {
		err = model.SoftDeleteByServiceID(c.db)
	} else {
		err = model.DeleteByServiceID(c.db)
	}
	if err != nil {
		logrus.Errorf("service.Controller.DeleteService err: %v, id: %d, soft: %v", err, id, soft)
	}
	return
}
