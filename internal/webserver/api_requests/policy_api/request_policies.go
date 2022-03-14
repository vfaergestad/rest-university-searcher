package policy_api

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/webserver/api_requests"
	"assignment-2/internal/webserver/json_utility"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const (
	POLICY_API_STATUS_URL = "https://covidtrackerapi.bsg.ox.ac.uk/api/"
	POLICY_API_URL        = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/"

	ALPHA_CODE_REGEX = "/^[a-z]|[A-Z]{3}$/"
	YEAR_REGEX       = "/^(2019|202\\d)$/"
	MONTH_REGEX      = "/^(0[1-9]|1[012])$/"
	DAY_REGEX        = "/^(0[1-9]|[12]\\d|3[01])$/"

	/*
		// Got from: https://stackoverflow.com/questions/41085409/country-code-validation-with-iso
		ALPHA_CODE_REGEX = "/^A(BW|FG|GO|IA|L[AB]|ND|R[EGM]|SM|T[A\nFG]|U[ST]|ZE)|B(DI|E[LNS]|FA|G[DR]|H[RS]|IH|L[MRZ]|MU|" +
			"OL|\nR[ABN]|TN|VT|WA)|C(A[FN]|CK|H[ELN]|IV|MR|O[DGKLM]|PV|RI|U\n[BW]|XR|Y[MP]|ZE)|D(EU|JI|MA|NK|OM|ZA)|E(CU|" +
			"GY|RI|S[HPT]|\nTH)|F(IN|JI|LK|R[AO]|SM)|G(AB|BR|EO|GY|HA|I[BN]|LP|MB|N[B\nQ]|R[CDL]|TM|U[FMY])|H(KG|MD|ND|RV|" +
			"TI|UN)|I(DN|MN|ND|OT|R\n[LNQ]|S[LR]|TA)|J(AM|EY|OR|PN)|K(AZ|EN|GZ|HM|IR|NA|OR|WT)\n|L(AO|B[NRY]|CA|IE|KA|SO|" +
			"TU|UX|VA)|M(A[CFR]|CO|D[AGV]|EX|\nHL|KD|L[IT]|MR|N[EGP]|OZ|RT|SR|TQ|US|WI|Y[ST])|N(AM|CL|ER\n|FK|GA|I[CU]|LD|" +
			"OR|PL|RU|ZL)|OMN|P(A[KN]|CN|ER|HL|LW|NG|O\nL|R[IKTY]|SE|YF)|QAT|R(EU|OU|US|WA)|S(AU|DN|EN|G[PS]|HN|J\nM|L[BEV]|" +
			"MR|OM|PM|RB|SD|TP|UR|V[KN]|W[EZ]|XM|Y[CR])|T(C[A\nD]|GO|HA|JK|K[LM]|LS|ON|TO|U[NRV]|WN|ZA)|U(GA|KR|MI|RY|SA\n|" +
			"ZB)|V(AT|CT|EN|GB|IR|NM|UT)|W(LF|SM)|YEM|Z(AF|MB|WE)$/ix"
	*/

)

type PolicyApiResponse struct {
	PolicyActions  []interface{}          `json:"policyActions"`
	StringencyData map[string]interface{} `json:"stringencyData"`
}

func GetStatusCode() (int, error) {

	res, err := api_requests.DoRequest(POLICY_API_STATUS_URL, http.MethodHead)
	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil

}

func GetResponse(alphaCode string, year string, month string, day string) (PolicyApiResponse, error) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(ALPHA_CODE_REGEX, alphaCode)
	if err != nil {
		return PolicyApiResponse{}, err
	} else if !match {
		return PolicyApiResponse{}, errors.New(constants.MALFORMED_ALPHACODE_ERROR)
	}

	// Check if given date is valid.
	_, err = checkDate(year, month, day)
	if err != nil {
		return PolicyApiResponse{}, err
	}

	// Create URL and request response from API
	url := fmt.Sprintf("%s%s/%s-%s-%s", POLICY_API_URL, alphaCode, year, month, day)
	res, err := api_requests.DoRequest(url, http.MethodGet)
	if err != nil {
		return PolicyApiResponse{}, err
	}

	policy, err := json_utility.DecodeResponse(res, PolicyApiResponse{})
	if err != nil {
		return PolicyApiResponse{}, err
	}

	return policy.(PolicyApiResponse), nil

}

func checkDate(year string, month string, day string) (bool, error) {
	// Checks if given year is a valid covid year. (Between 2019 and 2030)
	match, err := regexp.MatchString(YEAR_REGEX, year)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_COVID_YEAR_ERROR)
	}

	// Checks if given month is a valid month.
	match, err = regexp.MatchString(MONTH_REGEX, month)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_MONTH_ERROR)
	}

	// Checks if given day is a valid day.
	match, err = regexp.MatchString(DAY_REGEX, day)
	if err != nil {
		return false, err
	} else if !match {
		return false, errors.New(constants.MALFORMED_DAY_ERROR)
	}

	return true, nil
}
