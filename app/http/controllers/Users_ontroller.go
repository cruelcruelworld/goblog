package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	BaseController
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_user, err := user.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		articles, err := article.GetByUserID(_user.ID)
		if err != nil {
			uc.ResponseForSQLError(w, err)
		} else {
			data := view.D{
				"Articles": articles,
			}
			view.Render(w, data, "articles.index", "articles._article_meta")
		}
	}
}