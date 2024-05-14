package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"unicode"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
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

	er := pkg.Page("login").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostLogin(c *gin.Context) {
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
	usr := c.PostForm("account")
	pass := c.PostForm("pass")
	// het first character of account.
	firstRune := rune(usr[0])
	// check student(in if) or teacher(in else) account and password .
	if unicode.IsLetter(firstRune) {
		student, err := rdb.GetStudent(c, usr)
		if err != nil {
			panic(err)
		}
		if student.Password != pass {
			item.Message = "Account or Password is wrong."
			er := pkg.Page("login").Execute(c.Writer, item)
			if er != nil {
				log.Fatal(er)
			}
			return
		}
		item.Account = student.Name
		ses := c.MustGet("session").(sessions.Session)
		ses.Set("login", true)
		ses.Set("account", student.ID)
		ses.Set("name", student.Name)
		err = ses.Save()
		if err != nil {
			// handle error
		}
		c.Redirect(http.StatusFound, "/student")
	} else {
		teacher, err := rdb.GetTeacher(c, usr)
		if err != nil {
			panic(err)
		}
		if teacher.Password != pass {
			item.Message = "Account or Password is wrong."
			er := pkg.Page("login").Execute(c.Writer, item)
			if er != nil {
				log.Fatal(er)
			}
			return
		}

		item.Account = teacher.Name
		ses := c.MustGet("session").(sessions.Session)
		ses.Set("login", true)
		ses.Set("account", teacher.ID)
		ses.Set("name", teacher.Name)
		err = ses.Save()
		if err != nil {
			// handle error
		}
		c.Redirect(http.StatusFound, "/teacher")
	}
}

func GetLogout(c *gin.Context) {
	ses := c.MustGet("session").(sessions.Session)
	ses.Set("login", false)
	ses.Set("account", "")
	ses.Set("name", "")
	err := ses.Save()
	if err != nil {
		// handle error
	}
	c.Redirect(http.StatusFound, "/login")
}
