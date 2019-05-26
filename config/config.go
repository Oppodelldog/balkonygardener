package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var Arduino ArduinoConfig
var Db DbConfig
var Watering WateringConfig

type ArduinoConfig struct {
	Device   string
	BaudRate int
}

type DbConfig struct {
	Filename string
}

type WateringConfig map[string]WateringEntryConfig

type WateringEntryConfig struct {
	Duration time.Duration `yaml:"Duration"`
	Comment  string        `yaml:"Name"`
	Time     string        `yaml:"Time"`
}

func init() {
	err := envconfig.Process("BG_ARDUINO", &Arduino)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_DB", &Db)
	if err != nil {
		log.Fatal(err.Error())
	}
	if configFile, ok := os.LookupEnv("BG_WATERING_CONFIG"); ok {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			logrus.Errorf("could not read watring config: %v", err)
		} else {
			err := yaml.Unmarshal([]byte(data), &Watering)
			if err != nil {
				logrus.Errorf("could parse watring config: %v", err)
			}
		}
	}
	logrus.Info("Config")
	logrus.Infof("Arduino : %+v", Arduino)
	logrus.Infof("Db      : %+v", Db)
	logrus.Infof("Watering: %+v", Watering)
}
