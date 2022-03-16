package api_requests

import (
	"assignment-2/internal/webserver/api_requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoRequestWithWrongMethod(t *testing.T) {
	url := "https://testing.url"
	method := "JUMP"

	_, err := api_requests.DoRequest(url, method)
	assert.NotEqual(t, nil, err)
}
