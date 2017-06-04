package main

import (
	"time"
	"github.com/Oppodelldog/balkonygardener/modules"
)

func main() {

	modules.StartModules()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

