package handlers

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"net/http"
	"path"
	"regexp"
	"strings"
)

func HandlerNotifications(w http.ResponseWriter, r *http.Request) {
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

func getWebhook(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	// Check if the given path is valid
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

func getAllWebhooks(w http.ResponseWriter) {

	webhooks, err := webhooks_db.GetAllWebHooks()
	if err != nil {
		if err.Error() == constants.WebhookDBIsEmpty {
			http.Error(w, err.Error(), http.StatusOK)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = utility.EncodeStruct(w, webhooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func getSingleWebhook(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	webhookId := path.Base(cleanPath)

	webhook, err := webhooks_db.GetWebhookById(webhookId)
	if err != nil {
		if err.Error() == constants.WebhookNotFoundError {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = utility.EncodeStruct(w, webhook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Error(w, "", http.StatusOK)
		return
	}

}

func registerWebhook(w http.ResponseWriter, r *http.Request) {
	type responseStruct struct {
		WebhookId string `json:"webhook_id"`
	}

	var webhook structs.Webhook
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&webhook)

	if webhook.Url == "" {
		http.Error(w, "url cannot be empty", http.StatusBadRequest)
		return
	}

	if webhook.Country == "" {
		http.Error(w, "country cannot be empty", http.StatusBadRequest)
		return
	}

	if webhook.Calls < 1 {
		http.Error(w, "calls must be 1 or greater", http.StatusBadRequest)
		return
	}

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

	response := responseStruct{WebhookId: webhookId}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = utility.EncodeStruct(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func deleteWebhook(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	pathList := strings.Split(cleanPath, "/")
	// Check if the given path is valid
	if len(pathList) != 5 {
		http.Error(w, constants.GetBadDeleteWebhookRequestError().Error(), http.StatusBadRequest)
		return
	}

	webhookId := path.Base(cleanPath)

	err := webhooks_db.DeleteWebhook(webhookId)
	if err != nil {
		if err.Error() == constants.WebhookNotFoundError {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "webhook deleted", http.StatusOK)
		return
	}
}
