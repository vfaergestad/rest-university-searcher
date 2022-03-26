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

func GetAllWebHooks() (map[string]structs.Webhook, error) {
	resultMap := map[string]structs.Webhook{}
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
		resultMap[doc.Ref.ID] = webhook

	}

	// Return an empty map if the database is empty
	if dbEmpty {
		return map[string]structs.Webhook{}, errors.New(constants.WebhookDBIsEmpty)
	}

	return resultMap, nil
}

func GetWebHook(url string, country string, calls int) (structs.Webhook, error) {
	webhookId := hash_util.HashWebhook(url, country, calls)
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.Webhook{}, err
	}

	var webhook structs.Webhook
	err = doc.DataTo(&webhook)
	if err != nil {
		return structs.Webhook{}, err
	}
	return webhook, nil
}

func GetWebHookById(webhookId string) (structs.Webhook, error) {
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.Webhook{}, err
	}

	var webhook structs.Webhook
	err = doc.DataTo(&webhook)
	if err != nil {
		return structs.Webhook{}, err
	}
	return webhook, nil
}

func AddWebHook(url string, country string, calls int) error {
	webhookId := hash_util.HashWebhook(url, country, calls)
	_, err := db.GetClient().Collection(collection).Doc(webhookId).Set(db.GetContext(), map[string]interface{}{
		"webhookId": webhookId,
		"url":       url,
		"country":   country,
		"calls":     calls,
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteWebhook(url string, country string, calls int) error {
	webhookId := hash_util.HashWebhook(url, country, calls)
	_, err := db.GetClient().Collection(collection).Doc(webhookId).Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}
