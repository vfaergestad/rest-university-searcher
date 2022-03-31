package cases_api

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/mock_apis"
	"assignment-2/internal/webserver/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	casesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCases))
	defer casesMock.Close()

	constants.SetTestCasesApiUrl(casesMock.URL)

	m.Run()
}

func TestGetStatusCode(t *testing.T) {
	statusCode, err := GetStatusCode()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code: %v, got: %v", http.StatusOK, statusCode)
	}
}

func TestGetResponse(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "Valid request",
			args: args{
				country: "Norway",
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getResponse(tt.args.country)
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}

func TestGetResponseStruct(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name             string
		args             args
		expectedResponse structs.CasesResponse
		expectedError    error
	}{
		{
			name: "Valid request",
			args: args{
				country: "Norway",
			},
			expectedResponse: structs.CasesResponse{
				Country:    "Norway",
				Date:       "2020-01-01",
				Confirmed:  1,
				Recovered:  2,
				Deaths:     3,
				GrowthRate: 4,
			},
			expectedError: nil,
		},
		{
			name: "Request with Invalid country",
			args: args{
				country: "Taiwan",
			},
			expectedResponse: structs.CasesResponse{},
			expectedError:    constants.GetCountryNotFoundInCasesApi("Taiwan"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := GetResponseStruct(tt.args.country)
			if err != tt.expectedError {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
				}
			}
			if response != tt.expectedResponse {
				t.Errorf("Expected response: %v, got: %v", tt.expectedResponse, response)
			}
		})
	}
}
