package sensor

import (
	"testing"
)

func Test_arduinoMessageDecoder(t *testing.T) {

	arduinoMessages := make(chan string)
	sensorInfoChannel := make(chan *Info)
	numberOfReceivedSensorInfosChannel := make(chan int)
	go arduinoMessageDecoder(arduinoMessages, sensorInfoChannel)

	go func(numberOfReceivedSensorInfosChannel chan int) {
		numberOfReceivedSensorInfos := 0
		running := true
		for running {
			select {
			case _, ok := <-sensorInfoChannel:
				if !ok {
					running = false
					break
				}
				numberOfReceivedSensorInfos++
			}
		}
		numberOfReceivedSensorInfosChannel <- numberOfReceivedSensorInfos
	}(numberOfReceivedSensorInfosChannel)

	arduinoMessages <- "A0:123456.12"
	arduinoMessages <- "A1:24432"
	close(arduinoMessages)

	numberOfReceivedSensorInfos := <-numberOfReceivedSensorInfosChannel

	if numberOfReceivedSensorInfos != 2 {
		t.Fatalf("exepcted to receive 2 sensorInfos, but got %v", numberOfReceivedSensorInfos)
	}
}
