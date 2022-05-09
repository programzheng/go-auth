package twitter

import (
	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/internal/providers/twitter"
)

type RequestTokenRequest struct {
	OauthCallback   string `json:"oauth_callback"`
	XAuthAccessType string `json:"x_auth_access_type"`
}

func RequestToken(c *gin.Context) {
	request := RequestTokenRequest{}
	c.ShouldBindJSON(&request)
	requestMap := map[string]interface{}{
		"oauth_callback":     "https://programzheng-projects.tk",
		"x_auth_access_type": request.XAuthAccessType,
	}
	twitter.TwitterRequestToken(requestMap)
}
