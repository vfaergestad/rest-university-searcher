package main

import (
	universitysearch "assignment-1"
	"assignment-1/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = universitysearch.PORT
	}

	http.HandleFunc(universitysearch.UNIINFO_LOC, handlers.HandlerUniInfo)
	http.HandleFunc(universitysearch.NEIGHBOURUNIS_LOC, handlers.HandlerNeighbourUnis)
	http.HandleFunc(universitysearch.DIAG_LOC, handlers.HandlerDiag)

	// Starting HTTP-server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
