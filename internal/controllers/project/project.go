package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/internal/controllers"
	"github.com/programzheng/go-auth/internal/models"
	"github.com/programzheng/go-auth/internal/models/project"
	"github.com/programzheng/go-auth/internal/resources"
	"github.com/programzheng/go-auth/internal/services/projectservice"
)

type CreateProjectRequest struct {
	Provider    *string `json:"provider"`
	ProjectName *string `json:"project_name"`
}

func CreateProject(c *gin.Context) {
	request := CreateProjectRequest{}
	controllers.GinBind(c, &request)

	key := projectservice.GenerateKey()
	p := project.Project{
		Provider:    request.Provider,
		ProjectName: request.ProjectName,
		Key:         key,
	}

	if err := models.GetDB().Save(&p).Error; err != nil {
		c.JSON(http.StatusUnauthorized, resources.GlobalResponse("error", err))
		return
	}

	c.JSON(http.StatusOK, resources.GlobalResponse("success", resources.H{
		"project": &p,
	}))
}
