package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/config"
	"github.com/programzheng/go-auth/internal/controllers/google"
	"github.com/programzheng/go-auth/internal/controllers/project"
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
	googleGroup.POST("get_oauth_url", google.GetOauthURL)
	googleGroup.POST("get_user_info", google.GetUserInfoByToken)
	googleGroup.POST("get_oauth_unique_id", google.GetGoogleOauthUniqueIDByIDToken)
	googleGroup.POST("projects_oauth_login", google.GoogleProjectOauthLogin)

	port := config.New().GetString("PORT")

	if port != "" {
		r.Run(":" + port)
		return
	}

	r.Run()
}
