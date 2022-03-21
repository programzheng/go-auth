package projectservice

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/programzheng/go-auth/internal/models"
	"github.com/programzheng/go-auth/internal/models/project"
)

type ProjectService struct {
	Model project.Project
}

func GerenateKey() string {
	t := time.Now()
	ts := t.String()

	hash := md5.Sum([]byte(ts))
	return hex.EncodeToString(hash[:])
}

func (ps *ProjectService) GetFirstModel() error {
	if err := models.GetDB().Where(&ps.Model).First(&ps.Model).Error; err != nil {
		return err
	}

	return nil
}
