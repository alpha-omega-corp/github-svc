package main

import (
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/services/database"
	"github.com/alpha-omega-corp/services/server"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	v := viper.New()
	cManager := server.NewConfigManager(v)
	c, err := cManager.HostsConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := database.NewHandler(c.Docker.Dsn)

	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbHandler.Database())

	appCli := &cli.App{
		Name:  "docker-svc",
		Usage: "bootstrap the service",
		Commands: []*cli.Command{
			migrateCommand(dbHandler.Database()),
		},
	}

	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func migrateCommand(db *bun.DB) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					fmt.Print(db.Stats())
					migrator := migrate.NewMigrator(db, migrate.NewMigrations())
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "reset",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					return db.ResetModel(c.Context, (*models.ContainerPackage)(nil))
				},
			},
		},
	}
}
