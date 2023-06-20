package controller

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// IndexController of Index Page
func IndexControler(resWriter http.ResponseWriter, req *http.Request) {
	templates = template.Must(template.ParseGlob("./view/*.html"))
	templates.ExecuteTemplate(resWriter, "index.html", nil)
}
