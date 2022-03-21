package cases_api

import (
	"assignment-2/internal/webserver/api_requests"
	"assignment-2/internal/webserver/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	casesApiUrl = "https://covid19-graphql.now.sh"
)

type queryStruct struct {
	Query string `json:"query"`
}

type casesApiResponse struct {
	Data data `json:"data"`
}

type data struct {
	Country countryStruct `json:"country"`
}

type countryStruct struct {
	Name       string           `json:"name"`
	MostRecent mostRecentStruct `json:"mostRecent"`
}

type mostRecentStruct struct {
	Date       string  `json:"date"`
	Confirmed  int     `json:"confirmed"`
	Recovered  int     `json:"recovered"`
	Deaths     int     `json:"deaths"`
	GrowthRate float64 `json:"growthRate"`
}

/*
func GetStatusCode() (int, error) {

}*/

func GetResponse(country string) (structs.CasesResponse, error) {
	responseStruct, err := getResponse(country)
	if err != nil {
		return structs.CasesResponse{}, err
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

func getResponse(country string) (casesApiResponse, error) {

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

	res, err := api_requests.PostRequest(casesApiUrl, strings.NewReader(string(result)))
	if err != nil {
		fmt.Println(err.Error())
	}
	byteResult, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var casesResponse casesApiResponse
	err = json.Unmarshal(byteResult, &casesResponse)
	if err != nil {
		return casesApiResponse{}, err
	}
	return casesResponse, nil
}
