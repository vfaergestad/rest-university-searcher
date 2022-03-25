package hash_util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashPolicy(country string, scope string) string {
	hash := hmac.New(sha256.New, getHashSecret())
	hash.Write([]byte(country + scope))
	return hex.EncodeToString(hash.Sum(nil))
}
