package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/Oppodelldog/balkonygardener/db"
	"encoding/json"
)

func Start() {
	go startRestfulApi()
}
func startRestfulApi() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/sensor", SensorIndex)
	router.HandleFunc("/sensor/{sensorId}", SensorShow)

	logrus.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Balkony Gardener!")
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
	floatValues, err := db.ListFloatValues(sensorId)
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
