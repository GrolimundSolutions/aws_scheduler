package schedulermain

import (
	"fmt"
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
			"Worker": 1,
			"DB":     db.DbId,
		}).Debug("Worker Starting")

		go func(group *sync.WaitGroup, db string, number int) {
			db_runner(group, db, number, app)
		}(&wg, db.DbId, i)
	}

	// Fill and start Workers for Clusters
	for j, cluster := range app.clusters {
		log.WithFields(log.Fields{
			"Worker": j,
			"DB":     cluster.DbId,
		}).Debug("Worker Started")
		go func(group *sync.WaitGroup, cluster string, number int) {
			cluster_runner(group, cluster, number)
		}(&wg, cluster.DbId, j)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

}
