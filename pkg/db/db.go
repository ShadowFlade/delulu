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
}

func (d *Db) GetStatistics() ([]map[string]string, error) {
	rows, err := d.db.Query("SELECT * FROM statistics")

	if err != nil {
		return nil, fmt.Errorf("i guess im dumb")
	}
	defer rows.Close()

	stats, err := parseValues(rows)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

type RowsResult interface {
	[][]string
}

func parseValues[T RowsResult](rows *sql.Rows) ([]map[string]string, error) {
	var results T
	var columns []interface{}
	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for range cols {
		columns = append(columns, new(string))
	}

	for rows.Next() {
		var result []string
		err := rows.Scan(columns...)

		if err != nil {
			return nil, fmt.Errorf("scan result error:%s", err)
		}

		for _, val := range columns {
			strVal := *val.(*string)
			result = append(result, strVal)
		}
		results = append(results, result)
	}

	mapResults := make([]map[string]string, len(results))

	for _, rowResult := range results {

		temp := map[string]string{}
		for index, val := range rowResult {
			temp[cols[index]] = val
		}

		mapResults = append(mapResults, temp)
        temp = nil
	}

	return mapResults, nil
}
