package api

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/balkonygardener/log"
	"net/http"
	"strconv"
	"time"

	"github.com/Oppodelldog/balkonygardener/config"

	"github.com/Oppodelldog/balkonygardener/water"
	"github.com/gorilla/mux"
)

func StartRestfulApi() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/pipeline/{pipelineId}/trigger/{duration}", TriggerPipeline)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, config.Frontend.IndexFile)
}

func TriggerPipeline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pipelineId := vars["pipelineId"]
	duration, err := strconv.Atoi(vars["duration"])
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	pinName := pipelineId
	cfg := config.WateringEntryConfig{
		Duration: time.Second * time.Duration(duration),
	}

	if water.IsWatering(pinName) {
		w.WriteHeader(404)
		return
	}

	go func() {
		err := water.Water(pinName, cfg)
		if err != nil {
			log.Error(err)
			w.WriteHeader(500)
			return
		}
	}()

	jsonOKValue, err := json.Marshal("OK")
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	writeResponse(w, jsonOKValue)
}

func writeResponse(w http.ResponseWriter, jsonTableNames []byte) {
	_, err := fmt.Fprintln(w, string(jsonTableNames))
	if err != nil {
		log.Errorf("could not write response to client: %v", err)
	}
}
