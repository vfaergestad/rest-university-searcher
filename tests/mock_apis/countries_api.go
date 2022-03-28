package mock_apis

import (
	"assignment-2/internal/webserver/utility"
	"net/http"
)

type countryApiResponse struct {
	Name map[string]interface{}
}

func HandlerCountries(w http.ResponseWriter, r *http.Request) {
	response := countryApiResponse{map[string]interface{}{
		"name": map[string]interface{}{
			"common": "Norway",
		},
	}}

	err := utility.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
