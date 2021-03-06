package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/metricbeat/beater"
//	"github.com/elastic/beats/metricbeat/module/kafka"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/user/metrobeats/include"
	// Comment out the following line to exclude all official metricbeat modules and metricsets
	_ "github.com/elastic/beats/metricbeat/include"
)

var Name = "metrobeats"

func main() {
	if err := beat.Run(Name, "", beater.New); err != nil {
		os.Exit(1)
	}
}
