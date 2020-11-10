package service

import (
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/sqlx"
)

type Controller struct {
	db       sqlx.DBExecutor
	clientID client_srv_id.ClientSrvIDInterface
}

func NewController(db sqlx.DBExecutor, clientID client_srv_id.ClientSrvIDInterface) *Controller {
	return &Controller{db: db, clientID: clientID}
}
