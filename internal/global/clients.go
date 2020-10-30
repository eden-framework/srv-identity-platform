package global

import (
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/courier/client"
)

var ClientConfig = struct {
	ID *client_srv_id.ClientSrvID
}{
	ID: &client_srv_id.ClientSrvID{
		Client: &client.Client{
			Host: "localhost",
			Mode: "http",
			Port: 8800,
		},
	},
}
