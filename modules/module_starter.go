package modules

import "github.com/Sirupsen/logrus"

func StartModules() {
	appModules := loadModules()

	logrus.Infof("starting modules")
	for _, appModule := range appModules {
		logrus.Infof("starting %T", appModule)
		appModule.Start();
	}
	logrus.Infof("modules started")
}
