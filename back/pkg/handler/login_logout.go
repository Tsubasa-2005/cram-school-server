package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"unicode"
)

func Login(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Title   string
		Message string
		Account string
		ID      string
	}{
		Title:   " ",
		Message: "アカウントIDとパスワードを入力してください。",
		Account: "",
		ID:      "login",
	}

	if rq.Method == "GET" {
		er := pkg.Page("login").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
		return
	}
	if rq.Method == "POST" {
		// get form data.
		usr := rq.PostFormValue("account")
		pass := rq.PostFormValue("pass")
		// het first character of account.
		firstRune := rune(usr[0])
		// check student(in if) or teacher(in else) account and password .
		if unicode.IsLetter(firstRune) {
			student, err := rdb.GetStudent(usr)
			if err != nil {
				panic(err)
			}
			if student.Password != pass {
				item.Message = "Account or Password is wrong."
				er := pkg.Page("login").Execute(w, item)
				if er != nil {
					log.Fatal(er)
				}
				return
			}
			item.Account = student.Name
			ses, _ := pkg.Cs.Get(rq, pkg.SesName)
			ses.Values["login"] = true
			ses.Values["account"] = student.ID
			ses.Values["name"] = student.Name
			ses.Save(rq, w)
			http.Redirect(w, rq, "/student", 302)
		} else {
			teacher, err := rdb.GetTeacher(usr)
			if err != nil {
				panic(err)
			}
			if teacher.Password != pass {
				item.Message = "Account or Password is wrong."
				er := pkg.Page("login").Execute(w, item)
				if er != nil {
					log.Fatal(er)
				}
				return
			}

			item.Account = teacher.Name
			ses, _ := pkg.Cs.Get(rq, pkg.SesName)
			ses.Values["login"] = true
			ses.Values["account"] = teacher.ID
			ses.Values["name"] = teacher.Name
			ses.Save(rq, w)
			http.Redirect(w, rq, "/teacher", 302)
		}
	}
}

func Logout(w http.ResponseWriter, rq *http.Request) {
	ses, _ := pkg.Cs.Get(rq, pkg.SesName)
	ses.Values["login"] = nil
	ses.Values["account"] = nil
	ses.Values["name"] = nil
	ses.Save(rq, w)
	http.Redirect(w, rq, "/login", 302)
}
