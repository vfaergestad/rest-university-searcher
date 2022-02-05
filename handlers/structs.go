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

type UniAndCountry struct {
	Name      string      `json:"name"`
	Country   string      `json:"country"`
	Isocode   string      `json:"isocode"`
	Webpages  types.Array `json:"webpages"`
	Languages Language    `json:"languages"`
	Map       string      `json:"map"`
}

type University struct {
	AlphaTwoCode  string   `json:"alpha_two_code"`
	Country       string   `json:"country"`
	StateProvince string   `json:"state_province"`
	Domains       []string `json:"domains"`
	Name          string   `json:"name"`
	WebPages      []string `json:"web_pages"`
}

type Country struct {
	Languages []Language `json:"languages"`
	Maps      []string   `json:"maps"`
}
