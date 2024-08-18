package pkg

import (
	"bytes"
	"delulu/pkg/data"
	"delulu/pkg/db"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
)

var raceMap = map[string]string{
	"slavs":        "Славянин",
	"other":        "Относится к одной из редких этнических групп России",
	"middle_asian": "Житель средней азии",
	"exotic":       "Один из: черный, латиноамериканец, американец, европеец",
	"any":          "Этническая принадлежность не важна",
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
	Text      string
}
type Handlers struct {
}

func (this *Handlers) Result(c echo.Context) error {
	minAge, minAgeErr := strconv.Atoi(c.QueryParam("age-min"))
	maxAge, maxAgeErr := strconv.Atoi(c.QueryParam("age-max"))
	race := c.QueryParam("race")
	height, heightOk := strconv.Atoi(c.QueryParam("height"))
	money, moneyOk := strconv.Atoi(c.QueryParam("money"))
	excludeMarriedStr := c.QueryParam("married")
	excludeMarried := excludeMarriedStr == "on"

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

	chance := data.Stats.CalcChance(minAge, maxAge, race, height, money, excludeMarried) * 100

	var score int
	var img string
	var text string

	if chance <= 5 {
		score = 0
		imgs := []string{"587222.jpg", "975757.webp"}
		img = imgs[rand.Intn(len(imgs))]
		text = "Вам не место на этой планете"
	} else if chance <= 15 {
		score = 1
		img = "2.jpg"
		text = "Поздравляю, у вас феминизм!"
	} else if chance <= 30 {
		score = 2
		imgs := []string{"3.jpg", "3_1.jpg"}
		img = imgs[rand.Intn(len(imgs))]
		text = "Думаю, даже кошки вас уже не примут"
	} else if chance <= 65 {
		score = 3
		img = "4.jpg"
		text = "Потенциально хорошая жена"
	} else if chance <= 90 {
		score = 4
		img = "5.jpg"
		text = "Отличная жена и домохозяйка. Оставьте свой номер, с вами свяжутся"
	} else if chance > 90 {
		score = 5
		img = "6.jpg"
		text = "А вы точно женщина?"
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
		shouldEarn := money / 1000
		var shouldEarnStr string
		if shouldEarn != 0 {
			shouldEarnStr = fmt.Sprintf("Должен зарабатывать как минимум <span class='answer'>" + fmt.Sprint(shouldEarn) + "т.р.</span> ")
		} else {
			shouldEarnStr = fmt.Sprintf("Может зарабатывать <span class='answer'>любую сумму.</span> ")
		}
		list = append(list, template.HTML(shouldEarnStr))
	} else {
		list = append(list, template.HTML("Любой заработок"))
	}

	if excludeMarried {
		list = append(list, "Должен быть <span class='answer'>не женат</span>")
	} else {
		list = append(list, "Семейное положение <span class='answer'>не важно</span>")
	}

	formResults := FormResults{
		List:      list,
		IsMarried: excludeMarried,
		Chance:    chance,
		Score:     score,
		Img:       img,
		Text:      text,
		Page:      Pages.RESULT,
	}
	db := db.Db{}
	db.Connect()

	_, err := db.WriteStatistics(
		struct {
			AgeMin    int    `db:"age_min"`
			AgeMax    int    `db:"age_max"`
			Salary    int    `db:"salary"`
			Race      string `db:"race"`
			Height    int    `db:"height"`
			IsMarried bool   `db:"is_married"`
			Ip        string `db:"ip"`
		}{
			AgeMin:    minAge,
			AgeMax:    maxAge,
			Salary:    money,
			Race:      race,
			Height:    height,
			IsMarried: excludeMarried,
			Ip:        c.RealIP(),
		},
	)

	if err != nil {
		return err
	}

	c.Render(200, "index", formResults)
	return nil
}

func (this *Handlers) Feedback(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	description := c.Request().PostFormValue("description")
	email := c.Request().PostFormValue("email")
	db := db.Db{}
	db.Connect()
	_, err := db.WriteFeedback(struct {
		Name        string `db:"name"`
		Description string `db:"description"`
		Email       string `db:"email"`
	}{
		Name:        name,
		Description: description,
		Email:       email,
	})

	if err != nil {
		return err
	}

	c.Render(200, "form__feedback-result", nil)

	return nil

}

func (this *Handlers) CaptchaCheck(c echo.Context) error {
	recSecret := env.Get("RECAPTCHA_SECRET", "")

	formParams, err := c.FormParams()

	captchaBody := map[string]string{
		"response": formParams.Get("token"),
		"secret":   recSecret,
	}

	jsonValue, _ := json.Marshal(captchaBody)

	fmt.Println(formParams.Get("token"))
	fmt.Println(string(jsonValue), " json value")
	resp, err := http.Post("https://www.google.com/recaptcha/api/siteverify", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		c.Logger().Fatal("json error post")
	}

	// Read the response body
	defer resp.Body.Close() // Ensure the response body is closed
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c.Logger().Fatal("error reading response body")
	}
	fmt.Println(string(body), " body")
	return c.JSON(http.StatusOK, body)
}
