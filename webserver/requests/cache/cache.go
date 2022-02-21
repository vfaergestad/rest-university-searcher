package cache

import (
	"assignment-1/structs"
	"errors"
	"fmt"
	"strings"
)

var countryCache = make(map[string]structs.Country)

func AddCountry(c structs.Country) error {
	cName := strings.ToTitle(c.Name["common"].(string))
	if _, exists := countryCache[cName]; !exists {
		countryCache[cName] = c
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s is already in the cache.", cName))
	}
}

func GetCountry(name string) (structs.Country, error) {
	if c, exists := countryCache[strings.ToTitle(name)]; exists {
		return c, nil
	} else {
		return structs.Country{}, errors.New(fmt.Sprintf("%s is not in the cache.", name))
	}
}

func GetCountryByAlpha(alpha string) (structs.Country, error) {
	for _, c := range countryCache {
		if strings.ToUpper(c.Alpha) == strings.ToUpper(alpha) {
			return c, nil
		}
	}
	return structs.Country{}, errors.New(fmt.Sprintf("%s is not in the cache.", alpha))
}
