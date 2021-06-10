package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
	"log"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sql := "select id, username, password, email from users where username= ? and password = ?"
	row := db.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sql := "select id, username, password, email from users where username = ?"
	row := db.Db.QueryRow(sql, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func SaveUser(username string, password string, email string) error {
	sql := "insert into users(username, password, email)values(?,?,?)"
	_, err := db.Db.Exec(sql, username, password, email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
