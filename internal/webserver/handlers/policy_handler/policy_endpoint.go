package policy_handler

// Policy_handler is a handler for the /corona/v1/policy endpoint.

import (
	"assignment-2/internal/webserver/cache/policy_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/encode_struct"
	"assignment-2/internal/webserver/utility/logging"
	"assignment-2/internal/webserver/webhooks"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

const (
	// validDateRegex is a regular expression for validating a date.
	validDateRegex = "^\\d{4}-\\d{2}-\\d{2}$"
)

// HandlerPolicy is the entry point for the endpoint.
func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	// Checks if the method is GET.
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
	countryQuery = strings.ToUpper(countryQuery)

	// Invokes the webhook package to see if any webhooks needs to be counted up or needs to be invoked.
	go webhooks.Invoke(countryQuery)

	// Retrieves the date from the path
	scope := r.URL.Query().Get("scope")

	var policyResponseStruct structs.PolicyResponse

	// Checks if the date is empty
	if scope != "" {
		// Check if the scope is a valid date.
		match, err := regexp.MatchString(validDateRegex, scope)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if !match {
			http.Error(w, "Invalid scope. Must be on format YYYY-MM-DD.", http.StatusBadRequest)
			return
		} else {
			// Valid scope is given, retrieves the policy from the cache.
			policyResponseStruct, err = policy_cache.GetPolicy(countryQuery, scope)
			if err != nil {
				if constants.IsBadRequestError(err) {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				} else if constants.IsNotFoundError(err) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		// No scope was given. Trying to retrieve the latest policy, by going through the 7 past days.
		timeNow := time.Now()
		scope = timeNow.Format("2006-01-02")
		var err error
		for i := 0; i < 7; i++ {
			scope = timeNow.AddDate(0, 0, -i).Format("2006-01-02")
			if policyResponseStruct, err = policy_cache.GetPolicy(countryQuery, scope); err == nil {
				break
			}
		}
		if err != nil {
			if constants.IsBadRequestError(err) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else if constants.IsNotFoundError(err) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// Encodes the policy response struct to json and writes it to the response.
	err := encode_struct.EncodeStruct(w, policyResponseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
