package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"unicode"
)

func Signup(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Title   string
		Message string
		Account string
		ID      string
	}{
		Title:   " ",
		Message: "アカウントID、名前、パスワードを入力してください。",
		Account: "",
		ID:      "login",
	}

	if rq.Method == "GET" {
		er := pkg.Page("signup").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
		return
	}
	if rq.Method == "POST" {
		ID := rq.PostFormValue("account")
		name := rq.PostFormValue("name")
		pass := rq.PostFormValue("pass")
		firstRune := rune(ID[0])
		if unicode.IsLetter(firstRune) {
			// Check if the student ID already exists
			_, err := rdb.GetStudent(ID)
			if err == nil {
				// If the student ID already exists, show an error message and return
				item.Message = "The account ID is already in use. Please use a different ID."
				er := pkg.Page("signup").Execute(w, item)
				if er != nil {
					log.Fatal(er)
				}
				return
			}
			student := rdb.Student{
				ID:             ID,
				Name:           name,
				Password:       pass,
				MaxFirstClass:  0,
				MinFirstClass:  0,
				MaxSecondClass: 0,
				MinSecondClass: 0,
				MaxThirdClass:  0,
				MinThirdClass:  0,
			}
			if err := rdb.CreateStudent(student); err != nil {
				panic(err)
			}
			item.Account = student.Name
			ses, _ := pkg.Cs.Get(rq, pkg.SesName)
			ses.Values["login"] = true
			ses.Values["account"] = student.ID
			ses.Values["name"] = student.Name
			ses.Save(rq, w)
			http.Redirect(w, rq, "/student", 302)
		} else {
			// Check if the teacher ID already exists
			_, err := rdb.GetTeacher(ID)
			if err == nil {
				// If the teacher ID already exists, show an error message and return
				item.Message = "The account ID is already in use. Please use a different ID."
				er := pkg.Page("signup").Execute(w, item)
				if er != nil {
					log.Fatal(er)
				}
				return
			}
			teacher := rdb.Teacher{
				ID:       ID,
				Name:     name,
				Password: pass,
			}
			if err := rdb.CreateTeacher(teacher); err != nil {
				panic(err)
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
