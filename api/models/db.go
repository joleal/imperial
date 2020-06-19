package models

import (
	"database/sql"
	"fmt"
	"github.com/dovadi/dbconfig"
	"os"
	//blank import just to overwrite the base interface
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	user     string
	password string
	dbname   string
	conn     *sql.DB
}

func CreateDb() *Db {
	return &Db{
		user:   "ubuntu",
		dbname: "imperial",
	}
}

func (db *Db) connect() {
	connectionString := dbconfig.MysqlConnectionString("config/db_settings/db.json")

	conn, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.conn = conn
}

func (db *Db) closeConnection() {
	db.conn.Close()
	db.conn = nil
}

//Query executes a query
func (db *Db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.conn == nil {
		db.connect()
	}

	statement, err := db.conn.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}

	rows, err := statement.Query(args...) // execute our statement

	if err != nil {
		fmt.Println(err)
	}

	db.closeConnection()
	return rows, err
}
