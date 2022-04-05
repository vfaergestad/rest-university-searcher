package notifications_handler

// Notifications_endpoint handles al incoming traffic to the /corona/v1/notifications endpoint.

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/encode_struct"
	"encoding/json"
	"net/http"
	"path"
	"regexp"
	"strings"
)

// HandlerNotifications is the entry point for the handler.
func HandlerNotifications(w http.ResponseWriter, r *http.Request) {
	// Checks which method is used, and redirects the request to the corresponding handler.
	switch r.Method {
	case http.MethodGet:
		getWebhook(w, r)
	case http.MethodPost:
		registerWebhook(w, r)
	case http.MethodDelete:
		deleteWebhook(w, r)
	default:
		http.Error(w, "Method is not supported. Currently Get, Post, and Delete are supported.", http.StatusMethodNotAllowed)
		return
	}

}

// getWebhook handles requests to get webhooks, either a single one, or all of them.
func getWebhook(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	// Check if the given path is valid, and finds out if the request is for a single webhook or all of them.
	switch len(pathList) {
	case 4:
		getAllWebhooks(w)
	case 5:
		getSingleWebhook(w, r)
	default:
		http.Error(w, constants.GetBadGetWebhookRequestError().Error(), http.StatusBadRequest)
		return
	}
}

// getAllWebhooks handler that responds with all the webhooks in the database.
func getAllWebhooks(w http.ResponseWriter) {
	// Gets all the webhooks from the database.
	webhooks, err := webhooks_db.GetAllWebhooks()
	if err != nil {
		if err.Error() == constants.WebhookDBIsEmpty {
			http.Error(w, err.Error(), http.StatusOK)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Encodes the webhooks to json and writes it to the response.
	err = encode_struct.EncodeStruct(w, webhooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// getSingleWebhook handles requests to get a single webhook.
func getSingleWebhook(w http.ResponseWriter, r *http.Request) {
	// Finds the webhook id in the url.
	cleanPath := path.Clean(r.URL.Path)
	webhookId := path.Base(cleanPath)

	// Gets the webhook from the database.
	webhook, err := webhooks_db.GetWebhook(webhookId)
	if err != nil {
		if err.Error() == constants.WebhookNotFoundError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Writes the webhook to the response.
	err = encode_struct.EncodeStruct(w, webhook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// RegisterWebhook handles requests to register a new webhook.
func registerWebhook(w http.ResponseWriter, r *http.Request) {
	// Struct for responding with the webhook id.
	type responseStruct struct {
		WebhookId string `json:"webhook_id"`
	}

	// Decodes the request body to a webhook.
	var webhook structs.Webhook
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&webhook)

	// Checks if the webhook url is empty.
	if webhook.Url == "" {
		http.Error(w, "url cannot be empty", http.StatusBadRequest)
		return
	}

	// Checks if the webhook country is empty.
	if webhook.Country == "" {
		http.Error(w, "country cannot be empty", http.StatusBadRequest)
		return
	}

	// Checks if the webhook calls is empty.
	if webhook.Calls < 1 {
		http.Error(w, "calls must be 1 or greater", http.StatusBadRequest)
		return
	}

	// Checks if the webhook country is an alpha code. If it is, the corresponding country is retrieved from the cache.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, webhook.Country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if match {
		webhook.Country, err = country_cache.GetCountry(webhook.Country)
		if err != nil {
			switch err.Error() {
			case constants.CountryNotFoundError:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	// Adds the webhook to the database.
	webhookId, err := webhooks_db.AddWebhook(webhook.Url, webhook.Country, webhook.Calls)
	if err != nil {
		if err.Error() == constants.WebhookAlreadyExistingError {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Encodes the webhook id to json and writes it to the response.
	response := responseStruct{WebhookId: webhookId}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = encode_struct.EncodeStruct(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// deleteWebhook handles requests to delete a webhook.
func deleteWebhook(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	// Check if the given path is valid
	if len(pathList) != 5 {
		http.Error(w, constants.GetBadDeleteWebhookRequestError().Error(), http.StatusBadRequest)
		return
	}

	// Extracts the webhook id from the url.
	webhookId := path.Base(cleanPath)

	// Deletes the webhook from the database.
	err := webhooks_db.DeleteWebhook(webhookId)
	if err != nil {
		if err.Error() == constants.WebhookNotFoundError {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// TODO: Response as JSON?
		http.Error(w, "webhook deleted", http.StatusOK)
		return
	}
}
