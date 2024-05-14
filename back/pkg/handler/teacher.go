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
	log.Println("GetTeacher 111")
	if user == nil {
		// Handle error here
		return
	}
	log.Println("GetTeacher 1")
	teacherUser, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}
	log.Println("GetTeacher 11")
	teachers, err := rdb.GetAllTeachers(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetTeacher 2")
	students, err := rdb.GetAllStudents(c)
	if err != nil {
		log.Fatal(err)
	}
	// Get the form data from the database.
	_, err1 := rdb.GetForm(c, 1)
	_, err2 := rdb.GetForm(c, 2)
	_, err3 := rdb.GetForm(c, 3)
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
	}
	log.Println("GetTeacher 3")
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
