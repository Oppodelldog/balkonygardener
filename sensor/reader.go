package sensor

import (
	"io"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func openReader() io.ReadCloser {
	if config.Arduino.Mock {
		return arduinoMockReader{}
	}

	return openSerialReader()
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

// arduinoMockReader for local development mocks out arduino serial input
type arduinoMockReader struct {
}

func (arduinoMockReader) Read(p []byte) (n int, err error) {
	time.Sleep(1 * time.Second)
	message := "A0:123456.789\r\n"
	p = append(p, []byte(message)...)
	return len(message), nil
}

func (arduinoMockReader) Close() error {
	return nil
}
