package cases_api

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

const (
	casesApiUrl = "https://covid19-graphql.now.sh"
)

func GetStatusCode() (int, error) {
	res, err := getResponse("Norway")
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil
}

func GetResponseStruct(country string) (structs.CasesResponse, error) {
	res, err := getResponse(country)
	if err != nil {
		return structs.CasesResponse{}, err
	}

	var responseStruct casesApiResponse
	responseStruct, err = decodeCases(res, responseStruct)
	if err != nil {
		return structs.CasesResponse{}, err
	}

	if responseStruct.Data.Country.Name == "" {
		return structs.CasesResponse{}, constants.GetCountryNotFoundInCasesApi(country)
	}

	return structs.CasesResponse{
		Country:    country,
		Date:       responseStruct.Data.Country.MostRecent.Date,
		Confirmed:  responseStruct.Data.Country.MostRecent.Confirmed,
		Recovered:  responseStruct.Data.Country.MostRecent.Recovered,
		Deaths:     responseStruct.Data.Country.MostRecent.Deaths,
		GrowthRate: responseStruct.Data.Country.MostRecent.GrowthRate,
	}, nil

}

func getResponse(country string) (*http.Response, error) {

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

	query := queryStruct{Query: queryString}

	result, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := utility.PostRequest(casesApiUrl, strings.NewReader(string(result)))
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, nil

}

func decodeCases(res *http.Response, target casesApiResponse) (casesApiResponse, error) {
	byteResult, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(byteResult, &target)
	if err != nil {
		return casesApiResponse{}, err
	}
	return target, nil
}
