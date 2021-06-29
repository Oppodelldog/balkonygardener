package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/Oppodelldog/balkonygardener/log"

	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
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

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = envconfig.Process("BG_LOG", &Log)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.SetLogLevel(Log.Level)

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
			log.Errorf("could not read watering config: %v", err)
		} else {
			err := yaml.Unmarshal(data, &Watering)
			if err != nil {
				log.Errorf("could parse watering config: %v", err)
			}
		}
	}
	log.Info("Config")
	log.Infof("Arduino : %+v", Arduino)
	log.Infof("Db      : %+v", Db)
	log.Infof("Gpio    : %+v", Gpio)
	log.Infof("Watering: %+v", Watering)
}
