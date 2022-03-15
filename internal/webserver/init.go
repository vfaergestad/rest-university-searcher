package webserver

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/handlers"
	"assignment-2/internal/webserver/uptime"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = "8080"

	// The paths that will be handled by each handler
	DEFAULT_PATH = "/corona/"
)

func InitServer() {

	// Points the different URL-paths to the correct handler
	http.HandleFunc(DEFAULT_PATH, handlers.HandlerDefault)
	http.HandleFunc(constants.CASES_PATH, handlers.HandlerCases)
	http.HandleFunc(constants.POLICY_PATH, handlers.HandlerPolicy)
	http.HandleFunc(constants.STATUS_PATH, handlers.HandlerStatus)
	http.HandleFunc(constants.NOTIFICATIONS_PATH, handlers.HandlerNotifications)

	// Starting HTTP-server
	log.Println("Starting server on port " + DEFAULT_PORT + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+DEFAULT_PORT, nil))

}
