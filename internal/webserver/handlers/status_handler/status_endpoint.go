package status_handler

import (
	"assignment-2/internal/webserver/api_requests/cases_api"
	"assignment-2/internal/webserver/api_requests/countries_api"
	"assignment-2/internal/webserver/api_requests/policy_api"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/utility"
	"assignment-2/internal/webserver/utility/uptime"
	"net/http"
)

type statusResponse struct {
	CasesApi   int    `json:"cases_api"`
	PolicyApi  int    `json:"policy_api"`
	CountryApi int    `json:"country_api"`
	Webhooks   int    `json:"webhooks"`
	Version    string `json:"version"`
	Uptime     string `json:"uptime"`
}

func HandlerStatus(w http.ResponseWriter, r *http.Request) {

	// Responds with error if method is anything else than GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	policyStatus, err := policy_api.GetStatusCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countryStatus, err := countries_api.GetStatusCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	casesStatus, err := cases_api.GetStatusCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webhooks, err := webhooks_db.GetDBSize()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := statusResponse{
		CasesApi:   casesStatus,
		PolicyApi:  policyStatus,
		CountryApi: countryStatus,
		Webhooks:   webhooks,
		Version:    "v1",
		Uptime:     uptime.GetUptimeString(),
	}

	err = utility.EncodeStruct(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
