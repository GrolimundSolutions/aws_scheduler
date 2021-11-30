package schedulermain

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	log "github.com/sirupsen/logrus"
	"sync"
)

func (app *application) loadDatabaseInfos() {
	query := `SELECT dbid, type, day, hour, action FROM table_schedule WHERE day=$1 AND hour=$2`
	rows, err := app.db.Query(query, app.day, app.hour)
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
			}).Info("Found item type Cluster")
			app.clusters = append(app.clusters, ClusterItem{dbid, day, hour, action})
		} else if dbType == "db" {
			log.WithFields(log.Fields{
				"dbid":   dbid,
				"day":    day,
				"hour":   hour,
				"action": action,
			}).Info("Found item type Database")
			app.databases = append(app.databases, DatabaseItem{dbid, day, hour, action})
		} else {
			log.WithFields(log.Fields{
				"dbid":   dbid,
				"day":    day,
				"hour":   hour,
				"action": action,
			}).Error("Unknown Type found")
		}

	}
	CheckError(err)
	log.Debug("Clusters:", app.clusters)
	log.Debug("Databases:", app.databases)
}

func (app *application) startScheduling() {
	var wg sync.WaitGroup
	wg.Add(len(app.clusters) + len(app.databases))

	// Fill and start Workers for Databases
	for i, db := range app.databases {
		log.WithFields(log.Fields{
			"Worker": i,
			"DB":     db.DbId,
		}).Debug("Worker Starting")

		go func(group *sync.WaitGroup, db string, number int, action string, client *rds.Client) {
			db_runner(group, db, number, action, client)
		}(&wg, db.DbId, i, db.action, app.rdsClient)
	}

	// Fill and start Workers for Clusters
	for j, cluster := range app.clusters {
		log.WithFields(log.Fields{
			"Worker": j,
			"DB":     cluster.DbId,
		}).Debug("Worker Started")
		go func(group *sync.WaitGroup, cluster string, number int, action string, client *rds.Client) {
			cluster_runner(group, cluster, number, action, client)
		}(&wg, cluster.DbId, j, cluster.action, app.rdsClient)
	}

	log.Info("Main: Waiting for workers to finish")
	wg.Wait()
	log.Info("Main: All workers finished")

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
