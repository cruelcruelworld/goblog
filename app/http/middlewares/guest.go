package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

func Guest(next HttpHandlerFunc) HttpHandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if auth.Check() {
			flash.Warning("登录用户无法访问此页面")
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}

		next(writer, request)
	}
}
