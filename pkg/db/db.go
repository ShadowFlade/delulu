package db

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofor-little/env"
	"github.com/jmoiron/sqlx"
)

type IStatistics struct {
	ID           int         `db:"id"`
	AGE_MIN      int         `db:"age_min"`
	AGE_MAX      int         `db:"age_max"`
	SALARY       string      `db:"salary"`
	RACE         string      `db:"race"`
	HEIGHT       int         `db:"height"`
	IS_MARRIED   bool        `db:"is_married"`
	IP           interface{} `db:"ip"`
	DATE_CREATED string      `db:"date_created"`
}

type Db struct {
	db              *sqlx.DB
	tx              *sqlx.Tx
	dbName          string
	login           string
	password        string
	statisticsTable string
	ipsTable        string
	cols            []string
}

var schema = `
CREATE TABLE statistics (
    id int auto_increment primary key,
    age_min varchar(255),
    age_max varchar(255),
    salary varchar(255),
    race varchar(255),
    height int,
    is_married tinyint(1),
    ip varchar(255),
    date_created date
);

CREATE TABLE feedback (
    id int auto_increment primary key,
    name varchar(255),
    description varchar(255),
    email varchar(255)
);

CREATE TABLE unique_ips_temp (
    id int auto_increment primary key,
    ip varchar(15),
    date_created date
)`
func (d *Db) Connect() *sqlx.DB {
	if err := env.Load("./.env.local"); err != nil {
		panic(err)
	}
	d.ipsTable = env.Get("DB_IPS_TABLE", "mom")
	d.statisticsTable = env.Get("DB_STATISTICS_TABLE", ";)")
	d.login = env.Get("DB_LOGIN", "i")
	d.password = env.Get("DB_PASS", "fucked")
	d.dbName = env.Get("DB_NAME", "urmom")

	connectStr := fmt.Sprintf("%s:@(127.0.0.1:3306)/%s", d.login, d.dbName)
	db, err := sqlx.Connect("mysql", connectStr)

	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	d.db = db

	return db

}

func (d *Db) GetStatistics() ([]IStatistics, error) {
	stats := []IStatistics{}

	err := d.db.Select(&stats, "select * from statistics")
	if err != nil {
		return nil, err
	}
	return stats, nil
}

type RowsResult interface {
	[][]string
}

func (d *Db) WriteStatistics(stats interface{}) (int64, error) {
	tx := d.db.MustBegin()

	res, err := tx.NamedExec(`INSERT INTO statistics (age_min, age_max, salary, race, height, is_married, ip, date_created) VALUES (:age_min, :age_max,
        :salary, :race, :height, :is_married, :ip, now())`, stats)

	if err != nil {
		return 0.00, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0.00, err
	}

	errN := tx.Commit()

	if errN != nil {
		return 0.00, errN
	}

	return id, nil
}

func (d *Db) WriteFeedback(feedback interface{}) (int64, error) {
	id, err := d.Write("feedback", []string{"name", "description", "email"}, feedback)
	fmt.Println(id, " feedback id")
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Db) Write(table string, cols []string, feedback interface{}) (int64, error) {

	tx := d.db.MustBegin()
	vals := make([]string, 0, len(cols))

	for _, val := range cols {
		vals = append(vals, ":"+val)
	}

	query := fmt.Sprintf("insert into %s (%s) values (%s)", table, strings.Join(cols, ", "), strings.Join(vals, ", "))

	fmt.Println(query, " query")
	res, err := tx.NamedExec(query, feedback)

	if err != nil {
		return 0.00, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0.00, err
	}

	errN := tx.Commit()

	if errN != nil {
		return 0.00, errN
	}

	return id, nil
}
