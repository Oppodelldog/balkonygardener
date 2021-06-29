package water

import (
	"fmt"
	"sync"
	"time"

	"github.com/Oppodelldog/balkonygardener/log"

	"github.com/Oppodelldog/balkonygardener/rpio"

	"github.com/Oppodelldog/balkonygardener/config"

	"github.com/jasonlvhit/gocron"
	"github.com/pkg/errors"
)

func StartGardener() {
	for pinName, wateringConfig := range config.Watering {
		err := gocron.Every(1).Day().At(wateringConfig.Time).Do(func() {
			err := Water(pinName, wateringConfig)
			if err != nil {
				log.Errorf("error during watering %s(%s): %v", pinName, wateringConfig.Comment, err)
			}
		})

		if err != nil {
			log.Errorf("error performing watering %s(%s): %v", pinName, wateringConfig.Comment, err)
		}
	}
	gocron.Start()
}

func getGpioAdapter() rpio.GpioAdapter {
	if config.Gpio.Mock {
		return new(rpio.GpioAdapterMock)
	}
	return new(rpio.GpioHwioAdapter)
}

func IsWatering(pinName string) bool {
	if _, ok := wateringLock[pinName]; !ok {
		wateringLock[pinName] = new(boolLock)
	}

	return wateringLock[pinName].IsLocked()
}

var wateringLock = map[string]*boolLock{}

type boolLock struct {
	value bool
	m     sync.Mutex
}

func (l *boolLock) IsLocked() bool {
	l.m.Lock()
	defer l.m.Unlock()

	return l.value
}
func (l *boolLock) Lock() {
	l.m.Lock()
	defer l.m.Unlock()

	l.value = true
}
func (l *boolLock) Unlock() {
	l.m.Lock()
	defer l.m.Unlock()

	l.value = false
}

func Water(pinName string, config config.WateringEntryConfig) error {
	if IsWatering(pinName) {
		return errors.New("active job running")
	}
	wateringLock[pinName].Lock()
	defer wateringLock[pinName].Unlock()

	gpioAdapter := getGpioAdapter()

	flowersPin, err := gpioAdapter.GetPin(pinName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Could not GetPin '%s'", pinName))
	}

	gpioAdapter.Output(flowersPin)
	gpioAdapter.Low(flowersPin)

	time.Sleep(config.Duration)

	gpioAdapter.High(flowersPin)

	return nil
}
