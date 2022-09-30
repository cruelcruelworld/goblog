package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title, Body string
	ID          uint64
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}