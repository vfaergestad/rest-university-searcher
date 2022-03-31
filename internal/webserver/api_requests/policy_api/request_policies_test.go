package policy_api

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/mock_apis"
	"assignment-2/internal/webserver/structs"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	policyMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerPolicy))
	defer policyMock.Close()
	constants.SetTestPolicyApiUrl(policyMock.URL)

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
		alphaCode string
		year      string
		month     string
		day       string
	}
	tests := []struct {
		name             string
		args             args
		expectedResponse structs.PolicyApiResponse
		expectedError    error
	}{
		{
			name: "Valid Request",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "01",
				day:       "01",
			},
			expectedResponse: structs.PolicyApiResponse{
				PolicyActions: []structs.PolicyAction{
					{PolicyTypeCode: "NONE"},
				},
				StringencyData: map[string]interface{}{
					"date_value":        "2022-03-26",
					"country_code":      "NOR",
					"confirmed":         1396911,
					"deaths":            2339,
					"stringency_actual": nil,
					"stringency":        13.89,
				},
			},
		},
		{
			name: "Invalid year",
			args: args{
				alphaCode: "NOR",
				year:      "2000",
				month:     "01",
				day:       "01",
			},
			expectedError: errors.New(constants.MalformedCovidYearError),
		},
		{
			name: "Invalid month",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "13",
				day:       "01",
			},
			expectedError: errors.New(constants.MalformedMonthError),
		},
		{
			name: "Invalid day",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "01",
				day:       "32",
			},
			expectedError: errors.New(constants.MalformedDayError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := getResponse(tt.args.alphaCode, tt.args.year, tt.args.month, tt.args.day)
			if err != nil {
				if tt.expectedError == nil {
					t.Errorf("Error: %v", err)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
				}
			} else if tt.expectedError != nil {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if response.StringencyData != nil {
				if response.StringencyData["date_value"] != tt.expectedResponse.StringencyData["date_value"] {
					t.Errorf("Expected date_value: %v, got: %v", tt.expectedResponse.StringencyData["date_value"], response.StringencyData["date_value"])
				}
				if response.StringencyData["country_code"] != tt.expectedResponse.StringencyData["country_code"] {
					t.Errorf("Expected country_code: %v, got: %v", tt.expectedResponse.StringencyData["country_code"], response.StringencyData["country_code"])
				}
				if int(response.StringencyData["confirmed"].(float64)) != tt.expectedResponse.StringencyData["confirmed"].(int) {
					t.Errorf("Expected confirmed: %v, got: %v", tt.expectedResponse.StringencyData["confirmed"], response.StringencyData["confirmed"])
				}
				if int(response.StringencyData["deaths"].(float64)) != tt.expectedResponse.StringencyData["deaths"].(int) {
					t.Errorf("Expected deaths: %v, got: %v", tt.expectedResponse.StringencyData["deaths"], response.StringencyData["deaths"])
				}
				if response.StringencyData["stringency_actual"] != tt.expectedResponse.StringencyData["stringency_actual"] {
					t.Errorf("Expected stringency_actual: %v, got: %v", tt.expectedResponse.StringencyData["stringency_actual"], response.StringencyData["stringency_actual"])
				}
				if response.StringencyData["stringency"] != tt.expectedResponse.StringencyData["stringency"] {
					t.Errorf("Expected stringency: %v, got: %v", tt.expectedResponse.StringencyData["stringency"], response.StringencyData["stringency"])
				}
			}
			if response.PolicyActions != nil {
				if response.PolicyActions[0].PolicyTypeCode != tt.expectedResponse.PolicyActions[0].PolicyTypeCode {
					t.Errorf("Expected policy_type_code: %v, got: %v", tt.expectedResponse.PolicyActions[0].PolicyTypeCode, response.PolicyActions[0].PolicyTypeCode)
				}
			}
		})
	}
}

func TestGetStringencyAndPolicies(t *testing.T) {
	type args struct {
		alphaCode string
		year      string
		month     string
		day       string
	}
	tests := []struct {
		name               string
		args               args
		expectedStringency float64
		expectedPolicies   int
		expectedError      error
	}{
		{
			name: "Valid request without stringency actual",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "01",
				day:       "01",
			},
			expectedStringency: 13.89,
			expectedPolicies:   0,
		},
		{
			name: "Valid request with stringency actual",
			args: args{
				alphaCode: "SWE",
				year:      "2020",
				month:     "01",
				day:       "01",
			},
			expectedStringency: 12.69,
			expectedPolicies:   0,
		},
		{
			name: "Non existing alpha code",
			args: args{
				alphaCode: "NOS",
				year:      "2020",
				month:     "01",
				day:       "01",
			},
			expectedStringency: -1,
			expectedPolicies:   -1,
			expectedError:      errors.New(constants.PoliciesDataUnavailableError),
		},
		{
			name: "Invalid alpha code",
			args: args{
				alphaCode: "123",
				year:      "2020",
				month:     "01",
				day:       "01",
			},
			expectedStringency: -1,
			expectedPolicies:   -1,
			expectedError:      errors.New(constants.MalformedAlphaCodeError),
		},
		{
			name: "Invalid year",
			args: args{
				alphaCode: "NOR",
				year:      "2000",
				month:     "01",
				day:       "01",
			},
			expectedStringency: -1,
			expectedPolicies:   -1,
			expectedError:      errors.New(constants.MalformedCovidYearError),
		},
		{
			name: "Invalid month",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "13",
				day:       "01",
			},
			expectedStringency: -1,
			expectedPolicies:   -1,
			expectedError:      errors.New(constants.MalformedMonthError),
		},
		{
			name: "Invalid day",
			args: args{
				alphaCode: "NOR",
				year:      "2020",
				month:     "01",
				day:       "32",
			},
			expectedStringency: -1,
			expectedPolicies:   -1,
			expectedError:      errors.New(constants.MalformedDayError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringency, policies, err := GetStringencyAndPolicies(tt.args.alphaCode, tt.args.year, tt.args.month, tt.args.day)
			if err != nil {
				if tt.expectedError == nil {
					t.Errorf("Error: %v", err)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
				}
			} else if tt.expectedError != nil {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if stringency != tt.expectedStringency {
				t.Errorf("Expected stringency: %v, got: %v", tt.expectedStringency, stringency)
			}
			if policies != tt.expectedPolicies {
				t.Errorf("Expected policies: %v, got: %v", tt.expectedPolicies, policies)
			}
		})
	}
}
