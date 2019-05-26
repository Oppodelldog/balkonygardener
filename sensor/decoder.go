package sensor

import "github.com/sirupsen/logrus"

func arduinoMessageDecoder(arduinoMessages chan string, sensorInfoChannel chan *Info) {
	running := true
	for running {
		select {
		case arduinoInfo, ok := <-arduinoMessages:

			if !ok {
				running = false
				break
			} else {
				err, sensorInfo := parseSensorInfo(arduinoInfo)
				if err != nil {
					logrus.Errorf("could not parse sensor info: %v", err.Error())
				} else {
					//logrus.Info(sensorInfo)
					sensorInfoChannel <- sensorInfo
				}
			}
		}
	}
	close(sensorInfoChannel)
	logrus.Info("arduinoMessageDecoder ended")
}
