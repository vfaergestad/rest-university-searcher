package notifications_handler

import (
	"assignment-2/internal/webserver/cache/country_cache"
	"assignment-2/internal/webserver/constants"
	"assignment-2/internal/webserver/db"
	"assignment-2/internal/webserver/db/webhooks_db"
	"assignment-2/internal/webserver/mock_apis"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/utility/hash_util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var notificationsEndpoint *httptest.Server
var webhooksIds []string

func TestMain(m *testing.M) {
	countriesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCountries))
	defer countriesMock.Close()

	constants.SetTestCountryAPIUrl(countriesMock.URL)
	constants.SetTestServiceAccountLocation()

	err := db.InitializeFirestore()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = db.CloseFirestore()
		if err != nil {
			panic(err)
		}
	}()

	webhooksIds = webhooks_db.SetUpTestDatabase()

	err = country_cache.InitCache()
	if err != nil {
		panic(err)
	}

	notificationsEndpoint = httptest.NewServer(http.HandlerFunc(HandlerNotifications))
	defer notificationsEndpoint.Close()

	m.Run()
}

func TestGetDeleteWebhook(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   structs.Webhook
	}{
		{
			name: "Valid Get Single Webhook",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/" + webhooksIds[0],
			},
			method:             "GET",
			expectedStatusCode: http.StatusOK,
			expectedResponse: structs.Webhook{
				WebhookId: webhooksIds[0],
				Url:       "https://example.com",
				Country:   "Norway",
				Calls:     1,
			},
		},
		{
			name: "Invalid Get Single Webhook",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/invalid-id",
			},
			method:             "GET",
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid To Long Path",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/invalid-id/hei",
			},
			method:             "GET",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Valid Delete Webhook",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/" + webhooksIds[0],
			},
			method:             "DELETE",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Delete Webhook",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/invalid-id",
			},
			method:             "DELETE",
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Method",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/invalid-id",
			},
			method:             http.MethodHead,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   structs.Webhook{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Errorf("Error making %s request: %s", tt.method, err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}

			if tt.expectedResponse != (structs.Webhook{}) {
				var actual structs.Webhook
				_ = json.NewDecoder(res.Body).Decode(&actual)

				if actual != tt.expectedResponse {
					t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
				}
			}
		})
	}
	t.Cleanup(func() {
		webhooks_db.SetUpTestDatabase()
	})
}

func TestGetAllWebhooks(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		method             string
		expectedStatusCode int
		expectedResponse   []structs.Webhook
	}{
		{
			name: "Valid Get All Webhooks",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			method:             "GET",
			expectedStatusCode: http.StatusOK,
			expectedResponse: []structs.Webhook{
				{
					WebhookId: webhooksIds[0],
					Url:       "https://example.com",
					Country:   "Norway",
					Calls:     1,
				},
				{
					WebhookId: webhooksIds[1],
					Url:       "https://example2.com",
					Country:   "Sweden",
					Calls:     2,
				},
				{
					WebhookId: webhooksIds[2],
					Url:       "https://example3.com",
					Country:   "Denmark",
					Calls:     3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(tt.method, tt.args.url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Errorf("Error making %s request: %s", tt.method, err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}

			var actual []structs.Webhook
			_ = json.NewDecoder(res.Body).Decode(&actual)

			if len(actual) != len(tt.expectedResponse) {
				t.Errorf("Expected %d webhooks, got %d", len(tt.expectedResponse), len(actual))
			}
		})
	}
}

func TestRegisterWebhook(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name               string
		args               args
		body               structs.Webhook
		expectedStatusCode int
		expectedResponse   structs.Webhook
	}{
		{
			name: "Valid Register Webhook",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "Finland",
				Url:     "https://example4.com",
				Calls:   4,
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse: structs.Webhook{
				WebhookId: hash_util.HashWebhook("https://example4.com", "Finland", 4),
			},
		},
		{
			name: "Invalid Register Webhook; Empty url",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "Norway",
				Url:     "",
				Calls:   1,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Register Webhook; Empty country",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "",
				Url:     "https://example.com",
				Calls:   2,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Register Webhook; Empty calls",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "France",
				Url:     "https://example.com",
				Calls:   0,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Register Webhook; Invalid alpha-code",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "TES",
				Url:     "https://example.com",
				Calls:   2,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
		{
			name: "Invalid Register Webhook; Already exists",
			args: args{
				url: notificationsEndpoint.URL + "/corona/v1/notifications/",
			},
			body: structs.Webhook{
				Country: "Norway",
				Url:     "https://example.com",
				Calls:   1,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   structs.Webhook{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			body, _ := json.Marshal(tt.body)
			req, err := http.NewRequest(http.MethodPost, tt.args.url, strings.NewReader(string(body)))
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Errorf("Error making POST request: %s", err)
			}
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("Expected %d response, got %d", tt.expectedStatusCode, res.StatusCode)
			}

			if tt.expectedResponse != (structs.Webhook{}) {
				var actual structs.Webhook
				_ = json.NewDecoder(res.Body).Decode(&actual)

				if actual != tt.expectedResponse {
					t.Errorf("Expected %v, got %v", tt.expectedResponse, actual)
				}
			}
		})
	}

	t.Cleanup(func() {
		webhooks_db.SetUpTestDatabase()
	})
}
