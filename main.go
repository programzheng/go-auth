package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/config"
	"github.com/programzheng/go-auth/internal/controllers/google"
	"github.com/programzheng/go-auth/internal/controllers/project"
	"github.com/programzheng/go-auth/internal/controllers/twitter"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), ".")
)

func main() {
	r := gin.Default()

	r.GET("google_login_callback", google.GetOauthTokenByCode)

	apiGroup := r.Group("/api/v1")

	projectGroup := apiGroup.Group("/projects")
	projectGroup.POST("", project.CreateProject)

	googleGroup := apiGroup.Group("/google")
	{
		googleGroup.POST("get_oauth_url", google.GetOauthURL)
		googleGroup.POST("get_user_info", google.GetUserInfoByToken)
		googleGroup.POST("get_oauth_unique_id", google.GetGoogleOauthUniqueIDByIDToken)
		googleGroup.POST("projects_oauth_login", google.GoogleProjectOauthLogin)
	}

	twitterGroup := apiGroup.Group("/twitter")
	twitterOauthGroup := twitterGroup.Group("/oauth")
	{
		twitterOauthGroup.POST("request_token", twitter.RequestToken)
	}

	port := config.New().GetString("PORT")

	if port != "" {
		r.Run(":" + port)
		return
	}

	r.Run()
}
