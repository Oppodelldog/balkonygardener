package sensor

import (
	"fmt"
	"github.com/Oppodelldog/balkonygardener/log"
	"io"
	"math/rand"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
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
	log.Infof("Arduino Reader opening connection to %v", device)
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
	ids := []string{"A0", "A1", "A2", "A3", "US"}
	id := ids[rand.Intn(len(ids))]
	message := []byte(fmt.Sprintf("%s:%v\r\n", id, rand.Float32()*600))
	copy(p, message)
	return len(message), nil
}

func (arduinoMockReader) Close() error {
	return nil
}
