package structs

type CasesResponse struct {
	Country    string  `json:"country"`
	Date       string  `json:"date"`
	Confirmed  int     `json:"confirmed"`
	Recovered  int     `json:"recovered"`
	Deaths     int     `json:"deaths"`
	GrowthRate float64 `json:"growth_rate"`
}
