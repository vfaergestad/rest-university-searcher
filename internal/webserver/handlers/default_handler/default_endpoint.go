package default_handler

// default_endpoint handles all incoming requests that do not match any other endpoints.

import (
	"assignment-2/internal/webserver/constants"
	"net/http"
)

// HandlerDefault is the entry point for all incoming requests to the endpoint.
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	// Responds with a 404 Not Found error and a message with the correct endpoint paths.
	http.Error(w, constants.GetNotValidPathError().Error(), http.StatusNotFound)
}
