package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"net/http"
	"strconv"
	"strings"
)

var router *mux.Router
var db *sql.DB

type Article struct {
	Title, Body string
	ID          int64
}

func (article Article) Delete() (int64, error) {
	rs, err := db.Exec("DELETE FROM articles WHERE id =" + strconv.FormatInt(article.ID, 10))
	if err != nil {
		return 0, err
	}

	if id, err := rs.RowsAffected(); id > 0 {
		return id, err
	}

	return 0, nil
}

func (article Article) Link() string {
	showURL, err := router.Get("articles.show").URL("id", strconv.FormatInt(article.ID, 10))
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return showURL.String()
}

func getArticleByID (id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)

	return article, err
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		rowsAffected, err := article.Delete()

		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			if rowsAffected > 0 {
				indexURL, _ := router.Get("articles.index").URL()
				http.Redirect(w, r, indexURL.String(), http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章未找到")
			}
		}
	}
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

	router.Use(forceHTMLMiddleware)


	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
