package main

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),

	//	templates: template.Must(template.ParseFiles( this works, but above does not 
	//		"views/index.html",
	//		"views/blocks/content.html",
	//		"views/blocks/header.html",
	//	)),
	}
}

type Page struct{}

func newPage() Page {
	return Page{}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	page := newPage()
    fmt.Println(filepath.Match("./views/**/*.html","./views/blocks/content.html"));
	e.Renderer = newTemplate()

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
