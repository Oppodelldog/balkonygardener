package water

import (
	"fmt"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/log"

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

func getGpioAdapter() GpioAdapter {
	if config.Gpio.Mock {
		return new(GpioAdapterMock)
	}
	return new(GpioHwioAdapter)
}

func Water(pinName string, config config.WateringEntryConfig) error {
	gpioAdapter := getGpioAdapter()

	flowersPin, err := gpioAdapter.GetPin(pinName)
	if err != nil {
		return errors.Wrap(err, "Could not GetPing")
	}
	defer log.Error(gpioAdapter.ClosePin(flowersPin))

	err = gpioAdapter.PinMode(flowersPin, hwio.OUTPUT)
	if err != nil {
		return errors.Wrap(err, "Could not PinMode")
	}

	err = gpioAdapter.DigitalWrite(flowersPin, hwio.LOW)
	if err != nil {
		return errors.Wrap(err, "Could not DigitalWrite (LOW)")
	}

	log.Error(db.SaveString("water", fmt.Sprintf("OPEN WATER PIPELINE %s (%s)", pinName, config.Comment)))
	time.Sleep(config.Duration)
	log.Error(db.SaveString("water", fmt.Sprintf("CLOSE WATER PIPELINE %s (%s)", pinName, config.Comment)))

	err = gpioAdapter.DigitalWrite(flowersPin, hwio.HIGH)
	if err != nil {
		return errors.Wrap(err, "Could not DigitalWrite (HIGH)")
	}

	return nil
}
