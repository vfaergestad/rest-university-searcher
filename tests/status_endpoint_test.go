package tests

import (
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/utility/uptime"
	"encoding/json"
	"net/http"
	"testing"
)

type statusResponse struct {
	CasesApi   int    `json:"cases_api"`
	PolicyApi  int    `json:"policy_api"`
	CountryApi int    `json:"country_api"`
	Webhooks   int    `json:"webhooks"`
	Version    string `json:"version"`
	Uptime     string `json:"uptime"`
}

func TestGetRequestToStatus(t *testing.T) {
	webhooks, _ := webhooks_db.GetDBSize()
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedResponse   statusResponse
	}{
		{
			name: "Valid Get Request to Status Endpoint",
			args: args{
				url: StatusEndpoint.URL + "/corona/v1/status",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: statusResponse{
				CasesApi:   http.StatusOK,
				PolicyApi:  http.StatusOK,
				CountryApi: http.StatusOK,
				Webhooks:   webhooks,
				Version:    "v1",
				Uptime:     uptime.GetUptimeString(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			res, err := client.Get(tt.args.url)
			if err != nil {
				t.Errorf("Error making GET request: %s", err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}

			var actual statusResponse
			_ = json.NewDecoder(res.Body).Decode(&actual)
			tt.expectedResponse.Uptime = uptime.GetUptimeString()

			if actual != tt.expectedResponse {
				t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
			}
		})
	}
}

func TestPostRequestToStatus(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
	}{
		{
			name: "Valid Get Request to Status Endpoint",
			args: args{
				url: StatusEndpoint.URL + "/corona/v1/status",
			},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			res, err := client.Post(tt.args.url, "application/json", nil)
			if err != nil {
				t.Errorf("Error making POST request: %s", err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
