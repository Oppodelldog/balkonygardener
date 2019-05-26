package sensor

func StartSensorProcessing() {
	arduinoMessages := make(chan string)
	sensorInfoChannel := make(chan *Info)
	go arduinoReader(arduinoMessages)
	go arduinoMessageDecoder(arduinoMessages, sensorInfoChannel)
	go processSensorInfo(sensorInfoChannel)
}
