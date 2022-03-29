package mock_apis

import (
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"net/http"
	"path"
)

func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	country := path.Base(cleanPath)

	response := structs.PolicyApiResponse{
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

	err := utility.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
