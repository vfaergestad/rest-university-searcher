package handlers

import (
	"assignment-2/internal/webserver/api_requests/cases_api"
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/json_utility"
	"net/http"
	"path"
	"regexp"
	"strings"
)

func HandlerCases(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	// Retrieves the search-terms from the path
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")

	// Check if the given path is valid
	if len(pathList) != 5 {
		http.Error(w, constants.GetBadCasesRequestError().Error(), http.StatusBadRequest)
		return
	}

	// Defines the different queries in the url
	countryQuery := path.Base(cleanPath)

	// Check if country contains any numbers
	match, err := regexp.MatchString(constants.NoNumbersRegex, countryQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !match {
		http.Error(w, "country cannot contain numbers", http.StatusBadRequest)
		return
	}

	countryQuery = strings.Title(countryQuery)

	match, err := regexp.MatchString(constants.AlphaCodeRegex, countryQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if match {
		country, err = getCountryName(country)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	// TODO: Improve error-handling
	casesResponseStruct, err := cases_api.GetResponse(countryQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json_utility.EncodeStruct(w, casesResponseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//TODO: FIX
func getCountryName(alphaCode string) (string, error) {
	countryName, err := country_cache.GetCountry(alphaCode)
	if err != nil {
		switch err.Error() {
		case constants.ExpiredCacheEntry:

		}
	}
}
