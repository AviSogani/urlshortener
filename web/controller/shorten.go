package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type UrlResponse struct {
	ShortUrl string `json:"short_url"`
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func ShortenUrl(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var inputUrlData UrlInputData
	err := json.NewDecoder(request.Body).Decode(&inputUrlData)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	finalUrlData := newUrlData()

	finalUrlData.LongUrl = inputUrlData.LongUrl

	// set TTL
	if inputUrlData.TtlSeconds != 0 {
		finalUrlData.TtlSeconds = inputUrlData.TtlSeconds
		finalUrlData.ExpiryTime = time.Now().Add(time.Duration(inputUrlData.TtlSeconds) * time.Second)
	}

	// set alias
	var shortenedURL string
	if inputUrlData.CustomAlias == "" {
		shortKey := generateShortKey()
		finalUrlData.Alias = shortKey
	} else {
		finalUrlData.Alias = inputUrlData.CustomAlias
	}

	shortenedURL = fmt.Sprintf("http://localhost:8082/%s", finalUrlData.Alias)

	GetUrlShortener().urlMap[finalUrlData.Alias] = finalUrlData

	response := UrlResponse{
		ShortUrl: shortenedURL,
	}

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
	writer.WriteHeader(http.StatusOK)
}

func newUrlData() *UrlData {
	return &UrlData{
		Alias:       "",
		LongUrl:     "",
		TtlSeconds:  DefaultTtl,
		AccessCount: 0,
		AccessTimes: []string{},
	}
}
