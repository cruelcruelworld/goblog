package category

import "goblog/app/models"

type Category struct {
	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
	models.BaseModel
}