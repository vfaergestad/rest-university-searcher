package structs

import "time"

type CasesResponse struct {
	Country    string  `json:"country"`
	Date       string  `json:"date"`
	Confirmed  int     `json:"confirmed"`
	Recovered  int     `json:"recovered"`
	Deaths     int     `json:"deaths"`
	GrowthRate float64 `json:"growth_rate"`
}

type CountryCacheEntry struct {
	AlphaCode   string
	CountryName string
	Time        time.Time
}

type Webhook struct {
	WebhookId string `json:"webhook_id"`
	Url       string `json:"url"`
	Country   string `json:"country"`
	Calls     int    `json:"calls"`
}
