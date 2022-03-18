package webserver

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/handlers"
	"assignment-2/internal/webserver/uptime"
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
	http.HandleFunc(constants.DefaultPath, handlers.HandlerDefault)
	http.HandleFunc(constants.CasesPath, handlers.HandlerCases)
	http.HandleFunc(constants.PolicyPath, handlers.HandlerPolicy)
	http.HandleFunc(constants.StatusPath, handlers.HandlerStatus)
	http.HandleFunc(constants.NotificationsPath, handlers.HandlerNotifications)

	// Starting HTTP-server
	log.Println("Starting server on port " + defaultPort + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}
