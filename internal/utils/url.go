package utils

import (
	"net/url"
	"strings"
)

func RawURLEncode(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}
