package google

import (
	"testing"
)

func TestGetGoogleOauthClientIDSecret(t *testing.T) {
	_, err := getGoogleOauthClientIDSecret()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret error:%v", err)
		return
	}

	t.Logf("success")
}

func TestGetUserInfoByToken(t *testing.T) {
	_, err := GetUserInfoByToken("")
	if err != nil {
		t.Errorf("TestGetUserInfoByToken GetUserInfoByToken error:%v", err)
		return
	}

	t.Logf("success")
}
