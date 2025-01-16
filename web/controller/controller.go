package controller

import "time"

type URLShortener struct {
	//mu     sync.Mutex
	urlMap map[string]*UrlData
}

type UrlInputData struct {
	LongUrl     string `json:"long_url,omitempty"`
	CustomAlias string `json:"custom_alias,omitempty"` // todo: make a new struct
	TtlSeconds  int    `json:"ttl_seconds,omitempty"`
}

type UrlData struct {
	Alias       string    `json:"alias"`
	LongUrl     string    `json:"long_url"`
	TtlSeconds  int       `json:"ttl_seconds,omitempty"`
	AccessCount int       `json:"access_count"`
	AccessTimes []string  `json:"access_times"`
	ExpiryTime  time.Time `json:"expiry_time"`
}

var shortener *URLShortener

func GetUrlShortener() *URLShortener {
	if shortener.urlMap != nil {
		return shortener
	} else {
		return &URLShortener{
			urlMap: make(map[string]*UrlData),
		}
	}
}

func initializeUrlMap() {
	shortener = &URLShortener{
		urlMap: make(map[string]*UrlData),
	}
}

func Init() {
	initializeUrlMap()
	initializeCron()
}
