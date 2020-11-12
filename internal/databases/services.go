package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Services --database Config.DB --with-comments
//go:generate eden generate tag Services --defaults=true
// @def primary ID
// @def unique_index U_service_id ServiceID
type Services struct {
	datatypes.PrimaryID
	// 业务ID
	ServiceID uint64 `json:"serviceID,string" db:"f_service_id"`
	BaseService
	datatypes.OperateTime
}

type BaseService struct {
	// 服务标识
	ServiceKey string `json:"serviceKey" db:"f_service_key"`
	// 服务名称
	Name string `json:"name" db:"f_name"`
	// 介绍
	Comment string `json:"comment,omitempty" db:"f_comment,default=''"`
}
