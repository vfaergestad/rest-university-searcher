package constants

const (
	DEFAULT_PORT = "8080"

	// The paths that will be handled by each handler
	DEFAULT_PATH       = "/corona/"
	CASES_PATH         = "/corona/v1/cases/"
	POLICY_PATH        = "/corona/v1/policy/"
	STATUS_PATH        = "/corona/v1/status/"
	NOTIFICATIONS_PATH = "/corona/v1/notifications/"

	POLICY_API_URL = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"
)
