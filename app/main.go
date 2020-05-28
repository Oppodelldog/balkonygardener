package main

import (
	"github.com/Oppodelldog/balkonygardener/rpio"
	"time"

	"github.com/Oppodelldog/balkonygardener/modules"
)

func main() {
	rpio.Open()
	defer rpio.Close()

	modules.StartModules()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
