package schedulermain

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

func db_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
	}).Debug("Worker Started")

	// Do some Stuff

	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
	}).Debug("Worker Done")
}

func cluster_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.WithFields(log.Fields{
		"Worker":  n,
		"Cluster": dbid,
	}).Debug("Worker Started")

	// Do some Stuff

	log.WithFields(log.Fields{
		"Worker":  n,
		"Cluster": dbid,
	}).Debug("Worker Done")
}
