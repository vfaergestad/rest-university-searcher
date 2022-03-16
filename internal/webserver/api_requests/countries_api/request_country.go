package countries_api

import (
	"assignment-2/internal/webserver/api_requests"
	"assignment-2/internal/webserver/constants"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const (
	countryAPIUrl = "https://restcountries.com/v3.1/"
)

type countryApiResponse struct {
	Name map[string]interface{} `json:"name"`
}

func GetStatusCode() (int, error) {
	res, err := api_requests.DoRequest(countryAPIUrl+"all", http.MethodHead)
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil
}

// GetCountryName
// Possible custom errors:
// 		- MalformedAlphaCodeError
//		- CountryNotFoundError
func GetCountryName(alphaCode string) (string, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return "", err
	} else if !match {
		return "", errors.New(constants.MalformedAlphaCodeError)
	}

	url := fmt.Sprintf("%salpha/%s?fields=name", countryAPIUrl, alphaCode)
	res, err := api_requests.DoRequest(url, http.MethodGet)
	if err != nil {
		return "", err
	}

	countryStruct, err := decodeCountry(res, countryApiResponse{})
	if err != nil {
		return "", err
	}

	countryName := countryStruct.Name["common"].(string)
	if countryName == "" {
		return "", errors.New(constants.CountryNotFoundError)
	} else {
		return countryName, nil
	}

}

func decodeCountry(res *http.Response, target countryApiResponse) (countryApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return countryApiResponse{}, err
	}
	return target, nil
}
