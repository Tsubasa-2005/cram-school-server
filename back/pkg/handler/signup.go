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

func GetSignup(c *gin.Context) {
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

	err := pkg.Page("signup").Execute(c.Writer, item)
	if err != nil {
		return
	}
}

func PostSignup(c *gin.Context) {
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
	ID := c.PostForm("account")
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	firstRune := rune(ID[0])
	if unicode.IsLetter(firstRune) {
		// Check if the student ID already exists
		_, err := rdb.GetStudent(c, ID)
		if err == nil {
			// If the student ID already exists, show an error message and return
			item.Message = "The account ID is already in use. Please use a different ID."
			er := pkg.Page("signup").Execute(c.Writer, item)
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
		if err := rdb.CreateStudent(c, student); err != nil {
			panic(err)
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
		// Check if the teacher ID already exists
		_, err := rdb.GetTeacher(c, ID)
		if err == nil {
			// If the teacher ID already exists, show an error message and return
			item.Message = "The account ID is already in use. Please use a different ID."
			er := pkg.Page("signup").Execute(c.Writer, item)
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
		if err := rdb.CreateTeacher(c, teacher); err != nil {
			panic(err)
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
