package schedulermain

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

func db_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.Debugf("DB_Worker: %d: Started, DB: %s", n, dbid)
	// Do some Stuff
	log.Debugf("DB_Worker: %d: Done, DB: %s", n, dbid)
}

func cluster_runner(wg *sync.WaitGroup, dbid string, n int) {
	defer wg.Done()
	log.Debugf("Cluster_Worker: %d: Started, DB: %s", n, dbid)
	// Do some Stuff
	log.Debugf("Cluster_Worker: %d: Done, DB: %s", n, dbid)
}
