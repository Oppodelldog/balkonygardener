package sensor

import (
	"time"

	"github.com/Oppodelldog/balkonygardener/config"

	"github.com/Oppodelldog/balkonygardener/db"
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
				if time.Since(lastSaveTimes[sensorInfo.Name]) > time.Duration(time.Duration(config.Arduino.CaptureInterval)*time.Second) {
					logrus.Infof("saving sensor info: %v, %v", sensorInfo.Name, sensorInfo.Value)
					err := db.SaveFloat(sensorInfo.Name, sensorInfo.Value)
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
