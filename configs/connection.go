package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	baseUrl := "root:@tcp(127.0.0.1:3306)/lazlanrafar?charset=utf8mb4&parseTime=True&loc=Local"

  	db, err := gorm.Open(mysql.Open(baseUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
		  
	return db
}