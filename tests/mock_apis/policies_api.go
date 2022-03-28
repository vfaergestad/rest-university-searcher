package mock_apis

import (
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"net/http"
)

func HandlerPolicy(w http.ResponseWriter, r *http.Request) {
	response := structs.PolicyApiResponse{
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
	}

	err := utility.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
