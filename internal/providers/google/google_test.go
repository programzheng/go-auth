package google

import (
	"net/url"
	"strings"
	"testing"
)

func TestGetGoogleOauthClientIDSecret(t *testing.T) {
	_, err := getGoogleOauthConfigFromClientIDSecretJSON()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret error:%v", err)
		return
	}

	t.Logf("success")
}

func TestGetGoogleOauthURL(t *testing.T) {
	config, err := getGoogleOauthConfigFromClientIDSecretJSON()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret getGoogleOauthConfigFromClientIDSecretJSON error:%v", err)
		return
	}

	state, err := GetGoogleOauthState()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret GetGoogleOauthState error:%v", err)
		return
	}

	oAuthURL, err := GetGoogleOauthURL()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret GetGoogleOauthURL error:%v", err)
		return
	}
	redirectURI := url.QueryEscape(config.RedirectURL)

	scope := url.QueryEscape(strings.Join(config.Scopes, " "))

	target := "https://accounts.google.com/o/oauth2/auth?client_id=" + config.ClientID + "&redirect_uri=" + redirectURI + "&response_type=code&scope=" + scope + "&state=" + state
	if oAuthURL != target {
		t.Errorf("TestGetGoogleOauthURL oAuthURL\ncurrent:%v\ntarget:%v", oAuthURL, target)
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
