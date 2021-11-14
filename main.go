package main

import (
	"database/sql"
	"fmt"
	"github.com/GrolimundSolutions/psql_example/utils"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "scheduler_db"
	tablename = "table_schedule"
)
*/

func getDayOfWeek() int {
	return int(time.Now().Weekday())
}
func getActuallyHour() int {
	return time.Now().Hour()
}

type DbItem struct {
	DbId   string
	day    int
	hour   int
	action string
}

type ClusterItem struct {
	DbId   string
	day    int
	hour   int
	action string
}

func main() {
	utils.Test()
	os.Exit(0)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ll, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		ll = log.DebugLevel
	}

	log.SetLevel(ll)

	var databases []DbItem
	var clusters []ClusterItem

	log.WithFields(log.Fields{
		"day":  getDayOfWeek(),
		"hour": getActuallyHour(),
	}).Info("Start Scheduler")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	log.WithFields(log.Fields{
		"DB_Host":     config.DBHost,
		"DB_Name":     config.DBName,
		"DB_Port":     config.DBPort,
		"DB_User":     config.DBUser,
		"DB_Password": "*****",
	}).Debug("Connected to DB")

	defer db.Close()

	query := fmt.Sprintf("SELECT dbid, type, day, hour, action FROM %s WHERE day=%d AND hour=%d", config.DBTable, getDayOfWeek(), getActuallyHour())
	rows, err := db.Query(query)
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var (
			dbid   string
			dbType string
			day    int
			hour   int
			action string
		)

		err = rows.Scan(&dbid, &dbType, &day, &hour, &action)
		CheckError(err)

		if dbType == "cluster" {
			log.WithFields(log.Fields{
				"dbid":   dbid,
				"day":    day,
				"hour":   hour,
				"action": action,
			}).Debug("Cluster found")
			clusters = append(clusters, ClusterItem{dbid, day, hour, action})
		} else if dbType == "db" {
			log.WithFields(log.Fields{
				"dbid":   dbid,
				"day":    day,
				"hour":   hour,
				"action": action,
			}).Debug("DB found")
			databases = append(databases, DbItem{dbid, day, hour, action})
		} else {
			log.WithFields(log.Fields{
				"dbid":   dbid,
				"day":    day,
				"hour":   hour,
				"action": action,
			}).Fatal("Unknown DB-Type found")
		}

	}
	CheckError(err)
	fmt.Println("Clusters:", clusters)
	fmt.Println("Databases:", databases)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
