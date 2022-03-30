package default_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultEndpoint(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args
		method             string
		expectedStatusCode int
	}{
		{
			name: "TestDefaultEndpoint",
			args: args{
				url: "/corona/",
			},
			method:             http.MethodGet,
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.args.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			// TODO: HVA ER BEST, DETTE, ELLER SÅNN JEG HAR GJORT ELLERS
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandlerDefault)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expectedStatusCode {
				t.Errorf("defaulthandler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
		})
	}
}
