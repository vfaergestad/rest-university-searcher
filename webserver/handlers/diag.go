package handlers

import (
	"assignment-1/constants"
	"assignment-1/structs"
	"assignment-1/uptime"
	"assignment-1/webserver/requests"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerDiag(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	diagnose := requestDiagnose(w)

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(diagnose)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}

func requestDiagnose(w http.ResponseWriter) structs.Diagnose {

	url := constants.UNIVERSITIESAPI_URL

	universityApiRes, err := requests.CreateAndDoRequest(http.MethodHead, url)
	if err != nil {

	}

	url = constants.COUNTRIESAPI_URL + "all"

	countryAPIRes, err := requests.CreateAndDoRequest(http.MethodHead, url)
	if err != nil {

	}

	universityApiStatus := universityApiRes.StatusCode
	countriesApiStatus := countryAPIRes.StatusCode

	return structs.Diagnose{
		UniversitiesApi: fmt.Sprintf("%d", universityApiStatus),
		CountriesApi:    fmt.Sprintf("%d", countriesApiStatus),
		Version:         "v1",
		Uptime:          uptime.GetUptime(),
	}

}
