package route

import (
	"edra/web/controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", controller.ShortenUrl).Methods("POST")
	router.HandleFunc("/{alias}", controller.RedirectUrl).Methods("GET")
	router.HandleFunc("/analytics/{alias}", controller.FetchAnalytics).Methods("GET")
	router.HandleFunc("/update/{alias}", controller.UpdateUrl).Methods("PUT")
	router.HandleFunc("/delete/{alias}", controller.DeleteUrl).Methods("DELETE")

	router.Handle("/", router)

	// Listen on port 8082
	if err := http.ListenAndServe(":8082", router); err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Starting server")
	}
}
