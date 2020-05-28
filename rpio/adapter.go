package rpio

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
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
	rpio.PinMode(rpio.Pin(pin), rpio.Output)
}

func (a *GpioHwioAdapter) High(pin Pin) {
	rpio.WritePin(rpio.Pin(pin), rpio.High)
}

func (a *GpioHwioAdapter) Low(pin Pin) {
	rpio.WritePin(rpio.Pin(pin), rpio.Low)
}

func (*GpioHwioAdapter) GetPin(pinName string) (Pin, error) {
	pin, err := getPinFromName(pinName)
	if err != nil {
		return 0, err
	}
	return Pin(pin), nil
}

// GpioAdapterMock for local development without gpio access
type GpioAdapterMock struct {
}

func (m *GpioAdapterMock) Output(pin Pin) {
	logrus.Debugf("Pin %v is output", pin)
}

func (m *GpioAdapterMock) High(pin Pin) {
	logrus.Debugf("Pin %v is high", pin)
}

func (m *GpioAdapterMock) Low(pin Pin) {
	logrus.Debugf("Pin %v is low", pin)
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
