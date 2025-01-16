package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func UpdateUrl(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	inputAlias, _ := vars["alias"]

	var inputUrlData UrlInputData
	err := json.NewDecoder(request.Body).Decode(&inputUrlData)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	if inputUrlData.CustomAlias != "" && inputUrlData.TtlSeconds != 0 {
		urlMap := GetUrlShortener().urlMap
		val, ok := urlMap[inputAlias]
		if ok {
			if inputUrlData.TtlSeconds != 0 {
				val.TtlSeconds = inputUrlData.TtlSeconds
				val.ExpiryTime = time.Now().Add(time.Duration(inputUrlData.TtlSeconds) * time.Second)
				writer.Write([]byte("Successfully updated."))
				writer.WriteHeader(http.StatusOK)
				return
			}
			if inputUrlData.CustomAlias != "" {
				urlMap[inputUrlData.CustomAlias] = val
				deleteErr := deleteFromUrlMap(inputAlias)
				if deleteErr == nil {
					writer.Write([]byte("Successfully updated."))
					writer.WriteHeader(http.StatusOK)
					return
				} else {
					writer.Write([]byte("Alias does not exist or has expired."))
					writer.WriteHeader(http.StatusNotFound)
				}
			}
		}
	} else {
		writer.Write([]byte("Invalid request"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}
