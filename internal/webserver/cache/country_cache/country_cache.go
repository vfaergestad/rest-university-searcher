package country_cache

import (
	"assignment-2/internal/webserver/api_requests/countries_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/countries_db"
	"assignment-2/internal/webserver/structs"
	"errors"
	"time"
)

const cacheExpire = 1200

var cache map[string]structs.CountryCacheEntry

func InitCache() error {
	var err error
	cache, err = countries_db.GetAllCountries()
	if err != nil {
		return err
	} else {
		return nil
	}
}

//TODO: FIX
func GetCountry(alphaCode string) (structs.CountryCacheEntry, error) {
	c, exists := cache[alphaCode]
	if !exists {
		return structs.CountryCacheEntry{}, errors.New(constants.CountryNotFoundError)
	} else {
		if time.Since(c.Time).Hours() > cacheExpire {
			err := RemoveCountry(alphaCode)
			if err != nil {
				return structs.CountryCacheEntry{}, err
			}
			countryName, err := countries_api.GetCountryName(alphaCode)
			if err != nil {
				return structs.CountryCacheEntry{}, err
			}

		}
		return c, nil
	}
}

func AddCountry(alphaCode string, countryName string) error {
	_, exists := cache[alphaCode]
	if exists {
		return errors.New(constants.CountryAlreadyInCache)
	} else {
		cache[alphaCode] = structs.CountryCacheEntry{
			AlphaCode: alphaCode,
			Name:      countryName,
			Time:      time.Now(),
		}
		return nil
	}
}

func RemoveCountry(alphaCode string) error {
	_, exists := cache[alphaCode]
	if !exists {
		return errors.New(constants.CountryNotInCache)
	} else {
		delete(cache, alphaCode)
		return nil
	}
}
