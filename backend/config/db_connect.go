package config

import (
	"database/sql"
	"fmt"
)

var DbConn *sql.DB

// whole point of this will be to return variable db, which
//will help the other files to interact the db

// dbconn will help us to open connection with database
// videos https://www.youtube.com/watch?v=1E_YycpCsXw
// oli veel Open all teise positsioonina l]pus mingi pikem asi charset utf jne

func DBConn() {
	db, err := sql.Open("sqlite3", "rltforum.db")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB Connected!!")
	DbConn = db
}

// dont exactly get the polint of this function
//
func GetDB() *sql.DB {
	return DbConn
}

// FoorumDao will be used by functions which write into and out of database
// registering data, posts, comments
// as of 29.11 this part is not used anywhere
type FoorumDao struct {
	db *sql.DB
}

var foorum_dao *FoorumDao
