package google

import (
	"net/url"
	"strings"
	"testing"
)

func TestGetGoogleOauthClientIDSecret(t *testing.T) {
	_, err := getGoogleConfigFromClientIDSecretJSON()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret error:%v", err)
		return
	}

	t.Log("success")
}

func TestGetGoogleOauthURL(t *testing.T) {
	config, err := getGoogleConfigFromClientIDSecretJSON()
	if err != nil {
		t.Errorf("TestGetGoogleOauthClientIDSecret getGoogleConfigFromClientIDSecretJSON error:%v", err)
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

	t.Log("success")
}

func TestGetUserInfoByToken(t *testing.T) {
	_, err := GetUserInfoByToken("")
	if err != nil {
		t.Errorf("TestGetUserInfoByToken GetUserInfoByToken error:%v", err)
		return
	}

	t.Log("success")
}

func TestValidateGoogleOauthIDToken(t *testing.T) {
	payload, err := ValidateGoogleOauthIDToken("")
	if err != nil {
		t.Errorf("TestValidateGoogleOauthIDToken ValidateGoogleOauthIDToken error:%v", err)
		return
	}
	clientID, err := getGoogleClientIDSecretFromJSON()
	if err != nil {
		t.Errorf("TestValidateGoogleOauthIDToken getGoogleClientIDSecretFromJSON error:%v", err)
		return
	}
	if payload.Audience != clientID {
		t.Errorf("TestValidateGoogleOauthIDToken payload.Audience:%v != clientID:%v error", payload.Audience, clientID)
		return
	}

	t.Log("success")
}
