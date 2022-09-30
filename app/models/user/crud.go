package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func (user *User) Create() error {
	if err := model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(uid string) (User, error) {
	id := types.StringToUint64(uid)
	var user User
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) ComparePassword(password string) bool {
	return user.Password == password
}