package water

import (
	"github.com/jasonlvhit/gocron"
	"time"
	"github.com/Oppodelldog/balkonygardener/db"
	"github.com/mrmorphic/hwio"
	"fmt"
)

func StartGardener() {
	gocron.Every(1).Day().At("18:36").Do(watering)
	gocron.Start()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type WateringConfig struct {
	Duration time.Duration
	PinName  string
	Comment  string
}

var wateringConfigs = []WateringConfig{
	{
		PinName:  "gpio4",
		Comment:  "Blumen",
		Duration: time.Second * 45,
	},
	{
		PinName:  "gpio27",
		Comment:  "Baum",
		Duration: time.Second * 26,
	},
}

func watering() {
	defer hwio.CloseAll()

	hwio.SetDriver(hwio.NewRaspPiDTDriver())

	for _, wateringConfig := range wateringConfigs {
		flowersPin, err := hwio.GetPin(wateringConfig.PinName)
		if err != nil {
			break
		}

		err = hwio.PinMode(flowersPin, hwio.OUTPUT)
		panicOnError(err)
		err = hwio.DigitalWrite(flowersPin, hwio.LOW)
		panicOnError(err)

		db.SaveString("water", fmt.Sprintf("OPEN WATER PIPELINE %s (%s)", wateringConfig.PinName, wateringConfig.Comment))
		time.Sleep(wateringConfig.Duration)
		db.SaveString("water", fmt.Sprintf("CLOSE WATER PIPELINE %s (%s)", wateringConfig.PinName, wateringConfig.Comment))

		err = hwio.DigitalWrite(flowersPin, hwio.HIGH)
		panicOnError(err)
	}
}
