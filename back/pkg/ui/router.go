package ui

import (
	"cram-school-reserve-server/back/pkg/handler"
	"net/http"
)

func SetupRouter() {
	// home
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		handler.Home(w)
	})
	// signup handling.
	http.HandleFunc("/signup", func(w http.ResponseWriter, rq *http.Request) {
		handler.Signup(w, rq)
	})
	// login handling.
	http.HandleFunc("/login", func(w http.ResponseWriter, rq *http.Request) {
		handler.Login(w, rq)
	})
	// logout handling.
	http.HandleFunc("/logout", func(w http.ResponseWriter, rq *http.Request) {
		handler.Logout(w, rq)
	})
	// student handling.
	http.HandleFunc("/student", func(w http.ResponseWriter, rq *http.Request) {
		handler.Student(w, rq)
	})
	// teacher handling.
	http.HandleFunc("/teacher", func(w http.ResponseWriter, rq *http.Request) {
		handler.Teacher(w, rq)
	})
	// delete_student handling.
	http.HandleFunc("/teacher/delete_student", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteStudent(w, rq)
	})
	// edit_student_class handling.
	http.HandleFunc("/teacher/edit_student_class", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditStudentClass(w, rq)
	})
	// update_student_class handling.
	http.HandleFunc("/teacher/update_student_class", func(w http.ResponseWriter, rq *http.Request) {
		handler.UpdateStudentClass(w, rq)
	})
	// edit_student_name_password handler.
	http.HandleFunc("/teacher/edit_student_name_password", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditStudentNameAndPassword(w, rq)
	})
	// delete_teacher handling.
	http.HandleFunc("/teacher/delete_teacher", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteTeacher(w, rq)
	})
	// edit_teacher_name_password handler.
	http.HandleFunc("/teacher/edit_teacher_name_password", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditTeacherNameAndPassword(w, rq)
	})
	// edit_form handling.
	http.HandleFunc("/teacher/edit_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditForm(w, rq)
	})
	// create_form handling.
	http.HandleFunc("/teacher/create_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.CreateForm(w, rq)
	})
	// delete_form handling.
	http.HandleFunc("/teacher/delete_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteForm(w, rq)
	})
	// reserve
	http.HandleFunc("/student/reserve", func(w http.ResponseWriter, rq *http.Request) {
		handler.Reserve(w, rq)
	})
	http.ListenAndServe("", nil)

}
