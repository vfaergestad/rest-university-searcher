package main

import (
	"assignment-1/constants"
	"assignment-1/handlers"
	"assignment-1/uptime"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = constants.PORT
	}

	http.HandleFunc(constants.UNIINFO_LOC, handlers.HandlerUniInfo)
	http.HandleFunc(constants.NEIGHBOURUNIS_LOC, handlers.HandlerNeighbourUnis)
	http.HandleFunc(constants.DIAG_LOC, handlers.HandlerDiag)

	// Starting HTTP-server
	log.Println("Starting server on port " + port + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
