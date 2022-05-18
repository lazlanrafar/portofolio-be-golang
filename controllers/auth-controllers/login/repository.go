package loginAuth

import (
	model "lazlanrafar/models"

	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(username string, password string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) LoginRepository(username string, password string) (*model.User, string) {
	var user model.User
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	resultsLogin := db.Debug().Select("*").Where("username = ? AND password = ?", username, password).Find(&user)
	if resultsLogin.Error != nil {
		errorCode <- resultsLogin.Error.Error()
	}else {
		errorCode <- "success"
	}

	return &user, <-errorCode
}