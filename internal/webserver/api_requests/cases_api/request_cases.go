package cases_api

import (
	"assignment-2/internal/webserver/api_requests"
	"assignment-2/internal/webserver/structs"
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	casesApiUrl = "https://covid19-graphql.now.sh"
)

type casesApiResponse struct {
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
		Date:       responseStruct.Country.MostRecent.Date,
		Confirmed:  responseStruct.Country.MostRecent.Confirmed,
		Recovered:  responseStruct.Country.MostRecent.Recovered,
		Deaths:     responseStruct.Country.MostRecent.Deaths,
		GrowthRate: responseStruct.Country.MostRecent.GrowthRate,
	}, nil

}

func getResponse(country string) (casesApiResponse, error) {
	/*
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
	*/

	queryString := map[string]string{
		"query": fmt.Sprintf(`
				{
				country(name: %s) {
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
		`, country)}

	queryValue, err := json.Marshal(queryString)
	if err != nil {
		fmt.Println(err)
	}
	res, err := api_requests.PostRequest(casesApiUrl, bytes.NewBuffer(queryValue))
	if err != nil {
		fmt.Println(err.Error())
	}
	var casesResponse casesApiResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&casesResponse); err != nil {
		return casesApiResponse{}, err
	}
	return casesResponse, nil
}
