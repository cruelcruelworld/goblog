package category

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (category *Category) Create() error {
	if err := model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	} else {
		return nil
	}
}