package policy_api

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
	policyApiStatusUrl = "https://covidtrackerapi.bsg.ox.ac.uk/api/"
	policyApiUrl       = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"
)

type policyApiResponse struct {
	PolicyActions  []interface{}          `json:"policyActions"`
	StringencyData map[string]interface{} `json:"stringencyData"`
}

func GetStatusCode() (int, error) {

	res, err := api_requests.DoRequest(policyApiStatusUrl, http.MethodHead)
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil

}

func GetStringencyAndPolicies(alphaCode string, year string, month string, day string) (float64, int, error) {
	policyResponse, err := getResponse(alphaCode, year, month, day)
	if err != nil {
		return -1, -1, err
	}

	var stringency float64
	stringencyRaw := policyResponse.StringencyData["stringency_actual"]
	if stringencyRaw != nil {
		stringency = policyResponse.StringencyData["stringency_actual"].(float64)
		if stringency == 0 {
			stringency = policyResponse.StringencyData["stringency"].(float64)
		}
	} else {
		stringency = -1
	}

	policies := len(policyResponse.PolicyActions)
	if policies < 2 {
		policies = 0
	}
	return stringency, policies, nil
}

// GetResponse
// Possible custom-error-messages:
// 					- MALFORMED_ALPHACODE_ERROR
//					- MALFORMED_COVID_YEAR_ERROR
//					- MALFORMED_MONTH_ERROR
//					- MALFORMED_DAY_ERROR
func getResponse(alphaCode string, year string, month string, day string) (policyApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return policyApiResponse{}, err
	} else if !match {
		return policyApiResponse{}, errors.New(constants.MalformedAlphacodeError)
	}

	// Check if given date is valid.
	_, err = checkDate(year, month, day)
	if err != nil {
		return policyApiResponse{}, err
	}

	// Create URL and request response from API
	url := fmt.Sprintf("%s%s/%s-%s-%s", policyApiUrl, alphaCode, year, month, day)
	res, err := api_requests.DoRequest(url, http.MethodGet)
	if err != nil {
		return policyApiResponse{}, err
	}

	policy, err := decodePolicy(res, policyApiResponse{})
	if err != nil {
		return policyApiResponse{}, err
	}

	return policy, nil

}

func checkDate(year string, month string, day string) (bool, error) {
	// Checks if given year is a valid covid year. (Between 2019 and 2030)
	match, err := regexp.MatchString(constants.YearRegex, year)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MalformedCovidYearError)
	}

	// Checks if given month is a valid month.
	match, err = regexp.MatchString(constants.MonthRegex, month)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MalformedMonthError)
	}

	// Checks if given day is a valid day.
	match, err = regexp.MatchString(constants.DayRegex, day)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MalformedDayError)
	}

	return true, nil
}

func decodePolicy(res *http.Response, target policyApiResponse) (policyApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return policyApiResponse{}, err
	}
	return target, nil
}
