package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:3", "max_cn:40"},
		"body": []string{"required", "min_cn:10"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题不能为空",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:内容不能为空",
			"min_cn:内容长度需大于或等于 10 个字节",
		},
	}

	opts := govalidator.Options{
		Data:            &data,
		Rules:           rules,
		Messages:        messages,
		TagIdentifier:   "valid",
	}

	errs := govalidator.New(opts).ValidateStruct()

	return errs
}