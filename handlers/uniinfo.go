package handlers

import (
	"assignment-1/handlers/requests"
	"encoding/json"
	"net/http"
)

func HandlerUniInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	uniInfo := requests.GetUniCountryInfo(r)
	if uniInfo == nil {
		http.Error(w, "No universities found", http.StatusNoContent)
	}
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(uniInfo)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}
