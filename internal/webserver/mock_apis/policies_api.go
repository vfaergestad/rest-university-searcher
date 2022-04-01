package mock_apis

import (
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/encode_struct"
	"net/http"
	"path"
	"strings"
)

func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	country := pathList[1]
	var response structs.PolicyApiResponse
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

	err := encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
