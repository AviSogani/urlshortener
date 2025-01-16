package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UrlAnalytics struct {
	Alias       string        `json:"alias"`
	LongUrl     string        `json:"long_url"`
	AccessCount int           `json:"access_count"`
	AccessTimes []interface{} `json:"access_times"`
}

func FetchAnalytics(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var analytics UrlAnalytics

	vars := mux.Vars(request)
	inputAlias, ok := vars["alias"]

	urlmap := GetUrlShortener().urlMap

	val, ok := urlmap[inputAlias]
	if ok {
		_ = json.NewEncoder(writer).Encode(val)
	}

	err := json.NewEncoder(writer).Encode(&analytics)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}
