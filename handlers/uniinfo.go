package handlers

import (
	"assignment-1/constants"
	"assignment-1/handlers/requests"
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

func HandlerUniInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	uniInfo := getUniCountryInfo(r)
	if uniInfo == nil {
		http.Error(w, "No universities found", http.StatusNotFound)
	}
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(uniInfo)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}

func getUniCountryInfo(r *http.Request) []structs.UniAndCountry {
	search := path.Base(r.URL.Path)
	if search == "" {
		return nil
	}
	query := "search?name=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	universities, _ := requests.RequestUniInfo(url)
	if universities == nil {
		return nil
	}

	return getCombined(universities)

}

func getCombined(universities []structs.University) []structs.UniAndCountry {
	var uniInfos []structs.UniAndCountry
	for _, u := range universities {
		url := fmt.Sprintf("%sname/%s?fields=languages,maps", constants.COUNTRIESAPI_URL, u.Country)
		c, _ := requests.RequestCountryInfo(url)
		uniInfo := structs.CombineUniCountry(u, c)
		uniInfos = append(uniInfos, uniInfo)

	}
	return uniInfos
}
