package handlers

import (
	"assignment-1/structs"
	"assignment-1/webserver/requests/country"
	"assignment-1/webserver/requests/university"
	"encoding/json"
	"net/http"
	"path"
	"strings"
)

func HandlerUniInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	universities := getUniInfo(r)

	if len(universities) == 0 {
		http.Error(w, "No universities found", http.StatusNoContent)
		return
	}

	var fields []string
	if r.URL.Query().Get("fields") != "" {
		fields = strings.Split(r.URL.Query().Get("fields"), ",")
	} else {
		fields = nil
	}

	combinedInfo := getCombined(universities, fields)

	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(combinedInfo)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}

func getUniInfo(r *http.Request) []structs.University {
	search := path.Base(r.URL.Path)
	if search == "" {
		return nil
	}

	universities, _ := university.RequestUniversity(search)

	if universities == nil {
		return nil
	}

	return universities

}

func getCombined(universities []structs.University, fields []string) []structs.UniAndCountry {
	var uniInfos []structs.UniAndCountry
	for _, u := range universities {
		c, _ := country.GetCountryByName(u.Country)
		uniInfo := structs.CombineUniCountry(u, c, fields...)
		uniInfos = append(uniInfos, uniInfo)

	}
	return uniInfos
}
