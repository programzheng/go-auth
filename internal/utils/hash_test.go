package utils

import (
	"testing"
)

func TestHashHamcSha1HEX(t *testing.T) {
	s := "secret"
	k := "key"
	res := HashHamcSha1HEX(s, k)
	if res != "bb1e45ed87beae9c2fcac6f35012eefb019a6170" {
		t.Errorf("TestHashHamcSha1HEX error, result:%v", res)
	}
	t.Log("success")
}

func TestHashHamcSha1Base64(t *testing.T) {
	s := "secret"
	k := "key"
	res := HashHamcSha1Base64(s, k)
	if res != "ux5F7Ye+rpwvysbzUBLu+wGaYXA=" {
		t.Errorf("TestHashHamcSha1Base64 error, result:%v", res)
	}
	t.Log("success")
}
