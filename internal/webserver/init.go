package webserver

import (
	"assignment-2/internal/webserver/handlers"
	"assignment-2/internal/webserver/uptime"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = "8080"

	// The paths that will be handled by each handler
	DEFAULT_PATH       = "/corona/"
	CASES_PATH         = "/corona/v1/cases/"
	POLICY_PATH        = "/corona/v1/policy/"
	STATUS_PATH        = "/corona/v1/status/"
	NOTIFICATIONS_PATH = "/corona/v1/notifications/"
)

func InitServer() {

	// Points the different URL-paths to the correct handler
	http.HandleFunc(DEFAULT_PATH, handlers.HandlerDefault)
	http.HandleFunc(CASES_PATH, handlers.HandlerCases)
	http.HandleFunc(POLICY_PATH, handlers.HandlerPolicy)
	http.HandleFunc(STATUS_PATH, handlers.HandlerStatus)
	http.HandleFunc(NOTIFICATIONS_PATH, handlers.HandlerNotifications)

	// Starting HTTP-server
	log.Println("Starting server on port " + DEFAULT_PORT + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+DEFAULT_PORT, nil))

}
