package handlers

import (
	"assignment-1/constants"
	"assignment-1/handlers/requests"
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandlerNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	pathList := strings.Split(r.URL.Path, "/")
	countryQuery := pathList[len(pathList)-2]
	uniQuery := pathList[len(pathList)-1]

	var countries []structs.Country
	country := requests.RequestCountryInfo(fmt.Sprintf("%sname/%s?borders", constants.COUNTRIESAPI_URL, countryQuery))
	for _, bor := range country.Borders {
		url := fmt.Sprintf("%salpha/%s?fields=name,languange,maps", constants.COUNTRIESAPI_URL, bor)
		countries = append(countries, requests.RequestCountryInfo(url))
	}

	var uniInfo []structs.UniAndCountry
	for _, c := range countries {
		query := fmt.Sprintf("search?name=%s&country=%s", uniQuery, c.Name["common"].(string))
		url := constants.UNIVERSITIESAPI_URL + query
		universities := requests.RequestUniInfo(url)
		for _, u := range universities {
			uniInfo = append(uniInfo, structs.CombineUniCountry(u, c))
		}
	}
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
