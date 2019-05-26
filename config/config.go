package config

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Arduino ArduinoConfig
var Db DbConfig
var Watering WateringConfig
var Gpio GpioConfig
var Log LogConfig
var Frontend FrontendConfig

type LogConfig struct {
	Level int
}

type ArduinoConfig struct {
	Device          string
	BaudRate        int
	Mock            bool
	CaptureInterval int
}

type FrontendConfig struct {
	IndexFile string
}

type GpioConfig struct {
	Mock bool
}

type DbConfig struct {
	Filename string
}

type WateringConfig map[string]WateringEntryConfig

func (c WateringConfig) FindConfig(pinName string) *WateringEntryConfig {
	for name, waterConfig := range c {
		if name == pinName {
			return &waterConfig
		}
	}
	return nil
}

type WateringEntryConfig struct {
	Duration time.Duration `yaml:"Duration"`
	Comment  string        `yaml:"Name"`
	Time     string        `yaml:"Time"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = envconfig.Process("BG_LOG", &Log)
	if err != nil {
		log.Fatal(err.Error())
	}
	logrus.SetLevel(logrus.Level(Log.Level))

	err = envconfig.Process("BG_FRONTEND", &Frontend)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_ARDUINO", &Arduino)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_DB", &Db)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_GPIO", &Gpio)
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
	logrus.Infof("Gpio    : %+v", Gpio)
	logrus.Infof("Watering: %+v", Watering)
}
