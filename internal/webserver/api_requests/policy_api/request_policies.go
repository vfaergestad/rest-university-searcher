package policy_api

// Request_policies handles all requests and communications to the policy API.

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/request"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// GetStatusCode returns the status code of the API.
func GetStatusCode() (int, error) {
	res, err := request.HeadRequest(constants.PolicyApiStatusUrl)
	if err != nil {
		return -1, err
	}
	return res.StatusCode, nil
}

// GetStringencyAndPolicies returns the stringency and policies in the given country for the given scope.
// The country must be given as a ISO 31661 alpha-3 code.
// The scope must be a valid date in the format YYYY-MM-DD, and must be a year when covid existed.
func GetStringencyAndPolicies(alphaCode string, year string, month string, day string) (float64, int, error) {
	// Gets the response from the API.
	policyResponse, err := getResponse(alphaCode, year, month, day)
	if err != nil {
		return -1, -1, err
	}

	// Checks if the response has any data
	if policyResponse.StringencyData["msg"] == "Data unavailable" {
		log.Println(errors.New(constants.PoliciesDataUnavailableError))
		return -1, -1, errors.New(constants.PoliciesDataUnavailableError)
	}

	// Checks if the stringency-fields holds any values, and picks the right one.
	var stringency float64
	stringencyRaw := policyResponse.StringencyData["stringency_actual"]
	if stringencyRaw == nil {
		stringency = policyResponse.StringencyData["stringency"].(float64)
	} else {
		stringency = policyResponse.StringencyData["stringency_actual"].(float64)
	}

	// Checks how many policies there are. If there are only one, it checks if it is a valid policy.
	policies := len(policyResponse.PolicyActions)
	if policies == 1 {
		if policyResponse.PolicyActions[0].PolicyTypeCode == "NONE" {
			policies = 0
		}
	}
	return stringency, policies, nil
}

// getResponse returns the response from the API as a struct.
func getResponse(alphaCode string, year string, month string, day string) (structs.PolicyApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		log.Println(err)
		return structs.PolicyApiResponse{}, err
	} else if !match {
		log.Println(constants.MalformedAlphaCodeError)
		return structs.PolicyApiResponse{}, errors.New(constants.MalformedAlphaCodeError)
	}

	// Check if given date is valid.
	_, err = checkDate(year, month, day)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

	// Create URL and request response from API
	url := fmt.Sprintf("%s%s/%s-%s-%s", constants.PolicyApiUrl, alphaCode, year, month, day)
	res, err := request.GetRequest(url)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

	// Decodes the response to a struct.
	policy, err := decodePolicy(res, structs.PolicyApiResponse{})
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

	return policy, nil

}

// checkDate checks if the given date is valid.
// The date must be a year when covid existed.
func checkDate(year string, month string, day string) (bool, error) {
	// Checks if given year is a valid covid year. (Between 2019 and 2030)
	match, err := regexp.MatchString(constants.YearRegex, year)
	if err != nil {
		log.Println(err)
		return false, err
	} else if !match {
		log.Println(constants.MalformedCovidYearError)
		return false, errors.New(constants.MalformedCovidYearError)
	}

	// Checks if given month is a valid month.
	match, err = regexp.MatchString(constants.MonthRegex, month)
	if err != nil {
		log.Println(err)
		return false, err
	} else if !match {
		log.Println(constants.MalformedMonthError)
		return false, errors.New(constants.MalformedMonthError)
	}

	// Checks if given day is a valid day.
	match, err = regexp.MatchString(constants.DayRegex, day)
	if err != nil {
		log.Println(err)
		return false, err
	} else if !match {
		log.Println(constants.MalformedDayError)
		return false, errors.New(constants.MalformedDayError)
	}

	return true, nil
}

// decodePolicy decodes the response from the API to a struct.
func decodePolicy(res *http.Response, target structs.PolicyApiResponse) (structs.PolicyApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		log.Println(err)
		return structs.PolicyApiResponse{}, err
	}
	return target, nil
}
