package policy_handler

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/db/policies_db"
	"assignment-2/internal/webserver/mock_apis"
	"assignment-2/internal/webserver/structs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var policyEndpoint *httptest.Server

func TestMain(m *testing.M) {
	policyMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerPolicy))
	defer policyMock.Close()

	constants.SetTestPolicyApiUrl(policyMock.URL)
	constants.SetTestServiceAccountLocation()

	err := db.InitializeFirestore()
	if err != nil {
		panic(err)
	}
	policies_db.SetTestMode()

	defer func() {
		err = db.CloseFirestore()
		if err != nil {
			panic(err)
		}
	}()

	policyEndpoint = httptest.NewServer(http.HandlerFunc(HandlerPolicy))
	defer policyEndpoint.Close()

	m.Run()
}

func TestRequestsFromPolicyApi(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   structs.PolicyResponse
	}{
		{
			name: "Valid Get Request With Correct Country Code And No Scope",
			args: args{
				url: policyEndpoint.URL + "/corona/v1/policy/NOR",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/SWE?scope=2021-01-01",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "SWE",
				Scope:       "2021-01-01",
				Stringency:  12.69,
				Policies:    0,
			},
		},
		{
			name: "Valid Get Request With To Long Path",
			args: args{
				url: policyEndpoint.URL + "/corona/v1/policy/NOR/SWE?scope=2021-01-01",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/NOR?scope=21-01-22",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/NORD",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/NOR?scope=2018-01-01",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/NOR?scope=2020-30-01",
			},
			method:             http.MethodGet,
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
				url: policyEndpoint.URL + "/corona/v1/policy/NOR?scope=2020-01-32",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.PolicyResponse{
				CountryCode: "",
				Scope:       "",
				Stringency:  0,
				Policies:    0,
			},
		},
		{
			name: "Invalid Post Request",
			args: args{
				url: policyEndpoint.URL + "/corona/v1/policy/NOR",
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

			if tt.expectedResponse != (structs.PolicyResponse{}) {
				var actual structs.PolicyResponse
				_ = json.NewDecoder(res.Body).Decode(&actual)

				if actual != tt.expectedResponse {
					t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
				}
			}
		})
	}
}
