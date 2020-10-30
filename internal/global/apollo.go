package global

import (
	"github.com/eden-framework/apollo"
)

var ApolloConfig = apollo.ApolloBaseConfig{
	AppId:            "srv-identity-platform",
	Host:             "localhost:8080",
	BackupConfigPath: "./apollo_config",
	Cluster:          "default",
}
