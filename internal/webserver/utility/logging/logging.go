package logging

import (
	"log"
	"net/http"
)

func LogRequest(request *http.Request) {
	log.Printf("%s - - %s %s %s", request.RemoteAddr, request.Method, request.URL, request.Proto)
}
