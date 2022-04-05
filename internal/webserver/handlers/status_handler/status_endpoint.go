package status_handler

// Status_handler handles the corona/v1/status endpoint.

import (
	"assignment-2/internal/webserver/api_requests/cases_api"
	"assignment-2/internal/webserver/api_requests/countries_api"
	"assignment-2/internal/webserver/api_requests/policy_api"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/utility/encode_struct"
	"assignment-2/internal/webserver/utility/logging"
	"assignment-2/internal/webserver/utility/uptime"
	"net/http"
)

// Struct for holding information to respond with.
type statusResponse struct {
	CasesApi   int    `json:"cases_api"`
	PolicyApi  int    `json:"policy_api"`
	CountryApi int    `json:"country_api"`
	Webhooks   int    `json:"webhooks"`
	Version    string `json:"version"`
	Uptime     string `json:"uptime"`
}

// HandlerStatus is the entry point for the endpoint.
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	// Responds with error if method is anything else than GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported. Currently only GET are supported.", http.StatusMethodNotAllowed)
		return
	}

	// Retrieves the status codes from the APIs.
	policyStatus, _ := policy_api.GetStatusCode()
	countryStatus, _ := countries_api.GetStatusCode()
	casesStatus, _ := cases_api.GetStatusCode()
	webhooks, _ := webhooks_db.GetDBSize()

	// Creates a new statusResponse struct with the retrieved status codes and other information.
	response := statusResponse{
		CasesApi:   casesStatus,
		PolicyApi:  policyStatus,
		CountryApi: countryStatus,
		Webhooks:   webhooks,
		Version:    "v1",
		Uptime:     uptime.GetUptimeString(),
	}

	// Encodes the response struct to JSON, and writes it to the response.
	err := encode_struct.EncodeStruct(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
