package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/config"
	"github.com/programzheng/go-auth/internal/controllers/google"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), ".")
)

func main() {
	r := gin.Default()

	r.GET("google_login_callback", google.GetOauthTokenByCode)

	apiGroup := r.Group("/api/v1")
	googleGroup := apiGroup.Group("/google")
	googleGroup.POST("get_oauth_url", google.GetOauthURL)
	googleGroup.POST("get_user_info", google.GetUserInfoByToken)

	port := config.New().GetString("PORT")

	if port != "" {
		r.Run(":" + port)
		return
	}

	r.Run()
}
