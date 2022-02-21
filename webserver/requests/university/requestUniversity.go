package university

import (
	"assignment-1/constants"
	"assignment-1/structs"
	"assignment-1/webserver/requests"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func RequestUniversity(search string) ([]structs.University, error) {
	//fmt.Println(url)

	query := "search?name_contains=" + search
	url := constants.UNIVERSITIESAPI_URL + query

	res, err := requests.CreateAndDoRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	switch {
	case res.StatusCode == http.StatusNotFound:
		return []structs.University{}, errors.New(fmt.Sprintf("%d University not found", res.StatusCode))
	case res.StatusCode != http.StatusOK:
		return []structs.University{}, errors.New(fmt.Sprintf("Status code returned from universityAPI: %d", res.StatusCode))
	}

	var universities []structs.University
	if universities, err = decodeUniversities(res); err != nil {
		return nil, err
	}

	return universities, nil

}

func decodeUniversities(res *http.Response) ([]structs.University, error) {
	decoder := json.NewDecoder(res.Body)
	var universities []structs.University
	if err := decoder.Decode(&universities); err != nil {
		log.Println(err)
		return nil, err

	}
	return universities, nil
}
