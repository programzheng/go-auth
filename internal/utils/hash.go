package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func HashHamcSha1HEX(s string, key string) string {
	k := []byte(key)
	mac := hmac.New(sha1.New, k)
	mac.Write([]byte(s))
	res := fmt.Sprintf("%x", mac.Sum(nil))
	return res
}

func HashHamcSha1Base64(s string, key string) string {
	k := []byte(key)
	mac := hmac.New(sha1.New, k)
	mac.Write([]byte(s))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}