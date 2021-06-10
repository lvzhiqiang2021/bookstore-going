package controller

import (
	"bookstore/dao"
	"bookstore/model"
	utils "bookstore/pkg/utils"
	"fmt"
	"net/http"
	"text/template"
)

func IndexTest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Hello Index!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag {
		//跳转到首页
		GetPageBooksByPrice(w, r)
		//IndexTest(w, r)
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			uuid := utils.CreateUUID()
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			dao.AddSession(sess)
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			//渲染模板
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确")
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	//跳转到首页
	GetPageBooksByPrice(w, r)
	//IndexTest(w, r)
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)

	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在")
	} else {
		err := dao.SaveUser(username, password, email)
		if err != nil {
			t := template.Must(template.ParseFiles("views/pages/user/regist_error.html"))
			t.Execute(w, "注册失败,请联系网站管理员!")
		}else{
			t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
			t.Execute(w, "用户注册成功!")
		}

	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已经存在!"))
	} else {
		w.Write([]byte("<font style = 'color:green'>用户名不存在！</font>"))
	}
}
