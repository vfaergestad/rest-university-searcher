package country

import (
	"assignment-1/constants"
	"assignment-1/structs"
	"assignment-1/webserver/requests"
	"assignment-1/webserver/requests/cache"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetCountryByName(name string) (structs.Country, error) {
	if c, err := cache.GetCountry(name); err != nil {
		return c, nil
	} else {
		url := fmt.Sprintf("%sname/%s?fields=%s", constants.COUNTRIESAPI_URL, name, constants.COUNTRIESAPI_ALL_STANDARD_FIELDS)
		if c, err = requestCountry(url); err != nil {
			return structs.Country{}, err
		}
		return c, nil
	}

}

func GetCountryByAlpha(alpha string) (structs.Country, error) {
	if c, err := cache.GetCountryByAlpha(alpha); err != nil {
		return c, nil
	} else {
		url := fmt.Sprintf("%salpha/%s?fields=%s", constants.COUNTRIESAPI_URL, alpha, constants.COUNTRIESAPI_ALL_STANDARD_FIELDS)
		if c, err = requestCountry(url); err != nil {
			return structs.Country{}, err
		}
		return c, nil
	}
}

func requestCountry(url string) (structs.Country, error) {
	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return structs.Country{}, err
	}

	var c structs.Country
	if c, err = decodeCountry(res); err != nil {
		return structs.Country{}, err
	}

	if err = cache.AddCountry(c); err != nil {
		log.Println(err)
	}

	return c, nil
}

func decodeCountry(res *http.Response) (structs.Country, error) {
	decoder := json.NewDecoder(res.Body)
	var country structs.Country
	if err := decoder.Decode(&country); err != nil {
		log.Println(err)
		return structs.Country{}, err

	}
	return country, nil
}
