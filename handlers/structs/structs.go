package structs

type Diagnose struct {
	UniversitiesApi string `json:"universitiesapi"`
	CountriesApi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          int    `json:"uptime"`
}

type UniAndCountry struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
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
	CCA2      string                 `json:"cca2"`
}

func CombineUniCountry(u University, c Country) UniAndCountry {
	uniInfo := UniAndCountry{
		Name:      u.Name,
		Country:   u.Country,
		Isocode:   u.AlphaTwoCode,
		Webpages:  u.WebPages,
		Languages: c.Languages,
		Map:       c.Maps["openStreetMaps"],
	}
	return uniInfo

}
