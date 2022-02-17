package handlers

import (
	"assignment-1/constants"
	"assignment-1/handlers/requests"
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
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
	query := "search?name=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	universities := requests.RequestUniInfo(url)
	if universities == nil {
		return nil
	}

	return universities

}

func getCombined(universities []structs.University, fields []string) []structs.UniAndCountry {
	var uniInfos []structs.UniAndCountry
	for _, u := range universities {
		url := fmt.Sprintf("%sname/%s?fields=languages,maps", constants.COUNTRIESAPI_URL, u.Country)
		c := requests.RequestCountryInfo(url)
		uniInfo := structs.CombineUniCountry(u, c, fields...)
		uniInfos = append(uniInfos, uniInfo)

	}
	return uniInfos
}
