package controller

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// removes expired alias from map
func cleanUp() {
	urlMap := GetUrlShortener().urlMap
	if len(urlMap) > 0 {
		for key, val := range urlMap {
			expiryTime := val.ExpiryTime
			currentTime := time.Now()
			fmt.Println(expiryTime)
			fmt.Println(currentTime)
			if val.ExpiryTime.Before(time.Now()) || val.ExpiryTime.Equal(time.Now()) {
				_ = deleteFromUrlMap(key)
			}
		}
	}
}

func initializeCron() {

	// Create a new cron job instance
	c := cron.New()

	// Add the cron schedule and task to the cron job
	c.AddFunc("@every 1m", cleanUp)

	// Start the cron job scheduler
	c.Start()

}
