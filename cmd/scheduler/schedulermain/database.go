package schedulermain

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"time"
)

func (app *application) checkConnection() bool {
	count := 5

	for count > 0 {
		if app.db.Ping() == nil {
			log.Info("Connection to Database is OK")
			return true
		}
		count--
		log.WithFields(log.Fields{
			"retry": count,
			"err":   "Can't connect to Database",
		}).Warning("Checking connection")
		time.Sleep(time.Second * 2)
	}
	log.Fatal("Can't connect to Database")
	return false
}

func (app *application) initScheduler() {
	initDB(app)
	initAwsClients(app)
}

func initAwsClients(app *application) {
	// Load the Shared AWS Configuration (~/.aws/config)
	// or load from the environment variables
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	app.rdsClient = rds.NewFromConfig(cfg)
	log.Info("RDS Client initialized")
}

func initDB(app *application) {
	migrationsPath := "file://database/PROD_migrations"

	if app.ctx.Environment == "development" || app.ctx.Environment == "devl" || app.ctx.Environment == "develop" || app.ctx.Environment == "dev" {
		migrationsPath = "file://database/DEV_migrations"
	}

	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", app.ctx.DBUser, app.ctx.DBPassword, app.ctx.DBHost, app.ctx.DBPort, app.ctx.DBName)

	m, err := migrate.New(
		migrationsPath,
		dbConnectionString)
	if err != nil {
		log.Fatal("new: ", err)
	}
	if err := m.Up(); err != nil {
		log.Info("Start migrations")
		if err.Error() == "no change" {
			log.Info("INITIAL-Database --> No changes in database")
		} else {
			log.Fatal("up: ", err)
		}
	}
	log.Info("Migrations are done")
}
