package tests

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/handlers"
	"assignment-2/tests/mock_apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

var CasesEndpoint *httptest.Server
var PolicyEndpoint *httptest.Server

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

	err = country_cache.InitCache()
	if err != nil {
		panic(err)
	}

	CasesEndpoint = httptest.NewServer(http.HandlerFunc(handlers.HandlerCases))
	defer CasesEndpoint.Close()

	PolicyEndpoint = httptest.NewServer(http.HandlerFunc(handlers.HandlerPolicy))
	defer PolicyEndpoint.Close()

	m.Run()

}
