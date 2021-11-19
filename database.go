package main

import log "github.com/sirupsen/logrus"

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
		}).Info("Checking connection")
	}
	log.Fatal("Can't connect to Database")
	return false
}
