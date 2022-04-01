package countries_db

// Countries_db handles all communication to the countries' collection in the database.

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/structs"
	"cloud.google.com/go/firestore"
	"errors"
	"google.golang.org/api/iterator"
	"regexp"
	"strings"
	"time"
)

// collection is the name of the collection in the database.
const collection = "countries"

// GetAllCountries returns all country-entries in the database.
func GetAllCountries() (map[string]structs.CountryCacheEntry, error) {
	resultMap := map[string]structs.CountryCacheEntry{}
	dbEmpty := true

	iter := db.GetClient().Collection(collection).Documents(db.GetContext())

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		dbEmpty = false

		m := doc.Data()

		// Removes the entries that are older than the cache expire limit
		if time.Since(m["time"].(time.Time)).Hours() > constants.CacheExpire {
			_ = DeleteCountry(m["alphaCode"].(string))
		}

		// Creates a cache entry struct for each element, and puts it into the result map
		resultMap[m["alphaCode"].(string)] = structs.CountryCacheEntry{
			AlphaCode:   m["alphaCode"].(string),
			CountryName: m["countryName"].(string),
			Time:        m["time"].(time.Time),
		}

	}

	// Return an empty map if the database is empty
	if dbEmpty {
		return map[string]structs.CountryCacheEntry{}, errors.New(constants.CountryDBIsEmpty)
	}

	return resultMap, nil
}

// AddCountry adds a country to the database.
func AddCountry(alphaCode string, countryName string) error {
	// Checks if the alphaCode is valid
	match, err := regexp.MatchString(constants.AlphaCodeRegex, alphaCode)
	if err != nil {
		return err
	} else if !match {
		return errors.New(constants.MalformedAlphaCodeError)
	}

	alphaCode = strings.ToUpper(alphaCode)
	_, err = db.GetClient().Collection(collection).Doc(alphaCode).Set(db.GetContext(), map[string]interface{}{
		"alphaCode":   alphaCode,
		"countryName": countryName,
		"time":        firestore.ServerTimestamp,
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteCountry deletes a country from the database.
func DeleteCountry(alphaCode string) error {
	_, err := db.GetClient().Collection(collection).Doc(alphaCode).Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}
