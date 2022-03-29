package policy_cache

import (
	"assignment-2/internal/webserver/api_requests/policy_api"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/policies_db"
	"assignment-2/internal/webserver/structs"
	"strings"
)

func GetPolicy(countryCode string, scope string) (structs.PolicyResponse, error) {
	policy, err := policies_db.GetPolicy(countryCode, scope)
	if err != nil {
		policy, err = getPolicyFromApi(countryCode, scope)
		if err != nil {
			return structs.PolicyResponse{}, err
		} else {
			err = policies_db.AddPolicy(policy)
			if err != nil && err.Error() != constants.TestModeActiveError {
				return structs.PolicyResponse{}, err
			}
		}
	}
	return policy, nil
}

func getPolicyFromApi(country string, scope string) (structs.PolicyResponse, error) {
	scopeSplit := strings.Split(scope, "-")
	year := scopeSplit[0]
	month := scopeSplit[1]
	day := scopeSplit[2]

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
