package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/password"
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

func All() ([]User, error) {
	var users []User
	if err := model.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) ComparePassword(pwd string) bool {
	return password.CheckHash(pwd, user.Password)
}