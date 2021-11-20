package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
)

func (app *application) loadDatabaseInfos() {
	query := fmt.Sprintf(
		"SELECT dbid, type, day, hour, action FROM %s WHERE day=%d AND hour=%d",
		app.ctx.DBTable, app.day, app.hour)

	rows, err := app.db.Query(query)
	defer rows.Close()
	CheckError(err)

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
		log.Debugf("Starting DB_Worker: %d for dbid: %s", i, db.DbId)
		go func(group *sync.WaitGroup, db string, number int) {
			db_runner(&wg, db, i)
		}(&wg, db.DbId, i)
	}

	// Fill and start Workers for Clusters
	for j, cluster := range app.clusters {
		log.Debugf("Starting Cluster_Worker: %d for dbid: %s", j, cluster.DbId)
		go func(group *sync.WaitGroup, cluster string, number int) {
			cluster_runner(&wg, cluster, j)
		}(&wg, cluster.DbId, j)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

}

func db_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.Debugf("DB_Worker: %d: Started, DB: %s", n, dbid)
	// Do some Stuff
}

func cluster_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.Debugf("Cluster_Worker: %d: Started, DB: %s", n, dbid)
	// Do some Stuff
}
