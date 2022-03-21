package services

import "github.com/programzheng/go-auth/internal/models"

func IsErrRecordNotFound(err error) bool {
	return models.IsErrRecordNotFound(err)
}
