package main
import (
    "html/template"
    "io--primary: #005b99;
--important:#ff7f00; 
--bg:#f7f7f7;
--border:#d3d3d3; 
--stand-out:#ffcc00;
"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpls *template.Template
}

func (t *Template) Render (w io.Writer, name string, data interface{}, c echo.Context) error  {
    return t.tmpls.ExecuteTemplate( w, name , data )
}



