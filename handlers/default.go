package handlers

import (
	"assignment-1/constants"
	"fmt"
	"net/http"
)

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Please use paths %s, %s, or %s.", constants.DIAG_PATH, constants.UNIINFO_PATH, constants.NEIGHBOURUNIS_PATH), http.StatusOK)
}
