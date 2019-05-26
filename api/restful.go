package api

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/balkonygardener/config"
	"net/http"
	"strconv"
	"time"

	"github.com/Oppodelldog/balkonygardener/db"
	"github.com/Oppodelldog/balkonygardener/water"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

//noinspection GoUnusedParameter
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
	writeResponse(w, jsonTableNames)
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
	writeResponse(w, jsonFloatValues)
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
	writeResponse(w, jsonFloatValue)
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
	writeResponse(w, jsonFloatValues)
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
	cfg := config.WateringEntryConfig{
		Duration: time.Second * time.Duration(duration),
	}
	err = water.Water(pinName, cfg)
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
	writeResponse(w, jsonOKValue)

}
func getPinNameByPipelineId(pipelineId string) (string, error) {
	mappings := map[string]string{"gpio4": "gpio4",
		"gpio22": "gpio22",
		"gpio17": "gpio17"}

	if pinName, ok := mappings[pipelineId]; ok {
		return pinName, nil

	}

	return "", errors.New("Invalid pipelineId")
}

func writeResponse(w http.ResponseWriter, jsonTableNames []byte) {
	_, err := fmt.Fprintln(w, string(jsonTableNames))
	if err != nil {
		logrus.Errorf("could not write response to client: %v", err)
	}
}
