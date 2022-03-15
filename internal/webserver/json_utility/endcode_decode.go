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

func EncodeStruct(w http.ResponseWriter, target interface{}) error {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(target)
	return err
}
