package constants

import (
	"errors"
	"fmt"
)

const (
	MalformedAlphaCodeError = "the alpha-code must be 3 letters"
	MalformedCovidYearError = "not a valid covid year. Must be between 2019 and 2030"
	MalformedMonthError     = "not a valid month. Must be between 01 and 12"
	MalformedDayError       = "not a valid day. Must be between 01 and 31"
	InvalidMethodError      = "not a valid http method"
	CountryNotFoundError    = "country not found"
	CountryAlreadyInCache   = "country already in cache"
	CountryNotInCache       = "country not in cache"
	ExpiredCacheEntry       = "cache entry has expired, and has been deleted"
	CountryDBIsEmpty        = "the country database is empty"

	linkToDocs = "https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022-workspace/vegarfae/assignment-2/-/blob/main/README.md"
)

func IsBadRequestError(err error) bool {
	switch err.Error() {
	case MalformedAlphaCodeError,
		MalformedCovidYearError,
		MalformedMonthError,
		MalformedDayError,
		InvalidMethodError,
		CountryNotFoundError:
		return true
	default:
		return false

	}
}

func GetCountryNotFoundInCasesApi(country string) error {
	return errors.New(fmt.Sprintf("country not found in cases-api: %s", country))
}

func getDocumentationError() error {
	return errors.New(fmt.Sprintf("See %s for documentation", linkToDocs))
}

func GetNotValidPathError() error {
	return errors.New(
		fmt.Sprintf("Not a valid endpoint. \n\n"+
			"Please use paths %s, %s, %s, or %s. \n\n"+
			"%s", PolicyPath, CasesPath, NotificationsPath, StatusPath, getDocumentationError().Error()),
	)
}

func GetBadPoliciesRequestError() error {
	return errors.New(fmt.Sprintf("Not a valid request. Format: /corona/v1/policy/{:country_name}{?scope=YYYY-MM-DD}\n"+
		"\n"+
		"%s", getDocumentationError().Error()))
}

func GetBadCasesRequestError() error {
	return errors.New(fmt.Sprintf("Not a valid request. Format: /corona/v1/cases/{:country_name}\n"+
		"\n"+
		"%s", getDocumentationError().Error()))
}
