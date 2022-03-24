package webhooks_db

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/structs"
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

		m := doc.Data()

		// Creates a cache entry struct for each element, and puts it into the result map
		resultMap[m["webhookId"].(string)] = structs.Webhook{
			WebhookId: m["webhookId"].(string),
			Url:       m["url"].(string),
			Country:   m["country"].(string),
			Calls:     m["calls"].(int),
		}

	}

	// Return an empty map if the database is empty
	if dbEmpty {
		return map[string]structs.Webhook{}, errors.New(constants.WebhookDBIsEmpty)
	}

	return resultMap, nil
}

func GetWebHook(webhookId string) (structs.Webhook, error) {
	res := db.GetClient().Collection(collection).Doc(webhookId)
	doc, err := res.Get(db.GetContext())
	if err != nil {
		return structs.Webhook{}, err
	}

	m := doc.Data()
	_, exists := m["webhookId"]
	if !exists {
		return structs.Webhook{}, errors.New(constants.WebhookNotFoundError)
	} else {
		webhook := structs.Webhook{
			WebhookId: m["webhookId"].(string),
			Url:       m["url"].(string),
			Country:   m["country"].(string),
			Calls:     m["calls"].(int),
		}
		return webhook, nil
	}
}

func AddWebHook(webhook structs.Webhook) error {
	webhookId := webhook.WebhookId
	_, err := db.GetClient().Collection(collection).Doc(webhookId).Set(db.GetContext(), map[string]interface{}{
		"webhookId": webhookId,
		"url":       webhook.Url,
		"country":   webhook.Country,
		"calls":     webhook.Calls,
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteWebhook(webhook structs.Webhook) error {
	_, err := db.GetClient().Collection(collection).Doc(webhook.WebhookId).Delete(db.GetContext())
	if err != nil {
		return err
	} else {
		return nil
	}
}
