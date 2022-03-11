package webserver

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/webserver/handlers"
	"net/http"
)

func InitServer() {

	// Points the different URL-paths to the correct handler
	http.HandleFunc(constants.DEFAULT_PATH, handlers.HandlerDefault)
	http.HandleFunc(constants.CASES_PATH, handlers.HandlerCases)
	http.HandleFunc(constants.POLICY_PATH, handlers.HandlerPolicy)
	http.HandleFunc(constants.STATUS_PATH, handlers.HandlerStatus)
	http.HandleFunc(constants.NOTIFICATIONS_PATH, handlers.HandlerNotifications)

}
