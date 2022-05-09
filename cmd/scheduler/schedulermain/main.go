package schedulermain

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type application struct {
	db        *sql.DB
	ctx       *Config
	day       int
	hour      int
	databases []DatabaseItem
	clusters  []ClusterItem
	location  *time.Location
	rdsClient *rds.Client
}

// Version is provided by ldflags
var Version = "unspecified"

// Build is provided by ldflags
var Build = "unspecified"

func Run() {

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

	log.Infof("Version: %s, Build: %s", Version, Build)

	psqlconn := "empty"
	if len(config.DBRootCertPath) == 0 {
		psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
		log.WithFields(log.Fields{
			"dbConnectionString": fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, "*****", config.DBName),
		}).Debug("psqlconn")
	} else {
		psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ssl=true sslrootcert=%s",
			config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBRootCertPath)
		log.WithFields(log.Fields{
			"dbConnectionString": fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ssl=true sslrootcert=%s", config.DBHost, config.DBPort, config.DBUser, "*****", config.DBName, config.DBRootCertPath),
		}).Debug("psqlconn")
	}

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	// Define the TimeZone given from the config (TZ)
	loc, _ := time.LoadLocation(config.TZ)
	t := time.Now().In(loc)
	zone, offset := t.Zone()

	app := &application{
		db:        db,
		ctx:       &config,
		day:       getDayOfWeek(loc),
		hour:      getActuallyHour(loc),
		databases: nil,
		clusters:  nil,
		location:  loc,
		rdsClient: nil,
	}

	log.WithFields(log.Fields{
		"Environment":        config.Environment,
		"day":                getDayOfWeek(app.location),
		"hour":               getActuallyHour(app.location),
		"DB_Host":            config.DBHost,
		"DB_Name":            config.DBName,
		"DB_Port":            config.DBPort,
		"DB_User":            config.DBUser,
		"DB_Password":        "*****",
		"DB_RootCertPath":    config.DBRootCertPath,
		"log_level":          config.LogLevel,
		"AwsRegion":          config.AwsRegion,
		"AwsAccessKeySize":   len(config.AwsAccessKey),
		"AwsAccessKeyIdSize": len(config.AwsAccessKeyId),
		"TimeZone":           zone,
		"TimeZoneOffset":     offset,
	}).Info("Processing all configs and Connections")

	// Check and retry the Connection to the Database
	app.checkConnection()

	app.initScheduler()

	// Select all Items from the DB with the needed values (Day.now, Hour.now)
	app.loadDatabaseInfos()

	// Start the scheduler with the given values
	app.startScheduling()

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
