package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "between:3,40"},
		"body": []string{"required", "min:10"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题不能为空",
			"between:标题长度需介于 3-40",
		},
		"body": []string{
			"required:内容不能为空",
			"min:内容长度需大于或等于 10 个字节",
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