package main

import (
	"flag"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t.templates = template.Must(template.ParseGlob("./chat/templates/*.html"))
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	r := newRoom()
	e := echo.New()
	e.Renderer = &Template{}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "chat.html", "")
	})
	e.GET("/room", r.CreateRoom)
	go r.run()

	log.Println("Starting web server on", *addr)
	e.Logger.Fatal(e.Start(*addr))
}
