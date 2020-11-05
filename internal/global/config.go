package global

import (
	"github.com/eden-framework/courier/transport_grpc"
	"github.com/eden-framework/courier/transport_http"
	"github.com/eden-framework/eden-framework/pkg/client/mysql"
	"github.com/profzone/envconfig"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eden-framework/srv-identity-platform/internal/databases"
)

var Config = struct {
	LogLevel logrus.Level

	// db
	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	// administrator
	GRPCServer *transport_grpc.ServeGRPC
	HTTPServer *transport_http.ServeHTTP

	// service config
	EnableRegister           bool
	SecureCodeDefaultExpire  envconfig.Duration
	AccessTokenDefaultExpire envconfig.Duration

	JwtIssuer     string
	JwtPrivateKey envconfig.Password
}{
	LogLevel: logrus.DebugLevel,

	MasterDB: &mysql.MySQL{Database: databases.Config.DB},
	SlaveDB:  &mysql.MySQL{Database: databases.Config.DB},

	GRPCServer: &transport_grpc.ServeGRPC{
		Port: 8900,
	},
	HTTPServer: &transport_http.ServeHTTP{
		Port:     8800,
		WithCORS: true,
	},

	EnableRegister:           true,
	SecureCodeDefaultExpire:  envconfig.Duration(time.Minute),
	AccessTokenDefaultExpire: envconfig.Duration(2 * time.Hour),

	JwtIssuer:     "",
	JwtPrivateKey: "",
}
