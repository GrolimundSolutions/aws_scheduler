// RDS Status list: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/accessing-monitoring.html#Overview.DBInstance.Status

package schedulermain

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	log "github.com/sirupsen/logrus"
	"regexp"
)

func getDBInstanceStatus(output string) (string, error) {
	r, _ := regexp.Compile(`(?m)(?:DBInstanceStatus: ")(.*?)(?:")`)
	status := r.FindStringSubmatch(output)
	log.WithFields(log.Fields{
		"status_raw": status,
	}).Debug("getDBInstanceStatus")

	if len(status) > 1 {
		return status[1], nil
	} else {
		return "", errors.New("DBInstanceStatus not found")
	}
}

func getDBClusterStatus(output string) (string, error) {
	r, _ := regexp.Compile(`(?m)(?:DBClusterStatus: ")(.*?)(?:")`)
	status := r.FindStringSubmatch(output)
	log.WithFields(log.Fields{
		"status_raw": status,
	}).Debug("getDBClusterStatus")

	if len(status) > 1 {
		return status[1], nil
	} else {
		return "", errors.New("DBClusterStatus not found")
	}
}

func DescribeRDS_DB(database string, client *rds.Client) (string, error) {

	rdsOutput, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(database),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":  database,
			"error": err,
		}).Error("Error describe RDS-Database")
	}

	status, err := getDBInstanceStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":  database,
			"error": err,
		}).Error("Error get DBInstanceStatus")
	}

	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
	}).Info("RDS-Database Description")
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-Database Description")

	return status, err
}

func DescribeRDS_Cluster(cluster string, client *rds.Client) (string, error) {

	rdsOutput, err := client.DescribeDBClusters(context.TODO(), &rds.DescribeDBClustersInput{
		DBClusterIdentifier: aws.String(cluster),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":  cluster,
			"error": err,
		}).Error("Error describe RDS-Cluster")
	}

	status, err := getDBClusterStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":  cluster,
			"error": err,
		}).Error("Error get DBClusterStatus")
	}

	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
	}).Info("RDS-Cluster Description")
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-Cluster Description")

	return status, err
}

// Database

func StartRDS_DB(database string, rdsClient *rds.Client, n int) {

	rdsOutput, err := rdsClient.StartDBInstance(context.TODO(), &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(database),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   database,
			"error":  err,
			"Worker": n,
		}).Error("Error starting RDS")
	}

	status, err := getDBInstanceStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   database,
			"error":  err,
			"Worker": n,
		}).Error("Error get DBInstanceStatus")
	}
	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
		"Worker": n,
	}).Info("RDS-DB Starting")
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
		"Worker": n,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-DB Starting")

}

func StopRDS_DB(database string, rdsClient *rds.Client, n int) {

	rdsOutput, err := rdsClient.StopDBInstance(context.TODO(), &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(database),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   database,
			"error":  err,
			"Worker": n,
		}).Error("Error stopping RDS")
	}

	status, err := getDBInstanceStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   database,
			"error":  err,
			"Worker": n,
		}).Error("Error get DBInstanceStatus")
	}
	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
		"Worker": n,
	}).Info("RDS-DB Stopping")
	log.WithFields(log.Fields{
		"DBId":   database,
		"status": status,
		"Worker": n,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-DB Stopping")

}

// Cluster

func StartRDS_Cluster(cluster string, rdsClient *rds.Client, n int) {

	rdsOutput, err := rdsClient.StartDBCluster(context.TODO(), &rds.StartDBClusterInput{
		DBClusterIdentifier: aws.String(cluster),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   cluster,
			"error":  err,
			"Worker": n,
		}).Error("Error starting RDS-Cluster")
	}

	status, err := getDBClusterStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   cluster,
			"error":  err,
			"Worker": n,
		}).Error("Error get DBClusterInstanceStatus")
	}
	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
		"Worker": n,
	}).Info("RDS-Cluster Starting")
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
		"Worker": n,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-Cluster Starting")

}

func StopRDS_Cluster(cluster string, rdsClient *rds.Client, n int) {

	rdsOutput, err := rdsClient.StopDBCluster(context.TODO(), &rds.StopDBClusterInput{
		DBClusterIdentifier: aws.String(cluster),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   cluster,
			"error":  err,
			"Worker": n,
		}).Error("Error stopping RDS-Cluster")
	}

	status, err := getDBClusterStatus(awsutil.Prettify(rdsOutput))
	if err != nil {
		log.WithFields(log.Fields{
			"DBId":   cluster,
			"error":  err,
			"Worker": n,
		}).Error("Error get DBClusterInstanceStatus")
	}
	// as workaround to print the result, we need to use the V1 SDK (awsutil.Prettify()) Function
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
		"Worker": n,
	}).Info("RDS-Cluster stopping")
	log.WithFields(log.Fields{
		"DBId":   cluster,
		"status": status,
		"Worker": n,
		"raw":    awsutil.Prettify(rdsOutput),
	}).Debug("RDS-Cluster stopping")

}
