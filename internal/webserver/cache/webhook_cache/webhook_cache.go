package webhook_cache

import (
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/structs"
	"errors"
)

var cache map[string]structs.Webhook

func InitCache() error {
	var err error
	cache, err = webhooks_db.GetAllWebHooks()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetWebhook(id string) (structs.Webhook, error) {
	w, exists := cache[id]
	if !exists {
		return structs.Webhook{}, errors.New(constants.WebhookNotFoundError)
	} else {
		return w, nil
	}
}

func GetAllWebhooks() map[string]structs.Webhook {
	return cache
}

func AddWebhook(url string, country string, calls int) (string, error) {

}

func DeleteWebhook(id string) error {
	return nil
}

func createId(url string, country string, calls int) (string, error) {

}
