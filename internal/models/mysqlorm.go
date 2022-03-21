package models

import (
	"fmt"
	"log"

	"github.com/programzheng/go-auth/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
}

var env = config.New()

var (
	globalDB *gorm.DB
)

func init() {
	var err error
	//?parseTime=true for the database table column type is TIMESTAMP
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&loc=Local&parseTime=true",
		env.GetString("DB_USERNAME"),
		env.GetString("DB_PASSWORD"),
		env.GetString("DB_HOST"),
		env.GetString("DB_PORT"),
		env.GetString("DB_DATABASE"))
	fmt.Printf("connect: %v database\n", dsn)
	globalDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		log.Println("DataBase error:", err)
	}
}

func (mq *MySQL) GetDB() *gorm.DB {
	return globalDB
}
