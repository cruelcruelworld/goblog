package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/category"
)

func ValidateCategoryForm(category category.Category) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
		},
	}

	opts := govalidator.Options{
		Data:            &category,
		Rules:           rules,
		Messages:        messages,
		TagIdentifier:   "valid",
	}

	return govalidator.New(opts).ValidateStruct()

}