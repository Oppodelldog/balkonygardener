package modules

import "github.com/Oppodelldog/balkonygardener/water"

type watering struct{}

func (api *watering) Start() {
	go water.StartGardener()
}
func newWateringModule() *watering {
	return &watering{}
}
