package utils

import "testing"

func TestRawURLEncode(t *testing.T) {
	s := "foo bar"
	res := RawURLEncode(s)
	if res != "foo%20bar" {
		t.Errorf("TestRawURLEncode error, value:%v", res)
	}
	t.Log("success")
}
