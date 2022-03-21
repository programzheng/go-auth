package project

import (
	"github.com/programzheng/go-auth/internal/models"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Provider    *string `gorm:"not null;uniqueIndex:idx_project" json:"provider"`
	ProjectName *string `gorm:"not null;uniqueIndex:idx_project" json:"project_name"`
	Key         string
}

func init() {
	models.SetupTableModel(&Project{})
}
