package rpio

import (
	"log"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/stianeikeland/go-rpio/v4"
)

func Open() {
	if config.Gpio.Mock {
		return
	}

	err := rpio.Open()
	if err != nil {
		log.Fatalf("unable to open pin: %#v", err)
	}
}

func Close() {
	err := rpio.Close()
	if err != nil {
		log.Fatalf("unable to open pin: %#v", err)
	}
}
