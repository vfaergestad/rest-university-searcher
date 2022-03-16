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

)

func InitServer() {

	// Points the different URL-paths to the correct handler
	http.HandleFunc(constants.DefaultPath, handlers.HandlerDefault)
	http.HandleFunc(constants.CasesPath, handlers.HandlerCases)
	http.HandleFunc(constants.PolicyPath, handlers.HandlerPolicy)
	http.HandleFunc(constants.StatusPath, handlers.HandlerStatus)
	http.HandleFunc(constants.NotificationsPath, handlers.HandlerNotifications)

	// Starting HTTP-server
	log.Println("Starting server on port " + DEFAULT_PORT + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+DEFAULT_PORT, nil))

}
