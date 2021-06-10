package pkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bookstore?parseTime=true&charset=utf8&loc=Local")

	if err != nil {
		panic(err.Error())
	}

}
