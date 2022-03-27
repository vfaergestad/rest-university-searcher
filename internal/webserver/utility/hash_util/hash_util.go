package hash_util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPolicy(country string, scope string) string {
	hash := hmac.New(sha256.New, getHashSecret())
	hash.Write([]byte(country + scope))
	return hex.EncodeToString(hash.Sum(nil))
}

func HashWebhook(url string, country string, calls int) string {
	hash := hmac.New(sha256.New, getHashSecret())
	hash.Write([]byte(fmt.Sprintf("%s%s%d", url, country, calls)))
	return hex.EncodeToString(hash.Sum(nil))
}
