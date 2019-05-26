package water

import (
	"fmt"
	"github.com/Oppodelldog/balkonygardener/config"
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
	for pinName, wateringConfig := range config.Watering {
		gocron.Every(1).Day().At(wateringConfig.Time).Do(func() {
			err := Water(pinName, wateringConfig)
			if err != nil {
				logrus.Errorf("error during watering %s(%s): %v", pinName, wateringConfig.Comment, err)
			}
		})
	}
	gocron.Start()
}

func Water(pinName string, config config.WateringEntryConfig) error {

	flowersPin, err := hwio.GetPin(pinName)
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

	log.Error(db.SaveString("water", fmt.Sprintf("OPEN WATER PIPELINE %s (%s)", pinName, config.Comment)))
	time.Sleep(config.Duration)
	log.Error(db.SaveString("water", fmt.Sprintf("CLOSE WATER PIPELINE %s (%s)", pinName, config.Comment)))

	err = hwio.DigitalWrite(flowersPin, hwio.HIGH)
	if err != nil {
		return errors.Wrap(err, "Could not DigitalWrite (HIGH)")
	}

	return nil
}
