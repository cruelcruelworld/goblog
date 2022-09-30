package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (user *User) Create() error {
	if err := model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
