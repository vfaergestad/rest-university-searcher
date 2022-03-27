package handlers

import (
	"assignment-2/internal/webserver/cache/policy_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/utility"
	"assignment-2/internal/webserver/webhooks"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

const (
	validDateRegex = "^\\d{4}-\\d{2}-\\d{2}$"
)

func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	// Retrieves the search-terms from the path
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")

	// Check if the given path is valid
	if len(pathList) != 5 {
		http.Error(w, constants.GetBadPoliciesRequestError().Error(), http.StatusBadRequest)
		return
	}

	// Defines the different queries in the url
	countryQuery := path.Base(cleanPath)

	go webhooks.Invoke(countryQuery)

	scope := r.URL.Query().Get("scope")

	if scope != "" {
		// Check if the scope is a valid date.
		match, err := regexp.MatchString(validDateRegex, scope)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if !match {
			http.Error(w, "Invalid scope. Must be on format YYYY-MM-DD.", http.StatusBadRequest)
			return
		}
	} else {
		timeNow := time.Now()
		scope = timeNow.Format("2006-01-02")
	}

	policyResponseStruct, err := policy_cache.GetPolicy(countryQuery, scope)
	if err != nil {
		if constants.IsBadRequestError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err = utility.EncodeStruct(w, policyResponseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
