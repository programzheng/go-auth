package projectservice

import (
	"github.com/programzheng/go-auth/internal/models"
	"github.com/programzheng/go-auth/internal/models/project"
)

type ProjectBindUserService struct {
	Model project.ProjectBindUser
}

func (pbus *ProjectBindUserService) GetFirstModel() error {
	if err := models.GetDB().Where(&pbus.Model).First(&pbus.Model).Error; err != nil {
		return err
	}

	return nil
}

func (pbus *ProjectBindUserService) Create() error {
	if err := models.GetDB().Save(&pbus.Model).Error; err != nil {
		return err
	}

	return nil
}
