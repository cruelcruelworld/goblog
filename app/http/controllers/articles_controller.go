package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/policies"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type ArticlesController struct {
	BaseController
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Article":          _article,
			"CanModifyArticle": policies.CanModifyArticle(_article),
		}, "articles.show", "articles._article_meta")
	}
}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Articles" : articles,
		}, "articles.index", "articles._article_meta")
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {

	data := view.D{
		"Title": "",
		"Body":  "",
	}

	view.Render(w, data, "articles.create", "articles._form_field")
}

func (ac *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	_article := article.Article{
		Title: title,
		Body:  body,
		UserID: auth.User().ID,
	}

	errors := requests.ValidateArticleForm(_article)
	if len(errors) == 0 {
		_article.Create()
		if _article.ID > 0 {
			http.Redirect(w, r, route.Name2URL("articles.show", "id", _article.GetStringID()), http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		data := view.D{
			"Article": _article,
			"Errors": errors,
		}

		view.Render(w, data, "articles.create", "articles._form_field")
	}
}

func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseUnauthorized(w, r)
		} else {
			data := view.D{
				"Article": _article,
			}
			view.Render(w, data, "articles.edit", "articles._form_field")
		}
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseUnauthorized(w, r)
		} else {
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {

				rs, err := _article.Update()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
					return
				}

				if rs > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何修改！")
				}
			} else {
				data := view.D{
					"Article": _article,
					"Errors":  errors,
				}

				view.Render(w, data, "articles.edit", "articles._form_field")
			}
		}
	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseUnauthorized(w, r)
		} else {
			rowsAffected, err := _article.Delete()

			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				if rowsAffected > 0 {
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}
