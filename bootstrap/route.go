package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

var router *mux.Router

func SetupRoute() *mux.Router {
	router = mux.NewRouter()
	routes.RegisterWebRoutes(router)

	return router
}
