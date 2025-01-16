package controller

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteUrl(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	inputAlias, ok := vars["alias"]

	if ok {
		err := deleteFromUrlMap(inputAlias)
		if err == nil {
			writer.Write([]byte("Successfully deleted."))
			writer.WriteHeader(http.StatusOK)
			return
		} else {
			writer.Write([]byte("Alias does not exist or has expired."))
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}

func deleteFromUrlMap(key string) error {
	urlMap := GetUrlShortener().urlMap
	_, ok := urlMap[key]
	if ok {
		delete(GetUrlShortener().urlMap, key)
	}
	return errors.New("alias does not exist or has expired")
}
