package main

import (
	"time"
	"github.com/Oppodelldog/balkonygardener/sensor"
	"github.com/Oppodelldog/balkonygardener/water"
	"github.com/Oppodelldog/balkonygardener/api"
	"github.com/Oppodelldog/balkonygardener"
	"github.com/Oppodelldog/balkonygardener/modules"
)

func main() {

	appModules := loadModules()

	for _, appModule := range appModules {
		appModule.Start();
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func loadModules() []modules.Module{
	return []modules.Module{
		modules.NewRestfulSenorApiModule(),
		modules.NewWateringModule(),
		modules.NewSensorReaderModule(),
	}
}
