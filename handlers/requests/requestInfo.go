package requests

import (
	"assignment-1/constants"
	"assignment-1/handlers/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

func GetUniCountryInfo(r *http.Request) []structs.UniAndCountry {
	search := path.Base(r.URL.Path)
	if search == "" {
		return nil
	}
	query := "search?name=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	universities := RequestUniInfo(url)
	if universities == nil {
		return nil
	}

	return getCombined(universities)

}

func getCombined(universities []structs.University) []structs.UniAndCountry {
	var uniInfos []structs.UniAndCountry
	for _, u := range universities {
		url := fmt.Sprintf("%sname/%s?fields=languages,maps", constants.COUNTRIESAPI_URL, u.Country)
		c := RequestCountryInfo(url)
		uniInfo := structs.UniAndCountry{
			Name:      u.Name,
			Country:   u.Country,
			Isocode:   u.AlphaTwoCode,
			Webpages:  u.WebPages,
			Languages: c.Languages,
			Map:       c.Maps["openStreetMaps"],
		}
		uniInfos = append(uniInfos, uniInfo)

	}
	return uniInfos
}

func RequestUniInfo(url string) []structs.University {
	fmt.Println(url)
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
	fmt.Println(url)

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
