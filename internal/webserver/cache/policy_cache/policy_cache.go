package policy_cache

// Policy_cache handles all request for policies. It gets entries from the database,
// or from the API if the policy is not in the database.

import (
	"assignment-2/internal/webserver/api_requests/policy_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/policies_db"
	"assignment-2/internal/webserver/structs"
	"strings"
)

// GetPolicy returns the policy for the given country and scope.
func GetPolicy(countryCode string, scope string) (structs.PolicyResponse, error) {
	// Get the policy from the database.
	policy, err := policies_db.GetPolicy(countryCode, scope)
	if err != nil {
		// If the policy is not in the database, get it from the API.
		policy, err = getPolicyFromApi(countryCode, scope)
		if err != nil {
			return structs.PolicyResponse{}, err
		} else {
			// Save the policy in the database.
			err = policies_db.AddPolicy(policy)
			if err != nil && err.Error() != constants.TestModeActiveError {
				return structs.PolicyResponse{}, err
			}
		}
	}
	return policy, nil
}

// getPolicyFromApi gets the policy for the given country and date from the API.
func getPolicyFromApi(country string, scope string) (structs.PolicyResponse, error) {
	// Extracts the date from the scope.
	scopeSplit := strings.Split(scope, "-")
	year := scopeSplit[0]
	month := scopeSplit[1]
	day := scopeSplit[2]

	// Gets the policy from the API.
	stringency, polices, err := policy_api.GetStringencyAndPolicies(country, year, month, day)
	if err != nil {
		return structs.PolicyResponse{}, err
	} else {
		return structs.PolicyResponse{
			CountryCode: strings.ToTitle(country),
			Scope:       scope,
			Stringency:  stringency,
			Policies:    polices,
		}, nil

	}
}
