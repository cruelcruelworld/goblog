package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body string	`gorm:"type:longtext;not null;" valid:"body"`
	models.BaseModel
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}