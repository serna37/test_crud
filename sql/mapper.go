package sql

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

/** db connection */
var db *gorm.DB

/*
Create DB Connection.
*/
func Conn() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	var err error
	db, err = gorm.Open(DRIVER, connectionString)
	//defer db.Close()
	if err != nil {
		log.Fatal("connection error")
	}
}
