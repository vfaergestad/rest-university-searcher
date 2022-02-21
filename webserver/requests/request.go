package requests

import (
	"net/http"
)

func CreateAndDoRequest(method string, url string) (*http.Response, error) {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("content-type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return res, nil
}
