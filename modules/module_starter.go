package modules

import (
	"github.com/Oppodelldog/balkonygardener/log"
)

func StartModules() {
	appModules := loadModules()

	log.Infof("starting modules")
	for _, appModule := range appModules {
		log.Infof("starting %T", appModule)
		appModule.Start()
	}
	log.Infof("modules started")
}
