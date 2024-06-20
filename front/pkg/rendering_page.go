package pkg

import (
	dbpkg "cram-school-reserve-server/back/infra/rdb"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"net/http"
	"text/template"
	"unicode"
)

func CheckLogin(c *gin.Context) interface{} {
	ses := c.MustGet("session").(sessions.Session)
	loginValue := ses.Get("login")
	if loginValue == nil || !loginValue.(bool) {
		c.Redirect(http.StatusFound, "/login")
		return nil
	}
	ac := ""
	accountValue := ses.Get("account")
	if accountValue != nil {
		ac = accountValue.(string)
	}

	db := c.MustGet("db").(*gorm.DB)

	firstRune := rune(ac[0])
	if unicode.IsLetter(firstRune) {
		var student dbpkg.Student
		db.Where("id = ?", ac).First(&student)
		return &student
	} else {
		var teacher dbpkg.Teacher
		db.Where("id = ?", ac).First(&teacher)
		return &teacher
	}
}

func Page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("front/templates/"+fname+".html",
		"front/templates/exceptionReserveFormTable.html", "front/templates/head.html", "front/templates/foot.html", "front/templates/student_table.html", "front/templates/reserve_form_table.html")
	return tmps
}
