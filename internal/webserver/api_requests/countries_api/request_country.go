package countries_api

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type countryApiResponse struct {
	Name map[string]interface{} `json:"name"`
}

func GetStatusCode() (int, error) {
	res, err := utility.HeadRequest(constants.CountryAPIUrl + "all")
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

	url := fmt.Sprintf("%salpha/%s?fields=name", constants.CountryAPIUrl, alphaCode)
	res, err := utility.GetRequest(url)
	if err != nil {
		return "", err
	}

	countryStruct, err := decodeCountry(res, countryApiResponse{})
	if err != nil {
		return "", err
	}

	countryName := countryStruct.Name["common"]
	if countryName == "" || countryName == nil {
		return "", errors.New(constants.CountryNotFoundError)
	} else {
		return countryName.(string), nil
	}

}

func decodeCountry(res *http.Response, target countryApiResponse) (countryApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return countryApiResponse{}, err
	}
	return target, nil
}
