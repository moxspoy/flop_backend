package users_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func UpdatePhoneNumber(user *database_model.Users) *gorm.DB {
	return database.DB.Save(&user)
}
