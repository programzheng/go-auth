package twitter

import (
	"testing"
)

// func TestTwitterRequestToken(t *testing.T) {
// 	requestTokenRequest := map[string]interface{}{
// 		"oauth_callback":     "foo",
// 		"x_auth_access_type": "bar",
// 	}
// 	res, err := TwitterRequestToken(requestTokenRequest)
// 	if err != nil {
// 		t.Errorf("TestTwitterRequestToken http post error:%v", err)
// 		t.Errorf("TestTwitterRequestToken http post result:%v", res)
// 		return
// 	}
// 	_, exist := res["oauth_token"]
// 	if !exist {
// 		t.Errorf("TestTwitterRequestToken oauth_token is not exist:%v", err)
// 		return
// 	}
// 	_, exist = res["oauth_token_secret"]
// 	if !exist {
// 		t.Errorf("TestTwitterRequestToken oauth_token is not exist:%v", err)
// 		return
// 	}
// 	_, exist = res["oauth_callback_confirmed"]
// 	if !exist {
// 		t.Errorf("TestTwitterRequestToken oauth_token is not exist:%v", err)
// 		return
// 	}
// 	t.Log("success")
// }

func getTestOauthParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["include_entities"] = "true"
	params["oauth_consumer_key"] = "xvz1evFS4wEEPTGEFPHBog"
	params["oauth_nonce"] = "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = "1318622958"
	params["oauth_token"] = "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"
	params["oauth_version"] = "1.0"

	return params
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Collecting%20parameters
func TestGenerateCollectingParameters(t *testing.T) {
	params := getTestOauthParams()
	res := generateCollectingParameters(params, map[string]interface{}{
		"status": "Hello Ladies + Gentlemen, a signed OAuth request!",
	})
	if res != "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21" {
		t.Errorf("TestGenerateCollectingParameters error, result:%v", res)
		return
	}
	t.Log("success")
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Creating%20the%20signature%20base%20string
func TestGenerateSignatureBase(t *testing.T) {
	params := getTestOauthParams()
	query := generateCollectingParameters(params, map[string]interface{}{
		"status": "Hello Ladies + Gentlemen, a signed OAuth request!",
	})
	res := generateSignatureBase("https://api.twitter.com/1.1/statuses/update.json", "POST", query)
	if res != "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521" {
		t.Errorf("TestGenerateSignatureBase error, result:%v", res)
		return
	}
	t.Log("success")
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Getting%20a%20signing%20key
func TestGenerateSigningKey(t *testing.T) {
	consumerSecret := "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw"
	oauthTokenSecret := "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"
	res := generateSigningKey(consumerSecret, oauthTokenSecret)
	if res != "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw&LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE" {
		t.Errorf("TestGenerateSigningKey error, result:%v", res)
		return
	}
	t.Log("success")
}

//https://developer.twitter.com/en/docs/authentication/oauth-1-0a/creating-a-signature#:~:text=Calculating%20the%20signature
func TestGenerateTwitterOauthSignature(t *testing.T) {
	params := getTestOauthParams()
	query := generateCollectingParameters(params, map[string]interface{}{
		"status": "Hello Ladies + Gentlemen, a signed OAuth request!",
	})
	signatureBase := generateSignatureBase("https://api.twitter.com/1.1/statuses/update.json", "POST", query)
	signingKey := generateSigningKey("kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw", "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE")
	res := generateTwitterOauthSignature(signatureBase, signingKey)
	if res != "hCtSmYh+iHYCEqBWrE7C7hYmtUk=" {
		t.Errorf("TestGenerateTwitterOauthSignature error, result:%v", res)
		return
	}
	t.Log("success")
}
