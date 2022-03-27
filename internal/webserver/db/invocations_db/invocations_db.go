package invocations_db

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"errors"
)

const collection = "invocations"

func AddInvocation(country string) error {
	count, err := GetInvocation(country)
	if err != nil {
		count = 0
	}

	_, _ = db.GetClient().Collection(collection).Doc(country).Set(db.GetContext(), map[string]int{
		"count": count + 1,
	})
	return nil
}

func GetInvocation(country string) (int, error) {
	res := db.GetClient().Collection(collection).Doc(country)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return -1, err
	}
	if !doc.Exists() {
		return 0, errors.New(constants.WebhookNotFoundError)
	}
	return int(doc.Data()["count"].(int64)), nil
}
