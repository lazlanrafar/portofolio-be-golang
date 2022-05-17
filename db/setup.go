package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB()  {
	dsn := "root:@tcp(127.0.0.1:3306)/lazlanrafar?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Migrate(db)

}