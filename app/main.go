package main

import (
	"time"

	"github.com/Oppodelldog/balkonygardener/config"
	"github.com/Oppodelldog/balkonygardener/rpio"

	"github.com/Oppodelldog/balkonygardener/modules"
)

func main() {
	config.Init()
	rpio.Open()
	defer rpio.Close()

	modules.StartModules()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
