package tests

import (
	"assignment-2/internal/webserver/structs"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestGetRequestsFromPolicyApi(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedResponse   structs.PolicyResponse
	}{
		{
			name: "Valid Get Request With Correct Country Code And No Scope",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "NOR",
				Scope:       time.Now().Format("2006-01-02"),
				Stringency:  13.89,
				Policies:    0,
			},
		},
		{
			name: "Valid Get Request With Correct Country Code And Scope",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/SWE?scope=2021-01-01",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "SWE",
				Scope:       "2021-01-01",
				Stringency:  13.89,
				Policies:    0,
			},
		},
		{
			name: "Valid Get Request With To Long Path",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR/SWE?scope=2021-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Scope Format",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR?scope=21-01-22",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Country Code",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NORD",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Year In Scope",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR?scope=2018-01-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Month In Scope",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR?scope=2020-30-01",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Day In Scope",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR?scope=2020-01-32",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
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

			var actual structs.PolicyResponse
			_ = json.NewDecoder(res.Body).Decode(&actual)

			if actual != tt.expectedResponse {
				t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
			}
		})
	}
}

func TestPostRequestsFromPolicyApi(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
	}{
		{
			name: "Invalid Post Request",
			args: args{
				url: PolicyEndpoint.URL + "/corona/v1/policy/NOR",
			},
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			res, err := client.Post(tt.args.url, "application/json", nil)
			if err != nil {
				t.Errorf("Error making Post request: %s", err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
