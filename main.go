package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

type application struct {
	db        *sql.DB
	ctx       Config
	day       int
	hour      int
	databases []DatabaseItem
	clusters  []ClusterItem
}

func main() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ll, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		ll = log.WarnLevel
	}

	// Define the Logger
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// get the log level from the config, set to warning if not set
	log.SetLevel(ll)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", psqlconn)
	defer db.Close()
	CheckError(err)

	log.WithFields(log.Fields{
		"day":                getDayOfWeek(),
		"hour":               getActuallyHour(),
		"DB_Host":            config.DBHost,
		"DB_Name":            config.DBName,
		"DB_Table":           config.DBTable,
		"DB_Port":            config.DBPort,
		"DB_User":            config.DBUser,
		"DB_Password":        "*****",
		"log_level":          config.LogLevel,
		"AwsRegion":          config.AwsRegion,
		"AwsAccessKeySize":   len(config.AwsAccessKey),
		"AwsAccessKeyIdSize": len(config.AwsAccessKeyId),
	}).Info("Processing all configs and Connections")

	app := &application{
		db:        db,
		ctx:       config,
		day:       getDayOfWeek(),
		hour:      getActuallyHour(),
		databases: nil,
		clusters:  nil,
	}

	// Check and retry the Connection to the Database
	app.checkConnection()

	// Select all Items from the DB with the needed values (Day.now, Hour.now)
	app.loadDatabaseInfos()

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
