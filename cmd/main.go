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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", struct {
			Page   string
			Header string
			Age1   int
			Age2   int
		}{Page: "index", Header: "header", Age1: 20, Age2: 30})
	})

	e.GET("/result", func(c echo.Context) error {
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

		chance := data.Stats.CalcChance(minAge, maxAge, race, height, money, isMarried)

		var score int
		var img string

		if chance <= 1 {
			score = 0
			img = "0.jpg"
		} else if chance <= 5 { // hold your horses lady
			score = 1
			img = "1.jpg"
		} else if chance <= 15 { // are you a feminist?
			score = 2
			img = "2.jpg"
		} else if chance <= 40 { // potential catlady
			score = 3
			imgs := []string{"3.jpg", "3_1.jpg"}
			img = imgs[rand.Intn(len(imgs))]
		} else if chance <= 65 { // down to earth
			score = 4
			img = "4.jpg"
		} else if chance <= 95 { // tradwife material
			score = 5
			img = "5.jpg"
		} else if chance > 95 { // probably not a woman
			score = 6
			img = "6.jpg"
		}

		list := make([]template.HTML, 0)
		fmt.Println(minAge, " min age")
		if minAgeErr == nil || maxAgeErr == nil {
			list = append(list, template.HTML("Между <span>"+fmt.Sprint(minAge)+"</span> и <span>"+fmt.Sprint(maxAge)))
		} else { // akshually u can write it in html/template but me stupid amd now im too lazy to do it
			list = append(list, "Любого возраста")
		}

		if race == "any" {
			list = append(list, "Любой этнической принадлежности")
		} else {
			list = append(list, template.HTML(raceMap[race]))
		}

		if heightOk == nil {
			list = append(list, template.HTML("Как минимум <span>"+fmt.Sprint(height)+"</span> см"))
		} else {
			list = append(list, "Любого роста")
		}

		if moneyOk == nil {
			list = append(list, template.HTML("Должен зарабатывать как минимум <span>"+fmt.Sprint(money/1000)+"</span> т.р."))
		} else {
			list = append(list, template.HTML("Любой заработок"))
		}

		if isMarried {
			list = append(list, "Должен быть не женат")
		} else {
			list = append(list, "Семейное положение не важно")
		}

		race = raceMap[race]
		fmt.Println(race, " race")

		formResults := FormResults{
			List:      list,
			IsMarried: isMarried,
			Chance:    chance,
			Score:     score,
			Img:       img,
			Page:      "result",
		}
		c.Render(200, "index", formResults)
		return nil
	})

	e.Logger.Fatal(e.Start(":42069"))
}
