package handlers

import (
	"net/http"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {

	// Responds with error if method is anything else than GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

}
