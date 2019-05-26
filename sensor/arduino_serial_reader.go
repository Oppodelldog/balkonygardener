package sensor

import (
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func arduinoReader(arduinoMessages chan string) {

	device := "/dev/ttyACM0"
	logrus.Infof("Arduino Reader opening connection to %v", device)
	c := &serial.Config{Name: device, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	logrus.Info("connection established, reading data...")

	readBuffer := make([]byte, 100)
	var messageBuffer []byte
	for {
		n, err := s.Read(readBuffer)
		if err != nil {
			logrus.Error(err)
			break
		}
		readBytes := readBuffer[:n]
		messageBuffer = append(messageBuffer, readBytes...)
		eol := 0

		for index, b := range messageBuffer {
			if b == '\r' {
				eol++
			}
			if b == '\n' {
				eol++

				if eol == 2 {
					info := string(messageBuffer[:index-1])
					arduinoMessages <- info
					messageBuffer = messageBuffer[index+1:]
					break
				}
			}
		}
	}
}
