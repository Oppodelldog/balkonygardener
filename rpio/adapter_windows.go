package rpio

import (
	"strconv"

	"github.com/Oppodelldog/balkonygardener/log"
)

type Pin int

type GpioAdapter interface {
	GetPin(pinName string) (Pin, error)
	Output(pin Pin)
	High(pin Pin)
	Low(pin Pin)
}

// GpioHwioAdapter provides methods to control gpio pins
type GpioHwioAdapter struct {
}

func (a *GpioHwioAdapter) Output(pin Pin) {

}

func (a *GpioHwioAdapter) High(pin Pin) {

}

func (a *GpioHwioAdapter) Low(pin Pin) {

}

func (*GpioHwioAdapter) GetPin(pinName string) (Pin, error) {
	return Pin(0), nil
}

// GpioAdapterMock for local development without gpio access
type GpioAdapterMock struct {
}

func (m *GpioAdapterMock) Output(pin Pin) {
	log.Debugf("Pin %v is output", pin)
}

func (m *GpioAdapterMock) High(pin Pin) {
	log.Debugf("Pin %v is high", pin)
}

func (m *GpioAdapterMock) Low(pin Pin) {
	log.Debugf("Pin %v is low", pin)
}

func (*GpioAdapterMock) ClosePin(Pin) error {
	return nil
}

func (*GpioAdapterMock) GetPin(pinName string) (Pin, error) {
	pin, err := getPinFromName(pinName)
	if err != nil {
		return 0, err
	}
	return Pin(pin), nil
}

func getPinFromName(pinName string) (int, error) {
	pinValue := pinName[len("gpio"):]
	pin, err := strconv.ParseInt(pinValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(pin), nil
}
