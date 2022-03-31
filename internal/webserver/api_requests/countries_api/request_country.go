package countries_api

// Request_country takes a given alpha-3 code and returns the corresponding country
import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

// Struct to hold the response from the API
type countryApiResponse struct {
	Name map[string]interface{} `json:"name"`
}

// GetStatusCode returns the status code of the API response.
func GetStatusCode() (int, error) {
	res, err := utility.HeadRequest(constants.CountryAPIUrl + "all")
	if err != nil {
		return -1, err
	}
	return res.StatusCode, nil
}

// GetCountryName takes a given alpha-3 code and returns the corresponding country
func GetCountryName(alphaCode string) (string, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return "", err
	} else if !match {
		return "", errors.New(constants.MalformedAlphaCodeError)
	}

	// Formats the url for the API request.
	url := fmt.Sprintf("%salpha/%s?fields=name", constants.CountryAPIUrl, alphaCode)

	// Makes the API request.
	res, err := utility.GetRequest(url)
	if err != nil {
		return "", err
	}

	// Decodes the response from the API into a struct.
	countryStruct, err := decodeCountry(res, countryApiResponse{})
	if err != nil {
		return "", err
	}

	// Extracts the country name from the struct.
	countryName := countryStruct.Name["common"]

	// Checks if the country name was found.
	if countryName == "" || countryName == nil {
		return "", errors.New(constants.CountryNotFoundError)
	} else {
		return countryName.(string), nil
	}

}

// decodeCountry decodes the response from the API into a struct.
func decodeCountry(res *http.Response, target countryApiResponse) (countryApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return countryApiResponse{}, err
	}
	return target, nil
}
