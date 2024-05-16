package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	studentUser, ok := user.(*rdb.Student)
	if !ok {
		// Handle error here
		return
	}
	// Initialize the start and end dates as empty strings
	var startDate1, endDate1, startDate2, endDate2, startDate3, endDate3 string
	var isWithinReservePeriod1, isWithinReservePeriod2, isWithinReservePeriod3 bool
	// Get the form data from the database for formID 1, 2, and 3.
	form1, err := rdb.GetForm(c, 1)
	if err == nil {
		startDate1 = form1.StartDate
		endDate1 = form1.EndDate
		startDate1, endDate1, isWithinReservePeriod1, err = rdb.FormatDate(form1.ReserveStartDate, form1.ReserveEndDate)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		startDate1 = ""
		endDate1 = ""
	}
	form2, err := rdb.GetForm(c, 2)
	if err == nil {
		startDate2 = form2.StartDate
		endDate2 = form2.EndDate
		startDate2, endDate2, isWithinReservePeriod2, err = rdb.FormatDate(form2.ReserveStartDate, form2.ReserveEndDate)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		startDate2 = ""
		endDate2 = ""
	}
	form3, err := rdb.GetForm(c, 3)
	if err == nil {
		startDate3 = form3.StartDate
		endDate3 = form3.EndDate
		startDate3, endDate3, isWithinReservePeriod3, err = rdb.FormatDate(form3.ReserveStartDate, form3.ReserveEndDate)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		startDate3 = ""
		endDate3 = ""
	}
	item := struct {
		Title                  string
		Message                string
		Account                string
		ID                     string
		StartDate1             string
		EndDate1               string
		StartDate2             string
		EndDate2               string
		StartDate3             string
		EndDate3               string
		IsWithinReservePeriod1 bool
		IsWithinReservePeriod2 bool
		IsWithinReservePeriod3 bool
	}{
		Title:                  " ",
		Message:                "生徒用ページへようこそ" + studentUser.Name + "さん。",
		Account:                studentUser.Name,
		ID:                     "student",
		StartDate1:             startDate1,
		EndDate1:               endDate1,
		StartDate2:             startDate2,
		EndDate2:               endDate2,
		StartDate3:             startDate3,
		EndDate3:               endDate3,
		IsWithinReservePeriod1: isWithinReservePeriod1,
		IsWithinReservePeriod2: isWithinReservePeriod2,
		IsWithinReservePeriod3: isWithinReservePeriod3,
	}
	er := pkg.Page("student").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostDeleteStudent(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	_, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}

	studentIDs := c.PostFormArray("any_student_id")
	if len(studentIDs) == 0 {
		studentIDs = c.PostFormArray("student_id")
	}

	for _, studentID := range studentIDs {
		studentID = strings.Trim(studentID, "[]")
		splitData := strings.Split(studentID, " ")
		for _, element := range splitData {
			studentID = element
			err := rdb.DeleteStudent(c, studentID)
			if err != nil {
				http.Error(c.Writer, "Failed update student class", http.StatusBadRequest)
				return
			}
		}
	}
	c.Redirect(http.StatusSeeOther, "/teacher")
}

func PostEditStudentClass(c *gin.Context) {
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

	studentIDs := c.PostFormArray("any_student_id")
	if len(studentIDs) == 0 {
		studentIDs = c.PostFormArray("student_id")
	}
	var studentNames []string
	for _, id := range studentIDs {
		student, err := rdb.GetStudent(c, id)
		if err != nil {
			// Handle error here
			continue
		}
		studentNames = append(studentNames, student.Name)
	}

	item := struct {
		Title    string
		Message  string
		Account  string
		Students []string
		IDs      []string
		ID       string
	}{
		Title:    "生徒取得コマ数変更",
		Message:  "生徒が選択できるコマ数を変更してください。",
		Account:  teacherUser.Name,
		Students: studentNames,
		IDs:      studentIDs,
		ID:       "teacher",
	}

	// Render the edit page with the student data.
	er := pkg.Page("student_class_edit").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostUpdateStudentClass(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {
		// Handle error here
		return
	}
	_, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}

	studentIDs := c.PostFormArray("student_id")
	log.Println(studentIDs)

	// Get the new MaxClass and MinClass values from the form data.
	maxFirstClass, err := strconv.Atoi(c.PostForm("max_first_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid max first class value", http.StatusBadRequest)
		return
	}
	minFirstClass, err := strconv.Atoi(c.PostForm("min_first_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid min first class value", http.StatusBadRequest)
		return
	}

	maxSecondClass, err := strconv.Atoi(c.PostForm("max_second_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid max second class value", http.StatusBadRequest)
		return
	}
	minSecondClass, err := strconv.Atoi(c.PostForm("min_second_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid min second class value", http.StatusBadRequest)
		return
	}

	maxThirdClass, err := strconv.Atoi(c.PostForm("max_third_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid max third class value", http.StatusBadRequest)
		return
	}
	minThirdClass, err := strconv.Atoi(c.PostForm("min_third_class"))
	if err != nil {
		http.Error(c.Writer, "Invalid min third class value", http.StatusBadRequest)
		return
	}

	// Update the student record in the database.
	for _, studentID := range studentIDs {
		studentID = strings.Trim(studentID, "[]")
		splitData := strings.Split(studentID, " ")
		for _, element := range splitData {
			studentID = element
			err := rdb.EditStudentClass(c, studentID, maxFirstClass, minFirstClass, maxSecondClass, minSecondClass, maxThirdClass, minThirdClass, c.Writer)
			if err != nil {
				http.Error(c.Writer, "Failed update student class", http.StatusBadRequest)
				return
			}
		}
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}

func GetEditStudentNameAndPassword(c *gin.Context) {
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

	ID := c.Query("student_id")
	if ID == "" {
		http.Error(c.Writer, "Student ID is required", http.StatusBadRequest)
		return
	}

	student, err := rdb.GetStudent(c, ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	item := struct {
		Title     string
		Message   string
		Account   string
		Student   *rdb.Student
		ID        string
		StudentID string
	}{
		Title:     "生徒名とパスワードの変更",
		Message:   student.Name + "さんの名前とパスワードを変更してください。",
		Account:   teacherUser.Name,
		Student:   student,
		ID:        "teacher",
		StudentID: student.ID,
	}

	// Render the edit page with the student data.
	er := pkg.Page("student_name_password_edit").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostEditStudentNameAndPassword(c *gin.Context) {
	ID := c.Query("student_id")
	newName := c.PostForm("new_name")
	newPassword := c.PostForm("new_password")

	if ID == "" || newName == "" || newPassword == "" {
		http.Error(c.Writer, "Student ID, new name and new password are required", http.StatusBadRequest)
		return
	}

	err := rdb.UpdateStudentNameAndPassword(c, ID, newName, newPassword)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}
