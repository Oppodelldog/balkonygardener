package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"log"
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

type WateringConfig struct {
	Waterings []WateringEntryConfig
	Time      string
}

type WateringEntryConfig struct {
	Duration time.Duration
	PinName  string
	Comment  string
}

func init() {
	Watering.Waterings = []WateringEntryConfig{
		{
			PinName:  "gpio4",
			Comment:  "Blumen",
			Duration: time.Second * 35,
		},
		{
			PinName:  "gpio17",
			Comment:  "Baum",
			Duration: time.Second * 20,
		},
		{
			PinName:  "gpio22",
			Comment:  "Rote Blume",
			Duration: time.Second * 10,
		},
	}
	err := envconfig.Process("BG_ARDUINO", &Arduino)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_DB", &Db)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("BG_WATERING", &Watering)
	if err != nil {
		log.Fatal(err.Error())
	}
	logrus.Info("Config")
	logrus.Infof("Arduino : %+v", Arduino)
	logrus.Infof("Db      : %+v", Db)
	logrus.Infof("Watering: %+v", Watering)
}
