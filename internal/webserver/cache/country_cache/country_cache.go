package country_cache

import (
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
	if exists {
		if time.Since(c.Time).Hours() < cacheExpire {
			return c.CountryName, nil
		} else {
			_ = removeCountry(alphaCode)
			return "", errors.New(constants.ExpiredCacheEntry)
		}
	} else {
		return "", errors.New(constants.CountryNotInCache)
	}
}

func AddCountry(alphaCode string, countryName string) error {
	_, exists := cache[alphaCode]
	if exists {
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
		return errors.New(constants.CountryNotInCache)
	} else {
		delete(cache, alphaCode)
		return nil
	}
}
