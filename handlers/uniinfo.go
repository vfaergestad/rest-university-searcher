package handlers

import (
	"assignment-1/constants"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

func HandlerUniInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
	}

	requestUniInfo(r)

}

func requestUniInfo(r *http.Request) []University {

	search := path.Base(r.URL.Path)
	query := "search?name=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating University request: %e", err.Error())
	}

	r.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Errorf("Error in university response: %e", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Errorf("Status code returned from universityAPI: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)

	return

	//url := constants.UNIVERSITIESAPI_URL + "search"

}
