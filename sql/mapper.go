package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

/** db connection */
var db *sql.DB

/*
Create DB Connection.
*/
func Conn() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	var err error
	db, err = sql.Open(DRIVER, connectionString)
	if err != nil {
		log.Fatal("connection error")
	}
}
