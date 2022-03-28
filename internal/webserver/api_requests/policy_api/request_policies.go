package policy_api

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

func GetStatusCode() (int, error) {

	res, err := utility.HeadRequest(constants.PolicyApiStatusUrl)
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

	if policyResponse.StringencyData["msg"] == "Data unavailable" {
		return -1, -1, errors.New(constants.PoliciesDataUnavailableError)
	}

	var stringency float64
	stringencyRaw := policyResponse.StringencyData["stringency_actual"]
	if stringencyRaw == nil {
		stringency = policyResponse.StringencyData["stringency"].(float64)
		if stringency == 0 {
			stringency = -1
		}
	} else {
		stringency = policyResponse.StringencyData["stringency_actual"].(float64)
	}

	policies := len(policyResponse.PolicyActions)
	if policies == 1 {
		if policyResponse.PolicyActions[0].PolicyTypeCode == "NONE" {
			policies = 0
		}
	}
	return stringency, policies, nil
}

// GetResponse
// Possible custom-error-messages:
// 					- MALFORMED_ALPHACODE_ERROR
//					- MALFORMED_COVID_YEAR_ERROR
//					- MALFORMED_MONTH_ERROR
//					- MALFORMED_DAY_ERROR
func getResponse(alphaCode string, year string, month string, day string) (structs.PolicyApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	} else if !match {
		return structs.PolicyApiResponse{}, errors.New(constants.MalformedAlphaCodeError)
	}

	// Check if given date is valid.
	_, err = checkDate(year, month, day)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

	// Create URL and request response from API
	url := fmt.Sprintf("%s%s/%s-%s-%s", constants.PolicyApiUrl, alphaCode, year, month, day)
	res, err := utility.GetRequest(url)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

	policy, err := decodePolicy(res, structs.PolicyApiResponse{})
	if err != nil {
		return structs.PolicyApiResponse{}, err
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

func decodePolicy(res *http.Response, target structs.PolicyApiResponse) (structs.PolicyApiResponse, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return structs.PolicyApiResponse{}, err
	}
	return target, nil
}
