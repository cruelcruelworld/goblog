package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

type PagesController struct {
}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	var data = view.D{}
	view.Render(w, data, "articles.about")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}