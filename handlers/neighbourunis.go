package handlers

import "net/http"

func HandlerNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

}
