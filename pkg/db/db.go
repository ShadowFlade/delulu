package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/gofor-little/env"
)

type IStatistics struct {
	ID           int
	AGE_MIN      int
	AGE_MAX      int
	AGE          string
	SALARY       string
	PRICE        float32
	RACE         string
	HEIGHT       int
	IS_MARRIED   bool
	IP           interface{}
	DATE_CREATED string
}

type Db struct {
	db              *sql.DB
	statisticsTable string
	ipsTable        string
}

func (d *Db) Connect() {
	if err := env.Load("./.env.local"); err != nil {
		panic(err)
	}
	cfg := mysql.Config{
		User:                 env.Get("DB_LOGIN", "i"),
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               env.Get("DB_NAME", "your"),
		AllowNativePasswords: true,
	}

	d.ipsTable = env.Get("DB_IPS_TABLE", "mom")
	d.statisticsTable = env.Get("DB_STATISTICS_TABLE", ";)")
	var err error
	d.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := d.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// func (d *Db) WriteStatistics(fields map[string]interface{}) error {
// 	var stats map[string]interface{};
// 	for _, val := range fields {
// 		d.db.Query()
// 	}
// }

func (d *Db) GetStatistics() ([]IStatistics, error) {
	var statistics []IStatistics
	rows, err := d.db.Query("SELECT * FROM statistics")
	if err != nil {
		return nil, fmt.Errorf("i guess im dumb")
	}
	defer rows.Close()
	for rows.Next() {
		var statRow IStatistics
		if err := rows.Scan(&statRow.ID, &statRow.AGE_MIN, &statRow.AGE_MAX, &statRow.AGE, &statRow.SALARY, &statRow.PRICE, &statRow.RACE, &statRow.HEIGHT, &statRow.IS_MARRIED, &statRow.IP, &statRow.DATE_CREATED); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q", err)
		}
		//when can this trigger?
		statistics = append(statistics, statRow)
		if err := rows.Err(); err != nil {

			return nil, fmt.Errorf("wut")
		}
	}

	return statistics, nil
}
