package country_cache

import (
	"assignment-2/internal/webserver/api_requests/countries_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/countries_db"
	"assignment-2/internal/webserver/structs"
	"errors"
	"log"
	"time"
)

var cache map[string]structs.CountryCacheEntry

func InitCache() error {
	var err error
	cache, err = countries_db.GetAllCountries()
	if err != nil {
		if err.Error() == constants.CountryDBIsEmpty {
			return nil
		}
		return err
	} else {
		return nil
	}
}

func GetCountry(alphaCode string) (string, error) {
	c, exists := cache[alphaCode]
	if exists && time.Since(c.Time).Hours() < constants.CacheExpire {
		return c.CountryName, nil
	} else {
		countryName, err := updateCountry(alphaCode)
		if err != nil {
			return "", err
		}
		return countryName, nil
	}
}

func updateCountry(alphaCode string) (string, error) {
	_ = removeCountry(alphaCode)
	_ = countries_db.DeleteCountry(alphaCode)
	countryName, err := countries_api.GetCountryName(alphaCode)
	if err != nil {
		if err.Error() == constants.CountryNotFoundError {
			return "", err
		}
	}
	_ = countries_db.AddCountry(alphaCode, countryName)
	_ = addCountry(alphaCode, countryName)

	return countryName, nil

}

func addCountry(alphaCode string, countryName string) error {
	_, exists := cache[alphaCode]
	if exists {
		log.Println(errors.New(constants.CountryAlreadyInCache))
		return errors.New(constants.CountryAlreadyInCache)
	} else {
		cache[alphaCode] = structs.CountryCacheEntry{
			AlphaCode:   alphaCode,
			CountryName: countryName,
			Time:        time.Now(),
		}
		return nil
	}
}

func removeCountry(alphaCode string) error {
	_, exists := cache[alphaCode]
	if !exists {
		log.Println(errors.New(constants.CountryNotInCache))
		return errors.New(constants.CountryNotInCache)
	} else {
		delete(cache, alphaCode)
		return nil
	}
}
