package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

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
	if HasTable(models) {
		return nil
	}
	//env is local
	if env.GetString("ENV") == "local" {
		if err := GetDB().AutoMigrate(models...); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := CreateTable(models...); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
