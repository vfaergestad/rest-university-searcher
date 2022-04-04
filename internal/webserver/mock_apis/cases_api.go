package mock_apis

// cases_api is a mock API of the cases API for testing purposes.

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

// HandlerCases is the entry point of the mock api.
func HandlerCases(w http.ResponseWriter, r *http.Request) {

	// Get the query from the request.
	var query queryStruct
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		fmt.Println(err.Error())
	}

	var response structs.CasesApiResponse

	// Checks if the query contains "Taiwan", and returns an empty struct.
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
		// Responds with a standard response.
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

	// Encode the response to JSON and write it to the response.
	err = encode_struct.EncodeStruct(w, response)
	if err != nil {
		panic(err)
	}
}
