package utils

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	res := GenerateRandomString(32)
	if len(res) != 32 {
		t.Errorf("TestGenerateRandomString length error:%v", res)
	}
	t.Log("success")
}
