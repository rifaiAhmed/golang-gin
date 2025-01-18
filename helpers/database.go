package helpers

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetupMySql() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER", ""), GetEnv("DB_PASS", ""), GetEnv("DB_HOST", "localhost"), GetEnv("DB_PORT", "8080"), GetEnv("DB_NAME", ""),
	)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect databse", err)
	}
	logrus.Print("succesfully connect to databse")
}
