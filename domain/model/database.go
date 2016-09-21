package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(name string, pass string, dbname string) *Database {
	db := new(Database)
	var err error

	//example TCP on a remote host, e.g. Amazon RDS:
	//id:password@tcp(your-amazonaws-uri.com:3306)/dbname

	dbstring := fmt.Sprintf("%s:%s@/%s", name, pass, dbname)
	db.mydbconn, err = sql.Open("mysql", dbstring)

	if err != nil {
		println("couldnt connect")
		log.Fatal("couldnt connect", "error")
		return nil
	}
	return db
}

type Database struct {
	mydbconn *sql.DB
}
