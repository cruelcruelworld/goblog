package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func Render(w http.ResponseWriter, data interface{}, tplfiles ...string)  {
	viewDir := "resources/views/"

	for i, f := range tplfiles {
		tplfiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.LogError(err)
	newFiles := append(files, tplfiles...)
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}