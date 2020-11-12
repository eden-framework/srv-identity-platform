package modules

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/srv-identity-platform/internal/constants/errors"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/modules/service"
)

func init() {
	Router.Register(courier.NewRouter(GetModuleByModuleID{}))
}

// 根据ModuleID获取Module
type GetModuleByModuleID struct {
	httpx.MethodGet

	ModuleID uint64 `name:"moduleID,string" in:"path"`
}

func (req GetModuleByModuleID) Path() string {
	return "/:moduleID"
}

func (req GetModuleByModuleID) Output(ctx context.Context) (result interface{}, err error) {
	data, err := service.NewController(global.Config.SlaveDB, global.ClientConfig.ID).GetModuleByModuleID(req.ModuleID)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			err = errors.ModuleNotFound
			return
		}
		err = errors.InternalError.StatusError().WithDesc(err.Error())
	}
	return data, err
}
