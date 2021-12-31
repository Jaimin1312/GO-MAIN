package common

import (
	"go/build"
	"html/template"
)

var tpl *template.Template

//GetTemplate is return template pointer
func GetTemplate() *template.Template {
	path := build.Default.GOPATH + "/src/packages/template/*"
	tpl = template.Must(template.ParseGlob(path))
	return tpl
}
