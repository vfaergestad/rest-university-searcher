package requests

import (
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func RequestUniInfo(url string) []structs.University {
	//fmt.Println(url)
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
	var universities []structs.University
	if err := decoder.Decode(&universities); err != nil {
		log.Fatal(err)
	}

	return universities

}

func RequestCountryInfo(url string) structs.Country {
	//fmt.Println(url)

	var alphaSearch bool
	urlParts := strings.Split(url, "/")
	if urlParts[len(urlParts)-2] == "alpha" {
		alphaSearch = true
	} else {
		alphaSearch = false
	}

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating country request: %e", err.Error())
	}

	r.Header.Add("content-type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Errorf("Error in country response: %e", err.Error())
	}

	decoder := json.NewDecoder(res.Body)
	if alphaSearch {
		var country structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}

		return country
	} else {
		var country []structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}

		return country[0]
	}
}
