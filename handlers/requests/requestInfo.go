package requests

import (
	"assignment-1/handlers/structs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func RequestUniInfo(url string) ([]structs.University, error) {
	//fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []structs.University{}, err
	}

	r.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)

	switch {
	case res.StatusCode == http.StatusNotFound:
		return []structs.University{}, errors.New(fmt.Sprintf("%d University not found", res.StatusCode))
	case res.StatusCode != http.StatusOK:
		return []structs.University{}, errors.New(fmt.Sprintf("Status code returned from universityAPI: %d", res.StatusCode))
	}

	decoder := json.NewDecoder(res.Body)
	var universities []structs.University
	if err := decoder.Decode(&universities); err != nil {
		log.Println(err)
	}

	return universities, nil

}

func RequestCountryInfo(url string) (structs.Country, error) {
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
		return structs.Country{}, err
	}

	r.Header.Add("content-type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return structs.Country{}, err
	}

	switch {
	case res.StatusCode == http.StatusNotFound:
		return structs.Country{}, errors.New(fmt.Sprintf("%d Country not found", res.StatusCode))
	case res.StatusCode != http.StatusOK:
		return structs.Country{}, errors.New(fmt.Sprintf("Status code returned from countryAPI: %d", res.StatusCode))
	}

	decoder := json.NewDecoder(res.Body)
	if alphaSearch {
		var country structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}

		return country, nil
	} else {
		var country []structs.Country
		if err := decoder.Decode(&country); err != nil {
			log.Fatal(err)
		}

		return country[0], nil
	}
}
