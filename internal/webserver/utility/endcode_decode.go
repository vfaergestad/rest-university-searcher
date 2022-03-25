package utility

import (
	"encoding/json"
	"net/http"
)

func EncodeStruct(w http.ResponseWriter, target interface{}) error {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(target)
	return err
}
