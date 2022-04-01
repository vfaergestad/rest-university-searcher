package constants

const (
	// The paths to the different handlers/endpoints
	CasesPath         = "/corona/v1/cases/"
	PolicyPath        = "/corona/v1/policy/"
	StatusPath        = "/corona/v1/status/"
	NotificationsPath = "/corona/v1/notifications/"
	DefaultPath       = "/corona/"

	// Regexes for different validations
	AlphaCodeRegex = "^[a-zA-Z]{3}$"
	YearRegex      = "^(2019|202\\d)$"
	MonthRegex     = "^(0[1-9]|1[012])$"
	DayRegex       = "^(0[1-9]|[12]\\d|3[01])$"
	NoNumbersRegex = "^([^0-9]*)$"

	// CacheExpire Amount of hours before a cache entry expires
	CacheExpire = 1200
)

// ServiceAccountLocation is the location of the service account file
var ServiceAccountLocation = "serviceAccountKey.json"

// SetTestServiceAccountLocation sets the location of the service account file to a location relative to the test main.
func SetTestServiceAccountLocation() {
	ServiceAccountLocation = "./../../../../serviceAccountKey.json"
}

// The urls to the different APIs
var (
	CasesApiUrl        = "https://covid19-graphql.now.sh"
	CountryAPIUrl      = "https://restcountries.com/v3.1/"
	PolicyApiStatusUrl = "https://covidtrackerapi.bsg.ox.ac.uk/api/"
	PolicyApiUrl       = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"
)

// SetTestCasesApiUrl sets the url to a given test url
func SetTestCasesApiUrl(url string) {
	CasesApiUrl = url + "/"
}

// SetTestCountryAPIUrl sets the url to a given test url
func SetTestCountryAPIUrl(url string) {
	CountryAPIUrl = url + "/"
}

// SetTestPolicyApiUrl sets the url to a given test url
func SetTestPolicyApiUrl(url string) {
	PolicyApiUrl = url + "/"
}
