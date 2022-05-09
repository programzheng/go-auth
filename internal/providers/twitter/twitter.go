package twitter

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/programzheng/go-auth/config"
	"github.com/programzheng/go-auth/internal/utils"
)

var env = config.New()

type RequestTokenResponse struct {
	OauthToken             string `json:"oauth_token"`
	OauthTokenSecret       string `json:"oauth_token_secret"`
	OauthCallbackConfirmed bool   `json:"oauth_callback_confirmed"`
}

func TwitterRequestToken(request map[string]interface{}) (map[string]interface{}, error) {
	jsnoData, err := json.Marshal(request)
	if err != nil {
		log.Printf("twitter provider RequestToken json marshal error:%v", err)
	}
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth/request_token", bytes.NewBuffer(jsnoData))
	if err != nil {
		log.Printf("twitter provider RequestToken new http post request error:%v", err)
	}
	setTwitterRequestTokenHeaders(req, request)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Printf("twitter provider RequestToken http post request error:%v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t := make(map[string]interface{})
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("twitter provider RequestToken StatusCode error:%v", err)
		}
		err = json.Unmarshal(b, &t)
		if err != nil {
			log.Printf("%v", err)
		}
		return t, errors.New("unknow error")
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

func setTwitterRequestTokenHeaders(req *http.Request, request map[string]interface{}) {
	headers := make(map[string]interface{})
	headers["oauth_consumer_key"] = env.GetString("TWITTER_CONSUMER_KEY")
	headers["oauth_nonce"] = utils.GenerateRandomString(32)
	headers["oauth_signature_method"] = "HMAC-SHA1"
	headers["oauth_timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	headers["oauth_version"] = env.GetString("TWITTER_OAUTH_VERSION")
	query := generateCollectingParameters(headers, request)
	oauthSignatureBase := generateSignatureBase(req.URL.String(), req.Method, query)
	signingKey := generateSigningKey(
		env.GetString("TWITTER_CONSUMER_KEY_SECRET"),
		env.GetString("TWITTER_OAUTH_CLIENT_SECRET"),
	)
	headers["oauth_signature"] = generateTwitterOauthSignature(
		oauthSignatureBase,
		signingKey,
	)

	req.Header.Set("oauth_consumer_key", headers["oauth_consumer_key"].(string))
	req.Header.Set("oauth_nonce", headers["oauth_nonce"].(string))
	req.Header.Set("oauth_signature_method", headers["oauth_signature_method"].(string))
	req.Header.Set("oauth_timestamp", headers["oauth_timestamp"].(string))
	req.Header.Set("oauth_version", headers["oauth_version"].(string))
	req.Header.Set("oauth_signature", headers["oauth_signature"].(string))
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Collecting%20parameters
func generateCollectingParameters(headers map[string]interface{}, request map[string]interface{}) string {
	for k, v := range headers {
		request[k] = v
	}
	params := request
	var buf strings.Builder
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := params[k].(string)
		keyEscaped := utils.RawURLEncode(k)
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(utils.RawURLEncode(vs))
	}
	return buf.String()
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Creating%20the%20signature%20base%20string
func generateSignatureBase(URL string, method string, query string) string {
	return method + "&" + url.QueryEscape(URL) + "&" + url.QueryEscape(query)
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Getting%20a%20signing%20key
func generateSigningKey(consumerSecret string, oauthTokenSecret string) string {
	return consumerSecret + "&" + oauthTokenSecret
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Calculating%20the%20signature
func generateTwitterOauthSignature(signatureBase string, signingKey string) string {
	s := signatureBase
	k := signingKey
	return utils.HashHamcSha1Base64(s, k)
}
