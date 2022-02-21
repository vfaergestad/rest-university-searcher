package structs

type Diagnose struct {
	UniversitiesApi string `json:"universitiesapi"`
	CountriesApi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          int    `json:"uptime"`
}

type UniAndCountry struct {
	Name      string            `json:"name,omitempty"`
	Country   string            `json:"country,omitempty"`
	Isocode   string            `json:"isocode,omitempty"`
	WebPages  []string          `json:"webpages,omitempty"`
	Languages map[string]string `json:"languages,omitempty"`
	Map       string            `json:"map,omitempty"`
}

type University struct {
	Name         string   `json:"name"`
	Country      string   `json:"country"`
	AlphaTwoCode string   `json:"alpha_two_code"`
	WebPages     []string `json:"web_pages"`
}

type Country struct {
	Name      map[string]interface{} `json:"name"`
	Languages map[string]string      `json:"languages"`
	Maps      map[string]string      `json:"maps"`
	Borders   []string               `json:"borders"`
	Alpha     string                 `json:"cca3"`
}

func CombineUniCountry(u University, c Country, fields ...string) UniAndCountry {
	var uniInfo UniAndCountry
	if len(fields) > 0 {
		for _, f := range fields {
			switch f {
			case "name":
				uniInfo.Name = u.Name

			case "country":
				uniInfo.Country = u.Country

			case "isocode":
				uniInfo.Isocode = u.AlphaTwoCode

			case "webpages":
				uniInfo.WebPages = u.WebPages

			case "languages":
				uniInfo.Languages = c.Languages

			case "map":
				uniInfo.Map = c.Maps["openStreetMaps"]
			}
		}
	} else {
		uniInfo = UniAndCountry{
			Name:      u.Name,
			Country:   u.Country,
			Isocode:   u.AlphaTwoCode,
			WebPages:  u.WebPages,
			Languages: c.Languages,
			Map:       c.Maps["openStreetMaps"],
		}
	}
	return uniInfo

}
