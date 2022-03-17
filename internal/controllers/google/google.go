package google

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/programzheng/go-auth/internal/providers/google"
)

type GetUserInfoByTokenRequest struct {
	Token string
}

func GetOauthURL(c *gin.Context) {
	url, err := google.GetGoogleOauthURL()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func GetOauthTokenByCode(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	err := google.IsValidGoogleOauthState(state)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	t, err := google.GetGoogleOauthTokenByCode(code)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  t.AccessToken,
		"type":          t.Type(),
		"refresh_token": t.RefreshToken,
		"expiry":        t.Expiry.Unix(),
	})
}

func GetUserInfoByToken(c *gin.Context) {
	request := GetUserInfoByTokenRequest{}
	c.BindJSON(&request)

	userInfo, err := google.GetUserInfoByToken(request.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
	})
}
