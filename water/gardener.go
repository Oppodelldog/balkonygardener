package water

import (
	"fmt"
	"github.com/Oppodelldog/balkonygardener/log"
	"time"

	"github.com/Oppodelldog/balkonygardener/db"
	"github.com/jasonlvhit/gocron"
	"github.com/mrmorphic/hwio"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func StartGardener() {
	hwio.SetDriver(hwio.NewRaspPiDTDriver())
	gocron.Every(1).Day().At("18:36").Do(watering)
	gocron.Start()
}

var wateringConfigs = []WateringConfig{
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

func watering() {

	for _, wateringConfig := range wateringConfigs {
		err := Water(wateringConfig)
		if err != nil {
			logrus.Errorf("error during watering %s(%s): %v", wateringConfig.PinName, wateringConfig.Comment, err)
		}
	}
}
func Water(config WateringConfig) error {

	flowersPin, err := hwio.GetPin(config.PinName)
	if err != nil {
		return errors.Wrap(err, "Could not GetPing")

	}
	defer log.Error(hwio.ClosePin(flowersPin))

	err = hwio.PinMode(flowersPin, hwio.OUTPUT)
	if err != nil {
		return errors.Wrap(err, "Could not PinMode")
	}

	err = hwio.DigitalWrite(flowersPin, hwio.LOW)
	if err != nil {
		return errors.Wrap(err, "Could not DigitalWrite (LOW)")
	}

	log.Error(db.SaveString("water", fmt.Sprintf("OPEN WATER PIPELINE %s (%s)", config.PinName, config.Comment)))
	time.Sleep(config.Duration)
	log.Error(db.SaveString("water", fmt.Sprintf("CLOSE WATER PIPELINE %s (%s)", config.PinName, config.Comment)))

	err = hwio.DigitalWrite(flowersPin, hwio.HIGH)
	if err != nil {
		return errors.Wrap(err, "Could not DigitalWrite (HIGH)")
	}

	return nil
}
