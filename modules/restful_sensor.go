package modules

import "github.com/Oppodelldog/balkonygardener/api"

type restfulSensorApi struct{}

func (module *restfulSensorApi) Start() {
	go api.StartAPIServer()
}
func newRestfulSenorApiModule() *restfulSensorApi {
	return &restfulSensorApi{}
}
