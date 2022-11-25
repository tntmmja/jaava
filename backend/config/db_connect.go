package config

import (
	"database/sql"
	"fmt"
)

func dbConn() (db *sql.DB) {

	db, err := sql.Open("sqlite3", "rltforum.db")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB Connected!!")
	return db
}
