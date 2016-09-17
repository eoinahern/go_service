package model

import (
	"database/sql"
	"fmt"
	"log"
)

//Username: 	bd145d3b601f2e
//Password: 	532d35c9
//heroku_1587748f259385b

type baseDao struct {
	db *Database
}

func NewBaseDao(dbconnin *Database) *baseDao {
	basedao := new(baseDao)
	basedao.db = dbconnin
	return basedao
}

func (base *baseDao) CountRows(table string) int {

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	rows, err := base.db.mydbconn.Query(query)
	if err != nil {
		println("couldnt count rows!!!")
		println(query)
		log.Fatal(err)
	}

	defer rows.Close()
	return checkcount(rows)
}

func checkcount(rows *sql.Rows) int {

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		checkerr(err)
	}
	return count
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
