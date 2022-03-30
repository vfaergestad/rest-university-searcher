package webserver

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/handlers/cases_handler"
	"assignment-2/internal/webserver/handlers/default_handler"
	"assignment-2/internal/webserver/handlers/notifications_handler"
	"assignment-2/internal/webserver/handlers/policy_handler"
	"assignment-2/internal/webserver/handlers/status_handler"
	"assignment-2/internal/webserver/utility/uptime"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

func InitServer() {

	err := db.InitializeFirestore()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = db.CloseFirestore()
		if err != nil {
			panic(err)
		}
	}()

	err = country_cache.InitCache()
	if err != nil {
		panic(err)
	}

	// Points the different URL-paths to the correct handler
	http.HandleFunc(constants.DefaultPath, default_handler.HandlerDefault)
	http.HandleFunc(constants.CasesPath, cases_handler.HandlerCases)
	http.HandleFunc(constants.PolicyPath, policy_handler.HandlerPolicy)
	http.HandleFunc(constants.StatusPath, status_handler.HandlerStatus)
	http.HandleFunc(constants.NotificationsPath, notifications_handler.HandlerNotifications)

	// Starting HTTP-server
	log.Println("Starting server on port " + defaultPort + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}
