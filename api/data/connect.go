package data

import (
	"database/sql"
	"fmt"

	// without the underscore _, you will get imported but not
	// used error message
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Db struct {
	user     string
	password string
	dbname   string
	conn     *sql.DB
}

func (db *Db) connect() {
	// connect to our database server with data source name
	// data source name configuration has the following parameters :
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	// example config :
	// user:password@tcp(127.0.0.1:3306)/database

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", db.user, db.password, db.dbname))

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

func (db *Db) Query(query string) *sql.Rows {
	if db.conn == nil {
		db.connect()
	}

	statement, err := db.conn.Prepare(query)
	fmt.Println(statement)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query() // execute our statement

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.closeConnection()
	return rows
}

type Game struct {
	id     int
	name   string
	winner string
}

func main() {

	db := &Db{
		user:     "3333",
		password: "333333.",
		dbname:   "3333",
	}

	rows := db.Query("select name, winner from game limit 10;")
	for rows.Next() {
		var name, winner string
		rows.Scan(&name, &winner)
		fmt.Printf("Title of game is %s and the winner was %s\n", name, winner)
	}

}
