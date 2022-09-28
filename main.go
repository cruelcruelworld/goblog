package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello, 这里是goblog!</h1>")

}

func aboutHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,  "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:363201375@qq.com\">363201375@qq.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID：" + id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "文章列表")
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "创建文章")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesCreateHandler).Methods("POST").Name("articles.store")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	homeURL,_ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL,_ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL: ", articleURL)


	http.ListenAndServe(":3000", router)
}
