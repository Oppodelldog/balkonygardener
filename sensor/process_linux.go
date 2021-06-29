package sensor

import (
	"time"

	"github.com/Oppodelldog/balkonygardener/log"

	"github.com/Oppodelldog/balkonygardener/config"
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
					log.Debugf("saving sensor info: %v, %v", sensorInfo.Name, sensorInfo.Value)
					var err error = nil
					if err != nil {
						log.Errorf("error while saving sensor info: %v", err)
					}
					lastSaveTimes[sensorInfo.Name] = time.Now()
				}
			}
		}
	}
	log.Info("processBalkonyInfo ended")
}
