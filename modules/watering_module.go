package modules

import "github.com/Oppodelldog/balkonygardener/water"

type WateringModule struct{}

func (api *WateringModule) Start() {
	go water.StartGardener()
}
func NewWateringModule() *WateringModule {
	return &WateringModule{}
}
