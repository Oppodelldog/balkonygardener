package sensor

import (
	"time"

	"github.com/Oppodelldog/balkonygardener/config"

	"github.com/sirupsen/logrus"
)

func processSensorInfo(sensorInfoChannel chan *Info) {
	lastSaveTimes := map[string]time.Time{}
	running := true
	for running {
		select {
		case sensorInfo, ok := <-sensorInfoChannel:
			if !ok {
				running = false
				break
			} else {
				if _, ok := lastSaveTimes[sensorInfo.Name]; !ok {
					lastSaveTimes[sensorInfo.Name] = time.Now()
					break
				}
				if time.Since(lastSaveTimes[sensorInfo.Name]) > time.Duration(config.Arduino.CaptureInterval)*time.Second {
					logrus.Debugf("saving sensor info: %v, %v", sensorInfo.Name, sensorInfo.Value)
					var err error = nil
					if err != nil {
						logrus.Errorf("error while saving sensor info: %v", err)
					}
					lastSaveTimes[sensorInfo.Name] = time.Now()
				}
			}
		}
	}
	logrus.Info("processBalkonyInfo ended")
}
