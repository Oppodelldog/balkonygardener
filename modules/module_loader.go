package modules

import "github.com/Oppodelldog/balkonygardener/config"

func loadModules() []Module {
	var modules []Module
	modules = append(modules, newRestfulSenorApiModule())
	modules = append(modules, newWateringModule())

	if (config.Arduino.Device != "" && config.Arduino.BaudRate > 0) || (config.Arduino.Mock == true) {
		modules = append(modules, newSensorReaderModule())
	}

	return modules
}
