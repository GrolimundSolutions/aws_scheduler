package schedulermain

import (
	"github.com/aws/aws-sdk-go-v2/service/rds"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func db_runner(wg *sync.WaitGroup, dbid string, n int, action string, client *rds.Client) {
	defer wg.Done()
	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
	}).Debug("Worker Started")

	switch action {
	case "start":
		log.WithFields(log.Fields{
			"Worker": n,
			"DB":     dbid,
			"Action": action,
		}).Debug("Starting DB")

		// 1. Get actual Status from the RDS
		status, err := DescribeRDS_DB(dbid, client)
		log.WithFields(log.Fields{
			"Worker": n,
			"DB":     dbid,
			"Action": action,
			"Status": status,
		}).Debug("Starting DB")
		if err != nil {
			return
		}

		// 2. If Status is not "stopped", then raise an error and Skipping
		if status != "stopped" {
			log.WithFields(log.Fields{
				"Worker": n,
				"DB":     dbid,
				"Action": action,
				"Status": status,
			}).Error("DB is not stopped")
			return
		}

		// -- Else, Start the DB
		StartRDS_DB(dbid, client, n)

		// Slow down the process
		time.Sleep(100 * time.Millisecond)

		// 3. Get the Status from Response
		status, err = DescribeRDS_DB(dbid, client)
		if err != nil {
			return
		}

		// -- If Status is "starting", All OK
		if status != "starting" {
			log.WithFields(log.Fields{
				"Worker": n,
				"DB":     dbid,
				"Action": action,
				"Status": status,
			}).Error("DB cant Start")
			return
		}

	case "stop":
		log.WithFields(log.Fields{
			"Worker": n,
			"DB":     dbid,
			"Action": action,
		}).Debug("Stopping DB")

		// 1. Get actual Status from the RDS
		status, err := DescribeRDS_DB(dbid, client)
		log.WithFields(log.Fields{
			"Worker": n,
			"DB":     dbid,
			"Action": action,
			"Status": status,
		}).Debug("Stopping DB")
		if err != nil {
			return
		}

		// 2. If Status is not "available", then raise an error and Skipping
		if status != "available" {
			log.WithFields(log.Fields{
				"Worker": n,
				"DB":     dbid,
				"Action": action,
				"Status": status,
			}).Error("DB is not running")
			return
		}

		// -- Else, Start the DB
		StopRDS_DB(dbid, client, n)

		// Slow down the process
		time.Sleep(100 * time.Millisecond)

		// 3. Get the Status from Response
		status, err = DescribeRDS_DB(dbid, client)
		if err != nil {
			return
		}

		// -- If Status is "stopping", All OK
		if status != "stopping" {
			log.WithFields(log.Fields{
				"Worker": n,
				"DB":     dbid,
				"Action": action,
				"Status": status,
			}).Error("DB cant Stopped")
			return
		}

	}

	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
	}).Debug("Worker Done")

}

func cluster_runner(wg *sync.WaitGroup, dbid string, n int, action string, client *rds.Client) {
	defer wg.Done()
	log.WithFields(log.Fields{
		"Worker":  n,
		"Cluster": dbid,
	}).Debug("Worker Started")

	switch action {
	case "start":
		log.WithFields(log.Fields{
			"Worker":  n,
			"Cluster": dbid,
			"Action":  action,
		}).Debug("Starting Cluster")

		// 1. Get actual Status from the RDS
		status, err := DescribeRDS_Cluster(dbid, client)
		log.WithFields(log.Fields{
			"Worker":  n,
			"Cluster": dbid,
			"Action":  action,
			"Status":  status,
		}).Debug("Starting Cluster")
		if err != nil {
			return
		}

		// 2. If Status is not "stopped", then raise an error and Skipping
		if status != "stopped" {
			log.WithFields(log.Fields{
				"Worker":  n,
				"Cluster": dbid,
				"Action":  action,
				"Status":  status,
			}).Error("Cluster is not stopped")
			return
		}

		// -- Else, Start the DB
		StartRDS_Cluster(dbid, client, n)

		// Slow down the process
		time.Sleep(100 * time.Millisecond)

		// 3. Get the Status from Response
		status, err = DescribeRDS_Cluster(dbid, client)
		if err != nil {
			return
		}

		// -- If Status is "starting", All OK
		if status != "starting" {
			log.WithFields(log.Fields{
				"Worker":  n,
				"Cluster": dbid,
				"Action":  action,
				"Status":  status,
			}).Error("Cluster cant Start")
			return
		}

	case "stop":
		log.WithFields(log.Fields{
			"Worker":  n,
			"Cluster": dbid,
			"Action":  action,
		}).Debug("Stopping Cluster")

		// 1. Get actual Status from the RDS
		status, err := DescribeRDS_DB(dbid, client)
		log.WithFields(log.Fields{
			"Worker":  n,
			"Cluster": dbid,
			"Action":  action,
			"Status":  status,
		}).Debug("Stopping Cluster")
		if err != nil {
			return
		}

		// 2. If Status is not "available", then raise an error and Skipping
		if status != "available" {
			log.WithFields(log.Fields{
				"Worker":  n,
				"Cluster": dbid,
				"Action":  action,
				"Status":  status,
			}).Error("Cluster is not running")
			return
		}

		// -- Else, Start the DB
		StopRDS_Cluster(dbid, client, n)

		// Slow down the process
		time.Sleep(100 * time.Millisecond)

		// 3. Get the Status from Response
		status, err = DescribeRDS_Cluster(dbid, client)
		if err != nil {
			return
		}

		// -- If Status is "stopping", All OK
		if status != "stopping" {
			log.WithFields(log.Fields{
				"Worker":  n,
				"Cluster": dbid,
				"Action":  action,
				"Status":  status,
			}).Error("Cluster cant Stopped")
			return
		}

	}

	log.WithFields(log.Fields{
		"Worker":  n,
		"Cluster": dbid,
	}).Debug("Worker Done")
}
