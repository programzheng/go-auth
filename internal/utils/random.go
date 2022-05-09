package utils

import (
	"bytes"
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const latin = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01233456789"
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
	}

	return buffer.String()
}
