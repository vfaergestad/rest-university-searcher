package request

import (
	"assignment-2/internal/webserver/mock_apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testUrl string

func TestMain(m *testing.M) {
	countriesMock := httptest.NewServer(http.HandlerFunc(mock_apis.HandlerCountries))
	defer countriesMock.Close()
	testUrl = countriesMock.URL

	m.Run()
}

func TestRequest(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid request",
			args: args{
				url: testUrl,
			},
			wantErr: false,
		},
		{
			name: "Invalid request",
			args: args{
				url: "g4slg2ns460``tp*øt3",
			},
			wantErr: true,
		},
		{
			name: "Empty request",
			args: args{
				url: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetRequest(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_, err = HeadRequest(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("HeadRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_, err = PostRequest(tt.args.url, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
