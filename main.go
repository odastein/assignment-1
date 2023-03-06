package main

import (
	"assignment-1/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	handlers.StartTime = time.Now()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("$PORT has not been set. Default port is: " + port)
	}

	http.HandleFunc(handlers.UniInfoPath, handlers.UniInfoHandler)
	http.HandleFunc(handlers.NeighbourUniPath, handlers.NeighbourUnisHandler)
	http.HandleFunc(handlers.DiagPath, handlers.DiagHandler)
	log.Println("Running on port: " + port)
	log.Fatal(http.ListenAndServe(": "+port, nil))
}
