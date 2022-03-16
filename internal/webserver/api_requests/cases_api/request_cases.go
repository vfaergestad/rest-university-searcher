package cases_api

import (
	"assignment-2/internal/webserver/constants"
	"errors"
	"regexp"
	"time"
)

type casesApiResponse struct {
	Country countryStruct `json:"country"`
}

type countryStruct struct {
	Name       string           `json:"name"`
	MostRecent mostRecentStruct `json:"mostRecent"`
}

type mostRecentStruct struct {
	Date       time.Time `json:"date"`
	Confirmed  int       `json:"confirmed"`
	Recovered  int       `json:"recovered"`
	Deaths     int       `json:"deaths"`
	GrowthRate float64   `json:"growthRate"`
}

// getResponse
// Possible custom-errors:
//		- MalformedAlphaCodeError
func getResponseWithAlpha(alphaCode string) (casesApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return casesApiResponse{}, err
	} else if !match {
		return casesApiResponse{}, errors.New(constants.MalformedAlphacodeError)
	}

}

func getResponse(country string) (casesApiResponse, error) {

}

func getBody(country string) {

}
