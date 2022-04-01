package country_cache

// CountryCache is a cache of country codes and their corresponding names.
import (
	"assignment-2/internal/webserver/api_requests/countries_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/countries_db"
	"assignment-2/internal/webserver/structs"
	"errors"
	"log"
	"time"
)

// Cache that holds the translations.
var cache map[string]structs.CountryCacheEntry

// InitCache initializes the cache by getting all the entries from the database.
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

// GetCountry returns the name of the country with the given ISO 3166-1 alpha-3 code.
// If the country is not found, it requests the country from the API and adds it to the cache.
func GetCountry(alphaCode string) (string, error) {
	// Checks if the country is in the cache.
	c, exists := cache[alphaCode]
	// Checks if the country is expired.
	if exists && time.Since(c.Time).Hours() < constants.CacheExpire {
		return c.CountryName, nil
	} else {
		// If the country is not in the cache, it requests it from the API.
		countryName, err := updateCountry(alphaCode)
		if err != nil {
			return "", err
		}
		return countryName, nil
	}
}

// updateCountry requests the country from the API and adds it to the cache and the database.
func updateCountry(alphaCode string) (string, error) {
	// Removes the country from the cache and database.
	_ = removeCountry(alphaCode)
	_ = countries_db.DeleteCountry(alphaCode)
	// Requests the country from the API.
	countryName, err := countries_api.GetCountryName(alphaCode)
	if err != nil {
		if err.Error() == constants.CountryNotFoundError {
			return "", err
		}
	}
	// Adds the country to the cache and database.
	_ = countries_db.AddCountry(alphaCode, countryName)
	_ = addCountry(alphaCode, countryName)

	return countryName, nil

}

// addCountry adds the country to the cache.
func addCountry(alphaCode string, countryName string) error {
	// Checks if the country is already in the cache.
	_, exists := cache[alphaCode]
	if exists {
		log.Println(errors.New(constants.CountryAlreadyInCache))
		return errors.New(constants.CountryAlreadyInCache)
	} else {
		// Adds the country to the cache.
		cache[alphaCode] = structs.CountryCacheEntry{
			AlphaCode:   alphaCode,
			CountryName: countryName,
			Time:        time.Now(),
		}
		return nil
	}
}

// removeCountry removes the country from the cache.
func removeCountry(alphaCode string) error {
	// Checks if the country is in the cache.
	_, exists := cache[alphaCode]
	if !exists {
		log.Println(errors.New(constants.CountryNotInCache))
		return errors.New(constants.CountryNotInCache)
	} else {
		// Removes the country from the cache.
		delete(cache, alphaCode)
		return nil
	}
}
