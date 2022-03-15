package constants

import (
	"errors"
	"fmt"
)

const (
	MALFORMED_ALPHACODE_ERROR  = "the alpha-code must be 3 letters"
	MALFORMED_COVID_YEAR_ERROR = "not a valid covid year. Must be between 2019 and 2030"
	MALFORMED_MONTH_ERROR      = "not a valid month. Must be between 01 and 12 / 1 and 12"
	MALFORMED_DAY_ERROR        = "not a valid day. Must be between 01 and 31"
	INVALID_METHOD_ERROR       = "not a valid http method"

	LINK_TO_DOCS = "https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022-workspace/vegarfae/assignment-2/-/blob/main/README.md"
)

func getDocumentationError() error {
	return errors.New(fmt.Sprintf("See %s for documentation", LINK_TO_DOCS))
}

func GetNotValidPathError() error {
	return errors.New(
		fmt.Sprintf("Not a valid endpoint. \n\n"+
			"Please use paths %s, %s, %s, or %s. \n\n"+
			"%s", POLICY_PATH, CASES_PATH, NOTIFICATIONS_PATH, STATUS_PATH, getDocumentationError().Error()),
	)
}

func GetBadPoliciesRequestError() error {
	return errors.New(fmt.Sprintf("Not a valid request. Format: /corona/v1/policy/{:country_name}{?scope=YYYY-MM-DD}&{limit={:number}}}\n"+
		"\n"+
		"%s", getDocumentationError().Error()))
}
