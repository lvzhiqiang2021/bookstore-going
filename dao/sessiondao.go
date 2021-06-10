package dao

import (
	"bookstore/model"
	db "bookstore/pkg/db"
	"net/http"
)

func AddSession(sess *model.Session) error {
	sql := "insert into sessions values(?,?,?)"
	_, err := db.Db.Exec(sql, &sess.SessionID, &sess.UserName, &sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sql := "delete from sessions where session_id = ?"
	_, err := db.Db.Exec(sql, sessID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessID string) (*model.Session, error) {
	sql := "select session_id, username, user_id from sessions where session_id = ?"
	Stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	row := Stmt.QueryRow(sessID)
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
