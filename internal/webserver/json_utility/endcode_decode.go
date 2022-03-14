package json_utility

import (
	"encoding/json"
	"net/http"
)

func DecodeResponse(res *http.Response, target interface{}) (interface{}, error) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&target); err != nil {
		return nil, err
	}
	return target, nil
}
