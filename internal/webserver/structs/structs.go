package structs

type Status struct {
	CasesApi  int    `json:"cases_api"`
	PolicyApi int    `json:"policy_api"`
	Webhooks  int    `json:"webhooks"`
	Version   string `json:"version"`
	Uptime    string `json:"uptime"`
}

type PolicyApiResponse struct {
	PolicyActions  []interface{}          `json:"policyActions"`
	StringencyData map[string]interface{} `json:"stringencyData"`
}
