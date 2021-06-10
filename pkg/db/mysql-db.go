package db

import (
	"bookstore/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	//Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bookstore?parseTime=true&charset=utf8&loc=Local")
	config.Init()
	Db, err = sql.Open("mysql", viper.GetString("mysql.source_name"))
	Db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	err = Db.Ping()
	if nil != err {
		panic(err)
	} else {
		log.Println("MySQL Startup Normal!")
	}
}
