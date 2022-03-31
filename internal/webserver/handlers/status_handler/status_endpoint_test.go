package status_handler

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/mock_apis"
	"assignment-2/internal/webserver/utility/uptime"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var statusEndpoint *httptest.Server

func TestMain(m *testing.M) {
	policyMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerPolicy))
	defer policyMock.Close()
	casesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCases))
	defer casesMock.Close()
	countriesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCountries))
	defer countriesMock.Close()

	constants.SetTestPolicyApiUrl(policyMock.URL)
	constants.SetTestCasesApiUrl(casesMock.URL)
	constants.SetTestCountryAPIUrl(countriesMock.URL)
	constants.SetTestServiceAccountLocation()

	err := db.InitializeFirestore()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = db.CloseFirestore()
		if err != nil {
			panic(err)
		}
	}()

	statusEndpoint = httptest.NewServer(http.HandlerFunc(HandlerStatus))
	defer statusEndpoint.Close()

	uptime.Init()

	m.Run()
}

func TestRequestToStatus(t *testing.T) {
	webhooks, _ := webhooks_db.GetDBSize()
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   statusResponse
	}{
		{
			name: "Valid Get Request to Status Endpoint",
			args: args{
				url: statusEndpoint.URL + "/corona/v1/status",
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
		{
			name: "Valid Post Request to Status Endpoint",
			args: args{
				url: statusEndpoint.URL + "/corona/v1/status",
			},
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Errorf("Error making %s request: %s", tt.method, err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}

			if tt.expectedResponse != (statusResponse{}) {
				tt.expectedResponse.Uptime = uptime.GetUptimeString()
				var actual statusResponse
				_ = json.NewDecoder(res.Body).Decode(&actual)

				if actual != tt.expectedResponse {
					t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
				}
			}
		})
	}
}
