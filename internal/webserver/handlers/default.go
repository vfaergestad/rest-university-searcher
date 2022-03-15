package handlers

import (
	"assignment-2/internal/webserver/constants"
	"net/http"
)

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	http.Error(w, constants.GetNotValidPathError().Error(), http.StatusNotFound)
}
