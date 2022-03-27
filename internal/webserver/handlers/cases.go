package handlers

import (
	"assignment-2/internal/webserver/api_requests/cases_api"
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/utility"
	"assignment-2/internal/webserver/webhooks"
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

	match, err = regexp.MatchString(constants.AlphaCodeRegex, countryQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if match {
		countryQuery, err = country_cache.GetCountry(countryQuery)
		if err != nil {
			switch err.Error() {
			case constants.CountryNotFoundError:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	go webhooks.Invoke(countryQuery)

	casesResponseStruct, err := cases_api.GetResponseStruct(countryQuery)
	if err != nil {
		if constants.IsBadRequestError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = utility.EncodeStruct(w, casesResponseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
