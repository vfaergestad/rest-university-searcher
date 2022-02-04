package handlers

import (
	university_search "assignment-1"
	"fmt"
	"net/http"
)

func HandlerDiag(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	diagnose := getDiagnose()

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}

func getDiagnose() Diagnose {

	url := university_search.UNIVERSITIESAPI_URL
	universityApiRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating university API request: %e", err.Error())
	}
	universityApiRequest.Header.Add("content-type", "application/json")

	url = university_search.COUNTRIESAPI_URL
	countriesApiRequest, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating country API request: %e", err.Error())
	}

	client := http.Client{}
	university, err := client.Do(universityApiRequest)
	if err != nil {
		fmt.Errorf("Error in university API response: %e", err.Error())
	}

	universityApiStatus := res.StatusCode
}
