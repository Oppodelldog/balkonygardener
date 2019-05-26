package sensor

import (
	"context"
	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"io"
)

func StartSensorProcessing() {
	arduinoMessages := make(chan string)
	sensorInfoChannel := make(chan *Info)
	go arduinoReader(context.Background(), arduinoMessages, openSerialReader())
	go arduinoMessageDecoder(arduinoMessages, sensorInfoChannel)
	go processSensorInfo(sensorInfoChannel)
}

func openSerialReader() io.ReadCloser {
	device := config.Arduino.Device
	logrus.Infof("Arduino Reader opening connection to %v", device)
	c := &serial.Config{Name: device, Baud: config.Arduino.BaudRate}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	return s
}
