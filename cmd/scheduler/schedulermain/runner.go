package schedulermain

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	log "github.com/sirupsen/logrus"
	"regexp"
	"sync"
)

func db_runner(wg *sync.WaitGroup, dbid string, n int, app *application) {
	defer wg.Done()
	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
	}).Debug("Worker Started")

	//
	rdsOutput, err := app.rdsClient.StopDBInstance(context.TODO(), &rds.StopDBInstanceInput{
		DBInstanceIdentifier: &dbid,
	})
	if err != nil {
		log.Fatal(err)
	}
	r, _ := regexp.Compile("(?m)(?:DBInstanceStatus: \")(.*?)(?:\")")
	res := r.FindStringSubmatch(awsutil.Prettify(rdsOutput))
	//

	log.WithFields(log.Fields{
		"Worker": n,
		"DB":     dbid,
		"Status": res[1],
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
