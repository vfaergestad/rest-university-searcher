package cases_api

import (
	"assignment-2/internal/webserver/structs"
	"context"
	"fmt"
	"github.com/machinebox/graphql"
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
	client := graphql.NewClient(casesApiUrl)

	req := graphql.NewRequest(`
		query ($country: String!) {
			country(name: $country) {
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
	`)

	req.Var("country", country)

	ctx := context.Background()

	// run it and capture the response
	var casesResponse casesApiResponse
	if err := client.Run(ctx, req, &casesResponse); err != nil {
		fmt.Println(err.Error())
		return casesApiResponse{}, err
	}
	return casesResponse, nil
}
