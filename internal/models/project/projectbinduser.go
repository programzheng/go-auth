package project

import (
	"github.com/programzheng/go-auth/internal/models"
	"gorm.io/gorm"
)

type ProjectBindUser struct {
	gorm.Model
	UserID           *string `gorm:"not null;uniqueIndex:idx_project_bind_user"`
	ProviderUniqueID *string `gorm:"not null;uniqueIndex:idx_project_bind_user"`
	ProjectID        *uint   `gorm:"not null;uniqueIndex:idx_project_bind_user"`
	Project          Project
}

func init() {
	models.SetupTableModel(&ProjectBindUser{})
}
