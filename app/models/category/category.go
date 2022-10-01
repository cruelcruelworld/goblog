package category

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type Category struct {
	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
	models.BaseModel
}

func (category Category) Link() string {
	return route.Name2URL("categories.show", "id", category.GetStringID())
}