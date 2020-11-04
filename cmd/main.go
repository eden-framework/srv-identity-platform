package main

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-framework/sqlx/migration"
	"github.com/eden-framework/srv-identity-platform/internal/modules/providers"
	"github.com/eden-framework/srv-identity-platform/internal/modules/token"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/eden-framework/srv-identity-platform/internal/databases"
	"github.com/eden-framework/srv-identity-platform/internal/global"
	"github.com/eden-framework/srv-identity-platform/internal/routers"
)

var cmdMigrationDryRun bool

func main() {
	app := application.NewApplication(runner,
		application.WithConfig(&global.Config),
		application.WithApollo(&global.ApolloConfig),
		application.WithConfig(&global.ClientConfig),
		application.WithConfig(&databases.Config),
		application.WithConfig(&global.CacheConfig),
		application.WithConfig(&global.ProviderConfig),
		application.WithInitializer(true, providers.Initializer, token.Initializer))

	cmdMigrate := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			migrate(&migration.MigrationOpts{
				DryRun: cmdMigrationDryRun,
			})
		},
	}
	cmdMigrate.Flags().BoolVarP(&cmdMigrationDryRun, "dry", "d", false, "migrate --dry")
	app.AddCommand(cmdMigrate)

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	logrus.SetLevel(global.Config.LogLevel)
	go global.Config.GRPCServer.Serve(ctx, routers.Router)
	return global.Config.HTTPServer.Serve(ctx, routers.Router)
}

func migrate(opts *migration.MigrationOpts) {
	if err := migration.Migrate(global.Config.MasterDB, opts); err != nil {
		panic(err)
	}
}
