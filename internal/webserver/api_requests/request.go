package api_requests

import (
	"io"
	"net/http"
)

func GetRequest(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func PostRequest(url string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func HeadRequest(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return res, nil
}
