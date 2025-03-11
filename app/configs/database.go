package configs

import (
	"admin-v1/app/helpers"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormConnection() *gorm.DB {
	var db *gorm.DB

	if helpers.GormDB != nil {
		return helpers.GormDB
	}

	var (
		devHostName = os.Getenv("MYSQL_HOST")
		devDbName   = os.Getenv("MYSQL_DB_NAME")
		devUser     = os.Getenv("MYSQL_USER")
		devPassword = os.Getenv("MYSQL_PASSWORD")
		devPort     = os.Getenv("MYSQL_PORT")
	)

	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", devUser, devPassword, devHostName, devPort, devDbName)+"?parseTime=true&charset=utf8mb4&loc=Local"),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	return db
}