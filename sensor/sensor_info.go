package sensor

import (
	"context"
)

func StartSensorProcessing() {
	arduinoMessages := make(chan string)
	sensorInfoChannel := make(chan *Info)
	go receiveArduinoMessages(context.Background(), arduinoMessages, openReader())
	go arduinoMessageDecoder(arduinoMessages, sensorInfoChannel)
	go processSensorInfo(sensorInfoChannel)
}
