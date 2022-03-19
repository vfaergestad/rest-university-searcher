package countries_db

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

const collection = "countries"

func GetAllCountries() (map[string]structs.CountryCacheEntry, error) {
	resultMap := map[string]structs.CountryCacheEntry{}
	dbEmpty := true

	iter := db.GetClient().Collection(collection).Documents(db.GetContext())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		dbEmpty = false
		if err != nil {
			return nil, err
		}
		// Note: You can access the document ID using "doc.Ref.ID"

		// A message map with string keys. Each key is one field, like "text" or "timestamp"
		m := doc.Data()

		resultMap[m["alphaCode"].(string)] = structs.CountryCacheEntry{
			AlphaCode:   m["alphaCode"].(string),
			CountryName: m["countryName"].(string),
			Time:        m["time"].(time.Time),
		}

	}

	if dbEmpty {
		return map[string]structs.CountryCacheEntry{}, errors.New(constants.CountryDBIsEmpty)
	}

	return resultMap, nil
}

func GetCountry(alphaCode string) (string, error) {
	res := db.GetClient().Collection(collection).Doc(alphaCode)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return "", err
	}

	m := doc.Data()
	country, exists := m["countryName"]
	if !exists {
		return "", errors.New(constants.CountryNotFoundError)
	} else {
		return country.(string), nil
	}
}

func AddCountry(alphaCode string, countryName string) error {
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

func DeleteCountry(alphaCode string) error {
	_, err := db.GetClient().Collection(collection).Doc(alphaCode).Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}
