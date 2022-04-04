package mock_apis

// country_api is a mock API of the country API for testing purposes.

import (
	"assignment-2/internal/webserver/utility/encode_struct"
	"net/http"
	"path"
)

type countryApiResponse struct {
	Name map[string]interface{}
}

// HandlerCountries is the entry point for the mock API.
func HandlerCountries(w http.ResponseWriter, r *http.Request) {
	// Get the country name from the URL.
	cleanPath := path.Clean(r.URL.Path)
	country := path.Base(cleanPath)

	var response countryApiResponse
	// Checks which country is requested, and returns the correct response.
	if country == "NOR" {
		response = countryApiResponse{map[string]interface{}{
			"common": "Norway",
		}}
	} else if country == "SWE" {
		response = countryApiResponse{map[string]interface{}{
			"common": "Sweden",
		}}
	} else if country == "DNK" {
		response = countryApiResponse{map[string]interface{}{
			"common": "Denmark",
		}}
	}

	// Encode the response to JSON, and write it to the response.
	err := encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
