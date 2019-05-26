package modules

import "github.com/Oppodelldog/balkonygardener/api"

type restfulSensorApi struct{}

func (module *restfulSensorApi) Start() {
	go api.StartRestfulApi()
}
func newRestfulSenorApiModule() *restfulSensorApi {
	return &restfulSensorApi{}
}
