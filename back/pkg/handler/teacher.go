package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeacher(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	teacherUser, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}
	teachers, err := rdb.GetAllTeachers(c)
	if err != nil {
		log.Fatal(err)
	}
	students, err := rdb.GetAllStudents(c)
	if err != nil {
		log.Fatal(err)
	}

	form1, err1 := rdb.GetForm(c, 1)
	if err1 != nil {
		form1 = &rdb.Form{}
	}
	form2, err2 := rdb.GetForm(c, 2)
	if err2 != nil {
		form2 = &rdb.Form{}
	}
	form3, err3 := rdb.GetForm(c, 3)
	if err3 != nil {
		form3 = &rdb.Form{}
	}

	item := struct {
		Title    string
		Message  string
		Account  string
		Teachers []rdb.Teacher
		Students []rdb.Student
		ID       string
		Err1     error
		Err2     error
		Err3     error
		Form1    *rdb.Form
		Form2    *rdb.Form
		Form3    *rdb.Form
	}{
		Title:    "",
		Message:  "教師用ページへようこそ" + teacherUser.Name + "さん。",
		Account:  teacherUser.Name,
		Teachers: teachers,
		Students: students,
		ID:       "teacher",
		Err1:     err1,
		Err2:     err2,
		Err3:     err3,
		Form1:    form1,
		Form2:    form2,
		Form3:    form3,
	}

	er := pkg.Page("teacher").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func GetDeleteTeacher(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	ID := c.Query("teacher_id")
	if ID == "" {
		http.Error(c.Writer, "Teacher ID is required", http.StatusBadRequest)
		return
	}

	err := rdb.DeleteTeacher(c, ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}

func GetEditTeacherNameAndPassword(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	teacherUser, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}

	ID := c.Query("teacher_id")
	if ID == "" {
		http.Error(c.Writer, "Teacher ID is required", http.StatusBadRequest)
		return
	}

	teacher, err := rdb.GetTeacher(c, ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	item := struct {
		Title     string
		Message   string
		Account   string
		Teacher   *rdb.Teacher
		ID        string
		TeacherID string
	}{
		Title:     "教師名とパスワードの変更",
		Message:   teacher.Name + "さんの名前とパスワードを変更してください。",
		Account:   teacherUser.Name,
		Teacher:   teacher,
		ID:        "teacher",
		TeacherID: teacher.ID,
	}

	// Render the edit page with the teacher data.
	er := pkg.Page("teacher_name_password_edit").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostEditTeacherNameAndPassword(c *gin.Context) {
	ID := c.Query("teacher_id")
	newName := c.PostForm("new_name")
	newPassword := c.PostForm("new_password")

	if ID == "" || newName == "" || newPassword == "" {
		http.Error(c.Writer, "Teacher ID, new name and new password are required", http.StatusBadRequest)
		return
	}

	err := rdb.UpdateTeacherNameAndPassword(c, ID, newName, newPassword)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")

}
