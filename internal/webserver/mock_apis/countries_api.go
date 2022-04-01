package mock_apis

import (
	"assignment-2/internal/webserver/utility/encode_struct"
	"net/http"
	"path"
)

type countryApiResponse struct {
	Name map[string]interface{}
}

func HandlerCountries(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	country := path.Base(cleanPath)

	var response countryApiResponse
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

	err := encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
