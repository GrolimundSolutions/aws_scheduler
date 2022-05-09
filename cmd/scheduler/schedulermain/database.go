package schedulermain

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"time"
)

func (app *application) checkConnection() bool {
	count := 5

	for count > 0 {
		err := app.db.Ping()
		if err == nil {
			log.Info("Connection to Database is OK")
			return true
		}
		count--
		log.WithFields(log.Fields{
			"retry": count,
			"err":   "Can't connect to Database",
		}).Info("Checking connection")
		log.WithFields(log.Fields{
			"retry": count,
			"err":   err,
		}).Debug("Checking connection")
		time.Sleep(time.Second * 3)
	}
	log.Fatal("Can't connect to Database")
	return false
}

func initDB(app *application) {
	migrationsPath := "file://database/PROD_migrations"
	dbConnectionString := "empty"

	if app.ctx.Environment == "development" || app.ctx.Environment == "devl" || app.ctx.Environment == "develop" || app.ctx.Environment == "dev" {
		migrationsPath = "file://database/DEV_migrations"
	}

	if len(app.ctx.DBRootCertPath) == 0 {
		dbConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", app.ctx.DBUser, app.ctx.DBPassword, app.ctx.DBHost, app.ctx.DBPort, app.ctx.DBName)
		log.WithFields(log.Fields{
			"dbConnectionString": fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", app.ctx.DBUser, "******", app.ctx.DBHost, app.ctx.DBPort, app.ctx.DBName),
		}).Debug("initDB")
	} else {
		dbConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?ssl=true&sslrootcert=%s", app.ctx.DBUser, app.ctx.DBPassword, app.ctx.DBHost, app.ctx.DBPort, app.ctx.DBName, app.ctx.DBRootCertPath)
		log.WithFields(log.Fields{
			"dbConnectionString": fmt.Sprintf("postgres://%s:%s@%s:%d/%s?ssl=true&sslrootcert=%s", app.ctx.DBUser, "******", app.ctx.DBHost, app.ctx.DBPort, app.ctx.DBName, app.ctx.DBRootCertPath),
		}).Debug("initDB")
	}

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
