package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Db struct {
	db              *sql.DB
	statisticsTable string
	ipsTable        string
}

func (d *Db) connect() {
	cfg := mysql.Config{
		User:   os.Getenv("DB_LOGIN"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "DB_NAME",
	}

	d.ipsTable = os.Getenv("unique_ips_temp")
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

func (d *Db) WriteStatistics(fields map[string]interface{}) error {
	for _, val := range fields {

	}
}

func (d *Db) GetStatistics() (error, map[string]interface{}) {
	var statistics []map[string]interface{}
	rows, err := d.db.Query("SELECT * FROM ? WHERE", d.statisticsTable)
	if err != nil {
		return fmt.Errorf("i guess im dumb"), nil
	}
	defer rows.Close()
	for rows.Next() {
		var statRow map[string]interface{}
		//TODO should i populate keys first lol?
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return fmt.Errorf("i guess im dumb"), nil
		}
		albums = append(albums, alb)
	}
}
