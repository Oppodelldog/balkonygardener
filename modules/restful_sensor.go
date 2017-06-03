package modules

import "github.com/Oppodelldog/balkonygardener/api"

type RestfulSensorApi struct{}

func (module *RestfulSensorApi )Start() {
	go api.StartRestfulApi()
}
func NewRestfulSenorApiModule() *RestfulSensorApi{
	return &RestfulSensorApi{}
}

