package countries_api

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/mock_apis"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	countriesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCountries))
	defer countriesMock.Close()
	constants.SetTestCountryAPIUrl(countriesMock.URL)

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

func TestGetCountryName(t *testing.T) {
	type args struct {
		alphaCode string
	}
	tests := []struct {
		name        string
		args        args
		countryName string
		err         error
	}{
		{
			name: "Test case NOR",
			args: args{
				alphaCode: "NOR",
			},
			countryName: "Norway",
		},
		{
			name: "Test case SWE",
			args: args{
				alphaCode: "SWE",
			},
			countryName: "Sweden",
		},
		{
			name: "Test case DNK",
			args: args{
				alphaCode: "DNK",
			},
			countryName: "Denmark",
		},
		{
			name: "Invalid alpha code",
			args: args{
				alphaCode: "NOS",
			},
			countryName: "",
			err:         errors.New(constants.CountryNotFoundError),
		},
		{
			name: "Empty alpha code",
			args: args{
				alphaCode: "",
			},
			countryName: "",
			err:         errors.New(constants.MalformedAlphaCodeError),
		},
		{
			name: "Not alpha code",
			args: args{
				alphaCode: "123",
			},
			countryName: "",
			err:         errors.New(constants.MalformedAlphaCodeError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetCountryName(tt.args.alphaCode); err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("GetCountryName() error = %v, want %v", err, tt.err)
					return
				}
			} else if got != tt.countryName {
				t.Errorf("GetCountryName() = %v, want %v", got, tt.countryName)
			}
		})
	}
}
