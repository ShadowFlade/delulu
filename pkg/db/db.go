package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func (d *Db) Connect() *sqlx.DB {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	if err := env.Load("./.env"); err != nil {
		panic(err)
	}
	d.ipsTable = env.Get("DB_IPS_TABLE", "mom")
	d.statisticsTable = env.Get("DB_STATISTICS_TABLE", ";)")
	d.login = env.Get("DB_LOGIN", "i")
	d.password = env.Get("DB_PASS", "fucked")
	d.dbName = env.Get("DB_NAME", "urmom")

	connectStr := fmt.Sprintf("%s:@(127.0.0.1:3306)/%s", d.login, d.dbName)
	db, err := sqlx.Connect("mysql", connectStr)

	d.checkAndCreateNewStructure(*db)
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

func (d *Db) checkAndCreateNewStructure(db sqlx.DB) (bool, error) {

	rows, err := db.Queryx("SELECT table_name FROM information_schema.tables WHERE table_schema = 'delulu'")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to store table names
	tables := []string{}

	// Iterate over the result set and retrieve table names
	tableCount := 0
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName, " table name")
		tables = append(tables, tableName)
		tableCount++
	}

	fmt.Println(tableCount, " table count")
	if tableCount == 0 {
		execMultipleSqlStatementsFile("./create_tables.sql", db)
		// _, err = sqlx.LoadFile(db, "./create_tables.sql")

		if err != nil {
			fmt.Println("Error loading and executing SQL file:", err)
		}
	}
	// Print the table names
	for _, table := range tables {
		fmt.Println(table)
	}
	return true, nil

}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		e.Error()
		panic(e)
	}
}
func execMultipleSqlStatementsFile(filename string, db sqlx.DB) {
	fmt.Println("starting splitting")
	//determine if there are multiple sql statements in a file (by detecting empty lines) and then split this file into multiple files so each one contains one mysql statement and then we exec them one by one
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Print(string(data), " data")
	sqlStatements := strings.Split(string(data), "\n\n")
	files := make([]os.File, 0)

	for index, sqlStatement := range sqlStatements {
		file, err := os.Create("sql.tmp" + fmt.Sprint(index) + ".sql")
		check(err)
		file.WriteString(sqlStatement)
		files = append(files, *file)
	}

	for _, file := range files {
		check(err)
		_, err = sqlx.LoadFile(db, file.Name())
		check(err)
	}

	//	f, err := os.Open(filename)
	//	check(err)
	//
	//	b1 := make([]byte, 5)
	//	h1, err := f.Read(b1);
	//	check(err)
	//
	// fmt.Printf("%d byes: %s\n", n1, string(b1[:n1]))
	//
	//	o2, err := f.Seek
}
