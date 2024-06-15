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

func (d *Db) GetStatistics() ([]string, error) {
	var stats []string
	rows, err := d.db.Query("SELECT * FROM statistics")

	if err != nil {
		return nil, fmt.Errorf("i guess im dumb")
	}
	defer rows.Close()

	stats, err = parseValues(rows)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func parseValues(rows *sql.Rows) ([]string, error) {
	var results []string
	var columns []interface{}
	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for _, val := range cols {
		columns = append(columns, new(string))
        fmt.Println(val)
	}

	for rows.Next() {
		err := rows.Scan(columns...)
        fmt.Println(columns," columns")

		if err != nil {
            return nil, fmt.Errorf("scan result error:%s", err)
		}

        for _, val := range columns {
            strVal := *val.(*string)
            results = append(results, strVal);
        }
	}
	return results, nil
}
