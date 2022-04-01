package mock_apis

import (
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/encode_struct"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type queryStruct struct {
	Query string `json:"query"`
}

func HandlerCases(w http.ResponseWriter, r *http.Request) {

	var query queryStruct
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		fmt.Println(err.Error())
	}

	var response structs.CasesApiResponse
	if strings.Contains(query.Query, "Taiwan") {
		response = structs.CasesApiResponse{Data: structs.Data{
			Country: structs.CountryStruct{
				Name: "",
				MostRecent: structs.MostRecentStruct{
					Date:       "",
					Confirmed:  0,
					Recovered:  0,
					Deaths:     0,
					GrowthRate: 0,
				},
			}}}
	} else {
		response = structs.CasesApiResponse{Data: structs.Data{
			Country: structs.CountryStruct{
				Name: "Norway",
				MostRecent: structs.MostRecentStruct{
					Date:       "2020-01-01",
					Confirmed:  1,
					Recovered:  2,
					Deaths:     3,
					GrowthRate: 4,
				},
			}}}

	}

	err = encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
