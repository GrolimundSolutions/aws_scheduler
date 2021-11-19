package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
	return
}
