package handlers

import (
	"assignment-1/constants"
	"assignment-1/handlers/requests"
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandlerNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	pathList := strings.Split(r.URL.Path, "/")
	countryQuery := pathList[len(pathList)-2]
	uniQuery := pathList[len(pathList)-1]
	var limit int
	if r.URL.Query()["limit"] != nil {
		if l, err := strconv.Atoi(r.URL.Query()["limit"][0]); err != nil || l < 0 {
			http.Error(w, "The limit is not a valid positive integer. Using 0 as limit.", http.StatusBadRequest)
			limit = 0
		} else {
			limit = l
		}
	} else {
		limit = 0
	}

	var countries []structs.Country
	country := requests.RequestCountryInfo(fmt.Sprintf("%sname/%s?borders", constants.COUNTRIESAPI_URL, countryQuery))
	for _, bor := range country.Borders {
		url := fmt.Sprintf("%salpha/%s?fields=name,languages,maps", constants.COUNTRIESAPI_URL, bor)
		countries = append(countries, requests.RequestCountryInfo(url))
	}

	var uniInfo []structs.UniAndCountry
	for _, c := range countries {
		query := fmt.Sprintf("search?name=%s&country=%s", uniQuery, c.Name["common"].(string))
		url := constants.UNIVERSITIESAPI_URL + query
		universities := requests.RequestUniInfo(url)
		for _, u := range universities {
			if limit == 0 || len(uniInfo) < limit {
				uniInfo = append(uniInfo, structs.CombineUniCountry(u, c))
			} else {
				break
			}
		}
		if len(uniInfo) >= limit {
			break
		}
	}
	if len(uniInfo) == 0 {
		http.Error(w, "No universities found", http.StatusNoContent)
		return
	}
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(uniInfo)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}
