package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
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

	router.Use(forceHTMLMiddleware)


	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
