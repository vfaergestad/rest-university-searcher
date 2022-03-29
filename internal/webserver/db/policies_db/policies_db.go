package policies_db

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/hash_util"
	"errors"
	"time"
)

const collection = "policies"

var testMode = false

/*
func GetAllPolicies() (map[string]structs.PolicyResponseCacheEntry, error) {
	resultMap := map[string]structs.PolicyResponseCacheEntry{}
	dbEmpty := true

	iter := db.GetClient().Collection(collection).Documents(db.GetContext())

	for {
		// Gets the next item in the collection
		doc, err := iter.Next()

		// Stops the loop if there is no more elements
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
			_ = DeletePolicy(doc.D)
		}

		// Creates a cache entry struct for each element, and puts it into the result map
		resultMap[m["id"].(string)] = structs.PolicyResponseCacheEntry{
			CountryCode: m["countryCode"].(string),
			Scope:       m["scope"].(string),
			Stringency:  m["stringency"].(float64),
			Policies:    m["policies"].(int),
			Time:        m["time"].(time.Time),
		}

	}

	// Return an empty map if the database is empty
	if dbEmpty {
		return map[string]structs.PolicyResponseCacheEntry{}, errors.New(constants.PolicyDBIsEmpty)
	}

	return resultMap, nil
}*/

func GetPolicy(country string, scope string) (structs.PolicyResponse, error) {
	if testMode {
		return structs.PolicyResponse{}, errors.New(constants.TestModeActiveError)
	}
	id := hash_util.HashPolicy(country, scope)
	res := db.GetClient().Collection(collection).Doc(id)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.PolicyResponse{}, errors.New(constants.PolicyNotInCacheError)
	}

	var policy structs.PolicyResponse
	err = doc.DataTo(&policy)
	if err != nil {
		return structs.PolicyResponse{}, err
	}

	if time.Since(policy.Time).Hours() > constants.CacheExpire {
		_ = DeletePolicy(country, scope)
		return structs.PolicyResponse{}, errors.New(constants.ExpiredCacheEntry)
	}

	return policy, nil

}

func AddPolicy(policy structs.PolicyResponse) error {
	if testMode {
		return errors.New(constants.TestModeActiveError)
	}
	id := hash_util.HashPolicy(policy.CountryCode, policy.Scope)
	_, err := db.GetClient().Collection(collection).Doc(id).Set(db.GetContext(), map[string]interface{}{
		"countryCode": policy.CountryCode,
		"scope":       policy.Scope,
		"stringency":  policy.Stringency,
		"policies":    policy.Policies,
		"time":        time.Now(),
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeletePolicy(country string, scope string) error {
	if testMode {
		return errors.New(constants.TestModeActiveError)
	}
	id := hash_util.HashPolicy(country, scope)
	_, err := db.GetClient().Collection(collection).Doc(id).Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SetTestMode() {
	testMode = true
}
