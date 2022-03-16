package handlers

import (
	"assignment-2/internal/webserver/api_requests/policy_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/json_utility"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

const (
	validDateRegex = "^\\d{4}-\\d{2}-\\d{2}$"
)

type policyResponse struct {
	CountryCode string  `json:"country_code"`
	Scope       string  `json:"scope"`
	Stringency  float64 `json:"stringency"`
	Policies    int     `json:"policies"`
}

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

	scopeSplit := strings.Split(scope, "-")
	year := scopeSplit[0]
	month := scopeSplit[1]
	day := scopeSplit[2]

	stringency, polices, err := policy_api.GetStringencyAndPolicies(countryQuery, year, month, day)
	if err != nil {
		switch err.Error() {
		case constants.MalformedAlphaCodeError,
			constants.MalformedCovidYearError,
			constants.MalformedMonthError,
			constants.MalformedDayError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
	}

	policyResponseStruct := policyResponse{
		CountryCode: strings.ToTitle(countryQuery),
		Scope:       scope,
		Stringency:  stringency,
		Policies:    polices,
	}

	err = json_utility.EncodeStruct(w, policyResponseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
