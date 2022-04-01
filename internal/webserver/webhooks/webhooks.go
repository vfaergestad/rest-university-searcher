package webhooks

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Invoke(country string) {
	// Checks if given alpha-code is a three letter string.
	match, err := regexp.MatchString(constants.AlphaCodeRegex, country)
	if err != nil {
		fmt.Println(err.Error())
	} else if match {
		country, err = country_cache.GetCountry(country)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = checkAndInvokeWebhooks(country)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func checkAndInvokeWebhooks(country string) error {
	webhooks, err := webhooks_db.GetAllWebHooks()
	if err != nil {
		return err
	}
	for _, w := range webhooks {

		if w.Country == country {
			w.Count = w.Count + 1
			if w.Count >= w.Calls {
				go callWebhook(w)
				w.Count = 0
			}
			_, err = webhooks_db.UpdateWebhook(w.Url, w.Country, w.Calls, w.Count)
		}
	}
	return nil
}

func callWebhook(webhook structs.Webhook) {
	body := webhook
	body.Url = ""
	result, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Attempting invocation of url " + webhook.Url)
	res, err := utility.PostRequest(webhook.Url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err.Error())
	}

	// Read the response
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Something is wrong with invocation response. Error:", err)
		return
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}
