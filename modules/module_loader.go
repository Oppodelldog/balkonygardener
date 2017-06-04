package modules

func loadModules() []Module{
	return []Module{
		newRestfulSenorApiModule(),
		newWateringModule(),
		newSensorReaderModule(),
	}
}
