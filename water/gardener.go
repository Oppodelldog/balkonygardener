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
	gocron.Every(1).Day().At(config.Watering.Time).Do(watering)
	gocron.Start()
}

func watering() {
	for _, wateringConfig := range config.Watering.Waterings {
		err := Water(wateringConfig)
		if err != nil {
			logrus.Errorf("error during watering %s(%s): %v", wateringConfig.PinName, wateringConfig.Comment, err)
		}
	}
}
func Water(config config.WateringEntryConfig) error {

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
