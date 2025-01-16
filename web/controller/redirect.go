package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func RedirectUrl(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	inputAlias, ok := vars["alias"]

	if !ok {
		http.Error(writer, "Alias is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key

	urlMap := GetUrlShortener().urlMap
	originalUrlData, found := urlMap[inputAlias]
	if !found {
		http.Error(writer, "Alias not found", http.StatusNotFound)
		return
	} else {
		lengthOfAccessTimes := len(originalUrlData.AccessTimes)
		if lengthOfAccessTimes == 10 {
			originalUrlData.AccessTimes = originalUrlData.AccessTimes[1:]
		}
		originalUrlData.AccessCount += 1
		originalUrlData.AccessTimes = append(originalUrlData.AccessTimes, time.Now().Format(time.RFC3339))
	}

	// Redirect the user to the original URL
	http.Redirect(writer, request, originalUrlData.LongUrl, http.StatusTemporaryRedirect)

}
