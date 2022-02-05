package handlers

import "go/types"

type Diagnose struct {
	UniversitiesApi string `json:"universitiesapi"`
	CountriesApi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          int    `json:"uptime"`
}

type Language struct {
	Abbreviation string
	Name         string
}

type University struct {
	Name      string      `json:"name"`
	Country   string      `json:"country"`
	Isocode   string      `json:"isocode"`
	Webpages  types.Array `json:"webpages"`
	Languages Language    `json:"languages"`
	Map       string      `json:"map"`
}
