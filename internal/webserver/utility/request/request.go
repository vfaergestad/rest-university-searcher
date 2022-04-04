package request

// Request handles all http requests

import (
	"io"
	"log"
	"net/http"
)

// GetRequest sends a GET request to the given url and returns the response.
func GetRequest(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

// PostRequest sends a POST request to the given url with the given body and returns the response.
func PostRequest(url string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

// HeadRequest sends a HEAD request to the given url and returns the response.
func HeadRequest(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
