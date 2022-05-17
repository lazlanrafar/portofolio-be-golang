package config

import (
	model "lazlanrafar/models"
	util "lazlanrafar/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

func Connection() *gorm.DB {
	baseUrl := util.GodotEnv("DATABASE_URL")
  	db, err := gorm.Open(mysql.Open(baseUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.EntityWork{})
		  
	return db
}