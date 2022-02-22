package main

import (
	"github.com/GrolimundSolutions/aws_scheduler/cmd/scheduler/schedulermain"
)

func main() {
	// Show build information
	//log.Infof("Version: %s, Build: %s", Version, Build)
	schedulermain.Run()
}
