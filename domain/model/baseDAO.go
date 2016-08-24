package model

import (
	"database/sql"
	"log"
)

type baseDao struct {
	db *Database
}

func NewBaseDao(dbconnin *Database) *baseDao {
	basedao := new(baseDao)
	basedao.db = dbconnin
	return basedao
}

func (base *baseDao) CountRows(table string) int {

	rows, err := base.db.mydbconn.Query("SELECT COUNT(*) FROM %s", table)
	if err != nil {
		println("couldnt count rows!!!")
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
