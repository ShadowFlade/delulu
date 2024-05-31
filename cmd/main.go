package main

import (
	"delulu/pkg/data"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"strconv"

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
	// TODO make recursive file inclusion - have it in bookmarks somewhere
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct{}

var raceMap = map[string]string{
	"slavs":        "Славянин",
	"other":        "Относится к одной из редких этнических групп России",
	"middle_asian": "Житель средней азии",
	"exotic":       "Один из: черный, латиноамериканец, американец, европеец",
	"any":          "Этническая принадлежность не важна",
}

func newPage() Page {
	return Page{}
}

type FormResults struct {
	List      []template.HTML
	Age       int
	Race      string
	Height    int
	Money     int
	IsMarried bool
	Chance    float32
	Score     int
	Img       string
	Page      string
}

type IPages struct {
	INDEX     string
	RESULT    string
	ABOUT     string
	RESOURCES string
}

var Pages = IPages{
	INDEX:     "",
	RESULT:    "result",
	ABOUT:     "about",
	RESOURCES: "resources",
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/static", "static")

	e.GET("/"+Pages.INDEX, func(c echo.Context) error {
		return c.Render(200, "index", struct {
			Page   string
			Header string
			Age1   int
			Age2   int
		}{Page: "index", Header: "header", Age1: 20, Age2: 30})
	})

	e.GET("/"+Pages.RESULT, func(c echo.Context) error {
		minAge, minAgeErr := strconv.Atoi(c.QueryParam("age-min"))
		maxAge, maxAgeErr := strconv.Atoi(c.QueryParam("age-max"))
		race := c.QueryParam("race")
		height, heightOk := strconv.Atoi(c.QueryParam("height"))
		money, moneyOk := strconv.Atoi(c.QueryParam("money"))
		isMarried, _ := strconv.ParseBool(c.QueryParam("married"))

		if moneyOk != nil {
			return errors.New("Money is missing")
		}
		if heightOk != nil {
			return errors.New("heightOk is missing")
		}
		if minAgeErr != nil {
			return errors.New("minAgeErr is missing")
		}
		if maxAgeErr != nil {
			return errors.New("maxAgeErr is missing")
		}

		chance := data.Stats.CalcChance(minAge, maxAge, race, height, money, isMarried) * 100

		var score int
		var img string

		if chance <= 1 {
			score = 0
			img = "1.png"
		} else if chance <= 5 { // hold your horses lady
			score = 1
			img = "2.jpg"
		} else if chance <= 15 { // are you a feminist?
			score = 2
			imgs := []string{"3.jpg", "3_1.jpg"}
			img = imgs[rand.Intn(len(imgs))]
		} else if chance <= 40 { // potential catlady
			score = 3
			img = "4.jpg"

		} else if chance <= 65 { // down to earth
			score = 4
			img = "5.jpg"
		} else if chance <= 95 { // tradwife material
			score = 5
			img = "6.jpg"

		} else if chance > 95 { // probably not a woman
			score = 6
		}

		list := make([]template.HTML, 0)
		if minAgeErr == nil || maxAgeErr == nil {
			list = append(list, template.HTML("Между <span class='answer'>"+fmt.Sprint(minAge)+"</span> и <span class='answer'>"+fmt.Sprint(maxAge)))
		} else {
			list = append(list, "Любого возраста")
		}

		if race == "any" {
			list = append(list, "Любой этнической принадлежности")
		} else {
			list = append(list, template.HTML("<span class='answer'>"+fmt.Sprint(raceMap[race])+"</span>"))
		}

		if heightOk == nil {
			list = append(list, template.HTML("Как минимум <span class='answer'>"+fmt.Sprint(height)+"см</span> "))
		} else {
			list = append(list, "Любого роста")
		}

		if moneyOk == nil {
			list = append(list, template.HTML("Должен зарабатывать как минимум <span class='answer'>"+fmt.Sprint(money/1000)+"т.р.</span> "))
		} else {
			list = append(list, template.HTML("Любой заработок"))
		}

		if isMarried {
			list = append(list, "Должен быть <span class='answer'>не женат</span>")
		} else {
			list = append(list, "Семейное положение <span class='answer'>не важно</span>")
		}

		formResults := FormResults{
			List:      list,
			IsMarried: isMarried,
			Chance:    chance,
			Score:     score,
			Img:       img,
			Page:      Pages.RESULT,
		}
		c.Render(200, "index", formResults)
		return nil
	})

	e.GET("/"+Pages.ABOUT, func(c echo.Context) error {
		c.Render(200, "index", struct{ Page string }{
			Page: Pages.ABOUT,
		})
		return nil
	})

	e.GET("/"+Pages.RESOURCES, func(c echo.Context) error {
		c.Render(200, "index", struct{ Page string }{
			Page: Pages.RESOURCES,
		})
		return nil
	})

	e.Logger.Fatal(e.Start(":42069"))
}
