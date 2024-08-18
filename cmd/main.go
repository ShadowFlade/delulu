package main

import (
	"delulu/pkg"
	"delulu/pkg/db"
	"fmt"
	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	// TODO make recursive file inclusion - have it in bookmarks somewhere
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct{}

func newPage() Page {
	return Page{}
}

func main() {
	e := echo.New()
	db := db.Db{}
	db.Connect()

	if err := env.Load("./.env"); err != nil {
		fmt.Println("error")
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/static", "static")
	e.File("/toolcool-range-slider.min.js", "node_modules/toolcool-range-slider/dist/toolcool-range-slider.min.js")
	e.File("/tcrs-moving-tooltip.min.js", "node_modules/toolcool-range-slider/dist/plugins/tcrs-moving-tooltip.min.js")
	handlers := pkg.Handlers{}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", struct {
			Page             string
			Header           string
			Age1             int
			Age2             int
			RecaptchaSitekey string
		}{
			Page:             "index",
			Header:           "header",
			Age1:             20,
			Age2:             30,
			RecaptchaSitekey: env.Get("RECAPTCHA_SITEKEY", ""),
		})
	})
	fmt.Println(env.Get("RECAPTCHA_SITEKEY", ""))

	e.GET("/"+pkg.Pages.RESULT, handlers.Result)
	e.GET("/"+pkg.Pages.ABOUT, func(c echo.Context) error {
		c.Render(200, "index", struct{ Page string }{
			Page: pkg.Pages.ABOUT,
		})
		return nil
	})

	e.GET("/"+pkg.Pages.RESOURCES, func(c echo.Context) error {
		c.Render(200, "index", struct{ Page string }{
			Page: pkg.Pages.RESOURCES,
		})
		return nil
	})

	ipPort := env.Get("IP_PORT", ":3000")
	e.POST("/"+pkg.Pages.FEEDBACK, handlers.Feedback)
	e.POST("/captcha_check", handlers.CaptchaCheck)
	e.Logger.Fatal(e.Start(ipPort))
}
