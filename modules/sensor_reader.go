package modules

import "github.com/Oppodelldog/balkonygardener/sensor"

type SensorReaderModule struct{}

func (module *SensorReaderModule )Start() {
	go sensor.StartSensorProcessing()
}
func NewSensorReaderModule() *SensorReaderModule{
	return &SensorReaderModule{}
}
