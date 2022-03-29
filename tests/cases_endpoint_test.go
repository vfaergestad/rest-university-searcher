package tests

import (
	"assignment-2/internal/webserver/structs"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetRequestsFromCases(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedResponse   structs.CasesResponse
	}{
		{
			name: "Valid Get Request With Common Name",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/Norway",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.CasesResponse{
				Country:    "Norway",
				Date:       "2020-01-01",
				Confirmed:  1,
				Recovered:  2,
				Deaths:     3,
				GrowthRate: 4,
			},
		},
		{
			name: "Valid Get Request With Country Code",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/NOR",
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.CasesResponse{
				Country:    "Norway",
				Date:       "2020-01-01",
				Confirmed:  1,
				Recovered:  2,
				Deaths:     3,
				GrowthRate: 4,
			},
		},
		{
			name: "Invalid Get Request With To Long Path",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/NOR/Test",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.CasesResponse{
				Country:    "",
				Date:       "",
				Confirmed:  0,
				Recovered:  0,
				Deaths:     0,
				GrowthRate: 0,
			},
		},
		{
			name: "Invalid Get Request With To Numbers in Country",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/NOR123",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.CasesResponse{
				Country:    "",
				Date:       "",
				Confirmed:  0,
				Recovered:  0,
				Deaths:     0,
				GrowthRate: 0,
			},
		},
		{
			name: "Invalid Get Request With Invalid Country Code",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/Veg",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: structs.CasesResponse{
				Country:    "",
				Date:       "",
				Confirmed:  0,
				Recovered:  0,
				Deaths:     0,
				GrowthRate: 0,
			},
		},
		{
			name: "Invalid Get Request With Valid Country In But Invalid In Case API",
			args: args{
				url: CasesEndpoint.URL + "/corona/v1/cases/Taiwan",
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse: structs.CasesResponse{
				Country:    "",
				Date:       "",
				Confirmed:  0,
				Recovered:  0,
				Deaths:     0,
				GrowthRate: 0,
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

			var actual structs.CasesResponse
			_ = json.NewDecoder(res.Body).Decode(&actual)

			if actual != tt.expectedResponse {
				t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
			}
		})
	}
}

func TestInvalidPostRequestFromCases(t *testing.T) {
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
				url: CasesEndpoint.URL + "/corona/v1/cases/Norway",
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
