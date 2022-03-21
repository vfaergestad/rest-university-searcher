package api_requests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoRequestWithWrongMethod(t *testing.T) {
	url := "https://testing.url"
	method := "JUMP"

	_, err := GetRequest(url, method)
	assert.NotEqual(t, nil, err)
}
