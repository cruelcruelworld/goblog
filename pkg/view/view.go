package view

import (
	"goblog/pkg/auth"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplfiles ...string)  {
	RenderTemplate(w, "app", data, tplfiles...)
}

func RenderSimple(w io.Writer, data D, tplfiles ...string)  {
	RenderTemplate(w, "simple", data, tplfiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplfiles ...string)  {
	data["isLogined"] = auth.Check()

	newFiles := getTemplateFiles(tplfiles...)

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplfiles ...string) []string {
	viewDir := "resources/views/"

	for i, f := range tplfiles {
		tplfiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.LogError(err)

	return append(files, tplfiles...)
}