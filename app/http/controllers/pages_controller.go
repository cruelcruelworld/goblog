package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

type PagesController struct {
}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是goblog!</h1>")

}

func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	view.Render(w, data, "articles.about")
}

func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}