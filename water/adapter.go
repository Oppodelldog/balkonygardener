package water

import (
	"github.com/mrmorphic/hwio"
	"github.com/sirupsen/logrus"
)

type GpioAdapter interface {
	GetPin(pinName string) (hwio.Pin, error)
	PinMode(pin hwio.Pin, mode hwio.PinIOMode) error
	DigitalWrite(pin hwio.Pin, value int) (e error)
	ClosePin(pin hwio.Pin) error
}

// GpioHwioAdapter adapts hwio library which gives access to pis pins
type GpioHwioAdapter struct {
}

func (*GpioHwioAdapter) ClosePin(pin hwio.Pin) error {
	return hwio.ClosePin(pin)
}

func (*GpioHwioAdapter) GetPin(pinName string) (hwio.Pin, error) {
	return hwio.GetPin(pinName)
}

func (*GpioHwioAdapter) PinMode(pin hwio.Pin, mode hwio.PinIOMode) error {
	return hwio.PinMode(pin, mode)
}

func (*GpioHwioAdapter) DigitalWrite(pin hwio.Pin, value int) (e error) {
	return hwio.DigitalWrite(pin, value)
}

// GpioAdapterMock for local development without gpio access
type GpioAdapterMock struct {
}

func (*GpioAdapterMock) ClosePin(hwio.Pin) error {
	return nil
}

func (*GpioAdapterMock) GetPin(string) (hwio.Pin, error) {
	return hwio.Pin(1), nil
}

func (*GpioAdapterMock) PinMode(hwio.Pin, hwio.PinIOMode) error {
	return nil
}

func (*GpioAdapterMock) DigitalWrite(pin hwio.Pin, value int) (e error) {
	logrus.Debugf("Wrote to PIN %v: %v", pin, value)
	return nil
}
