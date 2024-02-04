package main
import (
    "html/template"
    "io"
    "net/mail"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpls *template.Template
}

func (t *Template) Render (w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.tmpls.ExecuteTemplate( w, name , data )
}



