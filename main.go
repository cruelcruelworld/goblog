package main

import (
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"
)

var router *mux.Router

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
