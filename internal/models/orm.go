package models

import (
	"errors"
	"log"

	"github.com/programzheng/go-auth/config"
	"gorm.io/gorm"
)

var env = config.New()

type DB interface {
	GetDB() *gorm.DB
	SetupTableModel()
}

func GetDB() *gorm.DB {
	switch env.GetString("DB_DIALECTOR") {
	case "mysql":
		return (&MySQL{}).GetDB()
	}

	return nil
}

func HasTable(dst interface{}) bool {
	return GetDB().Migrator().HasTable(dst)
}

func CreateTable(dst ...interface{}) error {
	return GetDB().Migrator().CreateTable(dst...)
}

func SetupTableModel(models ...interface{}) error {
	//env is local
	if config.GetProductionStatus() {
		for _, model := range models {
			if !HasTable(model) {
				if err := CreateTable(model); err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		if err := GetDB().AutoMigrate(models...); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
