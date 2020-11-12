package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model ModuleClients --database Config.DB --with-comments
//go:generate eden generate tag ModuleClients --defaults=true
// @def primary ID
// @def unique_index U_clients_id ClientID
// @def unique_index U_key AccessKey
// @def index I_module ModuleID
type ModuleClients struct {
	datatypes.PrimaryID
	// 业务ID
	ClientID uint64 `json:"clientID,string" db:"f_client_id"`
	// 所属模块
	ModuleID uint64 `json:"moduleID,string" db:"f_module_id"`
	// AccessKey
	AccessKey string `json:"accessKey" db:"f_access_key"`
	// AccessSecret
	AccessSecret string `json:"accessSecret" db:"f_accessSecret"`
	BaseClient
	datatypes.OperateTime
}

type BaseClient struct {
	// Endpoint
	Endpoint string `json:"endpoint" db:"f_endpoint"`
}
