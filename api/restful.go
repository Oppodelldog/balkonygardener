package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Oppodelldog/balkonygardener/db"
	"encoding/json"
	"strconv"
	"github.com/Oppodelldog/balkonygardener/water"
	"time"
	"github.com/pkg/errors"
)

func StartRestfulApi() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/sensor", SensorIndex)
	router.HandleFunc("/sensor/{sensorId}", SensorShow)
	router.HandleFunc("/sensor/{sensorId}/current", SensorShowCurrent)
	router.HandleFunc("/sensor/{sensorId}/hours/{hours}", SensorShowHours)
	router.HandleFunc("/pipeline/{pipelineId}/trigger/{duration}", TriggerPipeline)

	logrus.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func SensorIndex(w http.ResponseWriter, r *http.Request) {
	tableNames, err := db.ListTables()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	jsonTableNames, err := json.Marshal(tableNames)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Fprintln(w, string(jsonTableNames))
}

func SensorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId := vars["sensorId"]
	floatValues, err := db.ListValues(sensorId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	jsonFloatValues, err := json.Marshal(floatValues)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Fprintln(w, string(jsonFloatValues))
}

func SensorShowCurrent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId := vars["sensorId"]
	floatValue, err := db.GetLatestValue(sensorId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	jsonFloatValue, err := json.Marshal(floatValue)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Fprintln(w, string(jsonFloatValue))
}

func SensorShowHours(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId := vars["sensorId"]
	hours, err := strconv.Atoi(vars["hours"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	floatValues, err := db.GetValuesForHours(sensorId, hours)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	jsonFloatValues, err := json.Marshal(floatValues)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintln(w, string(jsonFloatValues))
}

func TriggerPipeline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pipelineId := vars["pipelineId"]
	duration, err := strconv.Atoi(vars["duration"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	pinName, err := getPinNameByPipelineId(pipelineId)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	config := water.WateringConfig{
		Duration: time.Second * time.Duration(duration),
		PinName:  pinName,
	}
	err = water.Water(config)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}

	jsonOKValue, err := json.Marshal("OK")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintln(w, string(jsonOKValue))

}
func getPinNameByPipelineId(pipelineId string) (string, error) {
	mappings := map[string]string{"gpio4": "gpio4",
		"gpio22":                      "gpio22",
		"gpio17":                      "gpio17"}

	if pinName, ok := mappings[pipelineId]; ok {
		return pinName, nil

	}

	return "", errors.New("Invalid pipelineId")
}
