package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
)

func Teacher(w http.ResponseWriter, rq *http.Request) {
	user := pkg.CheckLogin(w, rq)
	if user == nil {
		// Handle error here
		return
	}
	teacherUser, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}
	teachers, err := rdb.GetAllTeachers()
	if err != nil {
		log.Fatal(err)
	}
	students, err := rdb.GetAllStudents()
	if err != nil {
		log.Fatal(err)
	}
	// Get the form data from the database.
	_, err1 := rdb.GetForm(1)
	_, err2 := rdb.GetForm(2)
	_, err3 := rdb.GetForm(3)
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
	er := pkg.Page("teacher").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func DeleteTeacher(w http.ResponseWriter, rq *http.Request) {
	user := pkg.CheckLogin(w, rq)
	if user == nil {
		// Handle error here
		return
	}
	teacherID := rq.URL.Query().Get("teacher_id")
	if teacherID == "" {
		http.Error(w, "Teacher ID is required", http.StatusBadRequest)
		return
	}

	err := rdb.DeleteTeacher(teacherID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, rq, "/teacher", http.StatusSeeOther)
}

func EditTeacherNameAndPassword(w http.ResponseWriter, rq *http.Request) {
	user := pkg.CheckLogin(w, rq)
	if user == nil {
		// Handle error here
		return
	}
	teacherUser, ok := user.(*rdb.Teacher)
	if !ok {
		// Handle error here
		return
	}
	if rq.Method == "GET" {
		teacherID := rq.URL.Query().Get("teacher_id")
		if teacherID == "" {
			http.Error(w, "Teacher ID is required", http.StatusBadRequest)
			return
		}

		teacher, err := rdb.GetTeacher(teacherID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		er := pkg.Page("teacher_name_password_edit").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	} else if rq.Method == "POST" {
		teacherID := rq.FormValue("teacher_id")
		newName := rq.FormValue("new_name")
		newPassword := rq.FormValue("new_password")

		if teacherID == "" || newName == "" || newPassword == "" {
			http.Error(w, "Teacher ID, new name and new password are required", http.StatusBadRequest)
			return
		}

		err := rdb.UpdateTeacherNameAndPassword(teacherID, newName, newPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, rq, "/teacher", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
