package webhooks_db

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/hash_util"
	"errors"
	"google.golang.org/api/iterator"
)

const collection = "webhooks"

func GetAllWebHooks() ([]structs.Webhook, error) {
	var resultSlice []structs.Webhook
	dbEmpty := true

	iter := db.GetClient().Collection(collection).Documents(db.GetContext())

	for {
		// Gets the next item in the collection
		doc, err := iter.Next()

		// Stops the loop if there is no more elements
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		dbEmpty = false

		var webhook structs.Webhook
		err = doc.DataTo(&webhook)
		if err != nil {
			return nil, err
		}

		// Creates a cache entry struct for each element, and puts it into the result map
		resultSlice = append(resultSlice, webhook)

	}

	// Return an empty map if the database is empty
	if dbEmpty {
		return []structs.Webhook{}, errors.New(constants.WebhookDBIsEmpty)
	}

	return resultSlice, nil
}

func GetWebhook(url string, country string, calls int) (structs.Webhook, error) {
	webhookId := hash_util.HashWebhook(url, country, calls)
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.Webhook{}, err
	}
	if !doc.Exists() {

	}

	var webhook structs.Webhook
	err = doc.DataTo(&webhook)
	if err != nil {
		return structs.Webhook{}, err
	}
	return webhook, nil
}

func GetWebhookById(webhookId string) (structs.Webhook, error) {
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.Webhook{}, err
	}
	if !doc.Exists() {
		return structs.Webhook{}, errors.New(constants.WebhookNotFoundError)
	}

	var webhook structs.Webhook
	err = doc.DataTo(&webhook)
	if err != nil {
		return structs.Webhook{}, err
	}
	return webhook, nil
}

func AddWebhook(url string, country string, calls int) (string, error) {
	webhookId := hash_util.HashWebhook(url, country, calls)

	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, _ := res.Get(db.GetContext())
	if doc.Exists() {
		return "", errors.New(constants.WebhookAlreadyExistingError)
	}

	_, err := res.Set(db.GetContext(), map[string]interface{}{
		"webhookId": webhookId,
		"url":       url,
		"country":   country,
		"calls":     calls,
	})
	if err != nil {
		return "", err
	} else {
		return webhookId, nil
	}
}

func DeleteWebhook(webhookId string) error {
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if !doc.Exists() {
		return errors.New(constants.WebhookNotFoundError)
	}

	_, err = res.Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetDBSize() (int, error) {
	res, err := db.GetClient().Collection(collection).Documents(db.GetContext()).GetAll()
	if err != nil {
		return -1, err
	}
	return len(res), nil
}
