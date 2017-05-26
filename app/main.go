package main

import (
	"time"
	"github.com/Oppodelldog/balkonygardener/sensor"
	"github.com/Oppodelldog/balkonygardener/water"
	"github.com/Oppodelldog/balkonygardener/api"
)

func main() {

	sensor.StartSensorProcessing()

	water.StartGardener()

	api.Start()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
