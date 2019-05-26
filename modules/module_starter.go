package modules

import "github.com/sirupsen/logrus"

func StartModules() {
	appModules := loadModules()

	logrus.Infof("starting modules")
	for _, appModule := range appModules {
		logrus.Infof("starting %T", appModule)
		appModule.Start()
	}
	logrus.Infof("modules started")
}
