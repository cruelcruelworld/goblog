package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "categories.create")
}

func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request)  {
	name := r.PostFormValue("name")
	var _category category.Category
	_category.Name = name
	errors := requests.ValidateCategoryForm(_category)
	if len(errors) == 0 {
		_category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}

	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors": errors,
		}, "categories.create")
	}
}

func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	_category, _ := category.Get(id)

	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Articles":  articles,
			"PageData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}
