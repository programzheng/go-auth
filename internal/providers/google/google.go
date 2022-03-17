package google

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/programzheng/go-auth/config"
)

var env = config.New()

type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func getGoogleOauthConfigFromClientIDSecretJSON() (*oauth2.Config, error) {
	clientIDSecret := env.GetString("GOOGLE_OAUTH2_CLIENT_ID_SECRET")
	clientIDSecretByte := []byte(clientIDSecret)
	conf, err := google.ConfigFromJSON(clientIDSecretByte, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile")
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func GetGoogleOauthRedirectURL() (string, error) {
	redirectURL := env.GetString("GOOGLE_OAUTH2_CALLBACK")
	if redirectURL == "" {
		return "", errors.New("getGoogleOauthURL redirectURL is empty string")
	}

	return redirectURL, nil
}

func GetGoogleOauthState() (string, error) {
	state := env.GetString("GOOGLE_OAUTH2_STATE")
	if state == "" {
		return "", errors.New("getGoogleOauthURL state is empty string")
	}

	return state, nil
}

func GetGoogleOauthURL() (string, error) {
	state, err := GetGoogleOauthState()
	if err != nil {
		return "", err
	}

	config, err := getGoogleOauthConfigFromClientIDSecretJSON()
	if err != nil {
		return "", err
	}

	authURL := config.AuthCodeURL(state)

	return authURL, nil
}

func IsValidGoogleOauthState(state string) error {
	originState := env.GetString("GOOGLE_OAUTH2_STATE")
	if originState == "" {
		return errors.New("isValidGoogleOauthState state is empty string")
	}

	if originState != state {
		return errors.New("isValidGoogleOauthState state is valid fail")
	}

	return nil
}

func GetGoogleOauthTokenByCode(code string) (*oauth2.Token, error) {
	config, err := getGoogleOauthConfigFromClientIDSecretJSON()
	if err != nil {
		return nil, err
	}

	token, err := config.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserInfoByToken(token string) (*GoogleUserInfo, error) {
	config, err := getGoogleOauthConfigFromClientIDSecretJSON()
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{
		AccessToken: token,
	}
	client := config.Client(context.TODO(), t)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read response to []byte
	rawData, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Printf("GetUserInfoByToken error:%v", string(rawData))
		return nil, errors.New("get user info fail")
	}

	var googleUserInfo GoogleUserInfo
	json.Unmarshal(rawData, &googleUserInfo)

	return &googleUserInfo, nil
}
