package constants

const (
	CasesPath         = "/corona/v1/cases/"
	PolicyPath        = "/corona/v1/policy/"
	StatusPath        = "/corona/v1/status/"
	NotificationsPath = "/corona/v1/notifications/"
	DefaultPath       = "/corona/"

	AlphaCodeRegex = "^[a-zA-Z]{3}$"
	YearRegex      = "^(2019|202\\d)$"
	MonthRegex     = "^(0[1-9]|1[012])$"
	DayRegex       = "^(0[1-9]|[12]\\d|3[01])$"
	NoNumbersRegex = "^([^0-9]*)$"

	CacheExpire = 1200
)

var ServiceAccountLocation = "serviceAccountKey.json"

func SetTestServiceAccountLocation() {
	ServiceAccountLocation = "./../../../../serviceAccountKey.json"
}

var (
	CasesApiUrl        = "https://covid19-graphql.now.sh"
	CountryAPIUrl      = "https://restcountries.com/v3.1/"
	PolicyApiStatusUrl = "https://covidtrackerapi.bsg.ox.ac.uk/api/"
	PolicyApiUrl       = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"
)

func SetTestCasesApiUrl(url string) {
	CasesApiUrl = url + "/"
}

func SetTestCountryAPIUrl(url string) {
	CountryAPIUrl = url + "/"
}

func SetTestPolicyApiStatusUrl(url string) {
	PolicyApiStatusUrl = url + "/"
}

func SetTestPolicyApiUrl(url string) {
	PolicyApiUrl = url + "/"
}
