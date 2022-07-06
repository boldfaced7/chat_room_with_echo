package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./chat/templates/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "chat.html", "")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
