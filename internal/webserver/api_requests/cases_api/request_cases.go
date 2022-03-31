package cases_api

// Request_cases handles all request for information about cases from the cases-API.
import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetStatusCode returns the status code of the cases-API response.
func GetStatusCode() (int, error) {
	res, err := getResponse("Norway")
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil
}

// GetResponseStruct returns the response from the cases-API with the given country as a struct ready for encoding.
func GetResponseStruct(country string) (structs.CasesResponse, error) {
	// Get response from cases-API with the given country
	res, err := getResponse(country)
	if err != nil {
		return structs.CasesResponse{}, err
	}

	// Decodes the response into a struct
	var responseStruct structs.CasesApiResponse
	responseStruct, err = decodeCases(res, responseStruct)
	if err != nil {
		return structs.CasesResponse{}, err
	}

	// Checks if the response includes the correct country
	if responseStruct.Data.Country.Name == "" {
		return structs.CasesResponse{}, constants.GetCountryNotFoundInCasesApi(country)
	}

	// Parses the response-struct into a struct ready for encoding
	return structs.CasesResponse{
		Country:    country,
		Date:       responseStruct.Data.Country.MostRecent.Date,
		Confirmed:  responseStruct.Data.Country.MostRecent.Confirmed,
		Recovered:  responseStruct.Data.Country.MostRecent.Recovered,
		Deaths:     responseStruct.Data.Country.MostRecent.Deaths,
		GrowthRate: responseStruct.Data.Country.MostRecent.GrowthRate,
	}, nil

}

// getResponse returns the response from the cases-API with the given country as a string.
func getResponse(country string) (*http.Response, error) {

	// Formats the query for the cases-API, including the given country
	queryString := fmt.Sprintf(`
		query {
			country(name: "%s") {
				name
				mostRecent {
					date(format: "yyyy-MM-dd")
					confirmed
					recovered
					deaths
					growthRate
				}
			}
		}
	`, country)

	// Inputs the query into a query-struct
	query := queryStruct{Query: queryString}

	// Encodes the query-struct into a json-byte-array
	result, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Sends the query to the cases-API and returns the response
	res, err := utility.PostRequest(constants.CasesApiUrl, strings.NewReader(string(result)))
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, nil

}

// decodeCases decodes the response from the cases-API into a struct.
func decodeCases(res *http.Response, target structs.CasesApiResponse) (structs.CasesApiResponse, error) {
	// Reads the response from the cases-API and converts it to a byte-array
	byteResult, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Decodes the byte-array into a struct and returns it
	err = json.Unmarshal(byteResult, &target)
	if err != nil {
		return structs.CasesApiResponse{}, err
	}
	return target, nil
}
