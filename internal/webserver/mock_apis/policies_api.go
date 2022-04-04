package mock_apis

// Policies_api is a mock API of the policies API for testing purposes.

import (
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/encode_struct"
	"net/http"
	"path"
	"strings"
)

// HandlerPolicy is the entry point for the mock API.
func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	// Gets the country code from the URL.
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	country := pathList[1]

	var response structs.PolicyApiResponse

	// Checks which country code was requested, and returns the correct response.
	if country == "NOS" {
		response = structs.PolicyApiResponse{
			PolicyActions: []structs.PolicyAction{
				{PolicyTypeCode: "NONE"},
			},
			StringencyData: map[string]interface{}{
				"msg": "Data unavailable",
			},
		}
	} else if country == "SWE" {
		response = structs.PolicyApiResponse{
			PolicyActions: []structs.PolicyAction{
				{PolicyTypeCode: "NONE"},
			},
			StringencyData: map[string]interface{}{
				"date_value":        "2022-03-26",
				"country_code":      country,
				"confirmed":         1396911,
				"deaths":            2339,
				"stringency_actual": 12.69,
				"stringency":        13.89,
			},
		}
	} else if country == "FRA" {
		response = structs.PolicyApiResponse{
			PolicyActions: []structs.PolicyAction{
				{PolicyTypeCode: "NONE"},
			},
			StringencyData: map[string]interface{}{
				"date_value":        "2022-03-26",
				"country_code":      country,
				"confirmed":         1396911,
				"deaths":            2339,
				"stringency_actual": nil,
				"stringency":        nil,
			},
		}
	} else {
		response = structs.PolicyApiResponse{
			PolicyActions: []structs.PolicyAction{
				{PolicyTypeCode: "NONE"},
			},
			StringencyData: map[string]interface{}{
				"date_value":        "2022-03-26",
				"country_code":      country,
				"confirmed":         1396911,
				"deaths":            2339,
				"stringency_actual": nil,
				"stringency":        13.89,
			},
		}
	}

	// Encodes the response to JSON, and sends it to the client.
	err := encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
