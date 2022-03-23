package google

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/programzheng/go-auth/internal/controllers"
	"github.com/programzheng/go-auth/internal/providers/google"
	"github.com/programzheng/go-auth/internal/resources"
	"github.com/programzheng/go-auth/internal/services"
	"github.com/programzheng/go-auth/internal/services/projectservice"
)

type GetGoogleOauthUniqueIDByIDTokenRequest struct {
	IDToken string `json:"id_token"`
}

type GoogleProjectOauthLoginRequest struct {
	Provider    *string `json:"provider"`
	ProjectName *string `json:"project_name"`
	Key         string  `json:"key"`
	IDToken     string  `json:"id_token"`
	UserID      *string `json:"user_id"`
}

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
		"status": "success",
		"url":    url,
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
		"status":        "success",
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
		"status":    "success",
		"user_info": userInfo,
	})
}

func GetGoogleOauthUniqueIDByIDToken(c *gin.Context) {
	request := GetGoogleOauthUniqueIDByIDTokenRequest{}
	controllers.GinBind(c, &request)

	payload, err := google.ValidateGoogleOauthIDToken(request.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
		return
	}

	c.JSON(http.StatusOK, resources.GlobalResponse("success", resources.H{
		"unique_id": payload.Subject,
		"claims":    payload.Claims,
	}))
}

func GoogleProjectOauthLogin(c *gin.Context) {
	request := GoogleProjectOauthLoginRequest{}
	controllers.GinBind(c, &request)

	ps := &projectservice.ProjectService{}
	ps.Model.Provider = request.Provider
	ps.Model.ProjectName = request.ProjectName
	ps.Model.Key = request.Key
	if err := ps.GetFirstModel(); err != nil {
		c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
		return
	}

	payload, err := google.ValidateGoogleOauthIDToken(request.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
		return
	}

	pbus := &projectservice.ProjectBindUserService{}

	pbus.Model.UserID = request.UserID
	pbus.Model.ProviderUniqueID = &payload.Subject
	pbus.Model.ProjectID = &ps.Model.ID

	if err := pbus.GetFirstModel(); err != nil {
		if services.IsErrRecordNotFound(err) {
			if err := pbus.Create(); err != nil {
				c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
			return
		}
	}

	c.JSON(http.StatusOK, resources.GlobalResponse("success", resources.H{
		"user_id":   pbus.Model.UserID,
		"unique_id": pbus.Model.ProviderUniqueID,
	}))
}
