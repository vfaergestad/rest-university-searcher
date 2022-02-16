package handlers

import (
	"assignment-1/constants"
	"assignment-1/handlers/structs"
	"assignment-1/uptime"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerDiag(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	diagnose := requestDiagnose()

	w.Header().Add("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(diagnose)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

}

func requestDiagnose() structs.Diagnose {

	url := constants.UNIVERSITIESAPI_URL
	universityApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating university API request: %e", err.Error())
	}
	universityApiRequest.Header.Add("content-type", "application/json")

	url = constants.COUNTRIESAPI_URL + "all"
	countriesApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating country API request: %e", err.Error())
	}

	client := http.Client{}
	res, err := client.Do(universityApiRequest)
	if err != nil {
		fmt.Errorf("Error in university API response: %e", err.Error())
	}

	universityApiStatus := res.StatusCode

	res, err = client.Do(countriesApiRequest)
	if err != nil {
		fmt.Errorf("Error in countries API response: %e", err.Error())
	}

	countriesApiStatus := res.StatusCode

	return structs.Diagnose{
		UniversitiesApi: fmt.Sprintf("%d", universityApiStatus),
		CountriesApi:    fmt.Sprintf("%d", countriesApiStatus),
		Version:         "v1",
		Uptime:          uptime.GetUptime(),
	}

}
