package data

import (
	"database/sql"
	"fmt"
)

func main() {
	db := &Db{
		user:     "12345",
		password: "12345.",
		dbname:   "12345",
	}

	rows := db.Query("select name, winner from game limit 10;")
	for rows.Next() {
		var name, winner string
		rows.Scan(&name, &winner)
		fmt.Printf("Title of game is %s and the winner was %s\n", name, winner)
	}
}
