package modules

import "github.com/Oppodelldog/balkonygardener/sensor"

type sensorReader struct{}

func (module *sensorReader) Start() {
	go sensor.StartSensorProcessing()
}
func newSensorReaderModule() *sensorReader {
	return &sensorReader{}
}
