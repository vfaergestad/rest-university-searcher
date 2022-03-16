package api_requests

import (
	"assignment-2/internal/webserver/constants"
	"errors"
	"net/http"
)

func DoRequest(url string, method string) (*http.Response, error) {
	switch method {
	case
		http.MethodHead,
		http.MethodGet,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodTrace:

		r, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			return nil, err
		}

		return res, nil
	default:
		return nil, errors.New(constants.InvalidMethodError)
	}

}
