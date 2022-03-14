package policy_api

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/webserver/structs"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func GetStatusCode() (int, error) {
	year, month, day := time.Now().Date()
	url := fmt.Sprint("%s/NOR/%d-%d-%d", constants.POLICY_API_URL, year, month, day)

	r, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil

}

func GetResponse(alphaCode string, year string, month string, day string) (structs.PolicyApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.ALPHA_CODE_REGEX, alphaCode)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	} else if !match {
		return structs.PolicyApiResponse{}, errors.New(constants.MALFORMED_ALPHACODE_ERROR)
	}

	_, err = checkDate(year, month, day)
	if err != nil {
		return structs.PolicyApiResponse{}, err
	}

}

func checkDate(year string, month string, day string) (bool, error) {
	// Checks if given year is a valid covid year. (Between 2019 and 2030)
	match, err := regexp.MatchString(constants.YEAR_REGEX, month)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_MONTH_ERROR)
	}

	// Checks if given month is a valid month.
	match, err = regexp.MatchString(constants.MONTH_REGEX, month)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_MONTH_ERROR)
	}

	// Checks if given day is a valid day.
	match, err = regexp.MatchString(constants.DAY_REGEX, day)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_DAY_ERROR)
	}

	return true, nil
}
