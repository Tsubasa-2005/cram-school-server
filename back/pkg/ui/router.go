package ui

import (
	"cram-school-reserve-server/back/pkg/handler"
	"net/http"
)

func SetupRouter() {

	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		handler.Home(w)
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, rq *http.Request) {
		handler.Signup(w, rq)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, rq *http.Request) {
		handler.Login(w, rq)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, rq *http.Request) {
		handler.Logout(w, rq)
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, rq *http.Request) {
		handler.Student(w, rq)
	})

	http.HandleFunc("/teacher", func(w http.ResponseWriter, rq *http.Request) {
		handler.Teacher(w, rq)
	})

	http.HandleFunc("/teacher/delete_student", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteStudent(w, rq)
	})

	http.HandleFunc("/teacher/edit_student_class", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditStudentClass(w, rq)
	})

	http.HandleFunc("/teacher/update_student_class", func(w http.ResponseWriter, rq *http.Request) {
		handler.UpdateStudentClass(w, rq)
	})

	http.HandleFunc("/teacher/edit_student_name_password", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditStudentNameAndPassword(w, rq)
	})

	http.HandleFunc("/teacher/delete_teacher", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteTeacher(w, rq)
	})

	http.HandleFunc("/teacher/edit_teacher_name_password", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditTeacherNameAndPassword(w, rq)
	})

	http.HandleFunc("/teacher/edit_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.EditForm(w, rq)
	})

	http.HandleFunc("/teacher/create_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.CreateForm(w, rq)
	})

	http.HandleFunc("/teacher/delete_form", func(w http.ResponseWriter, rq *http.Request) {
		handler.DeleteForm(w, rq)
	})

	http.HandleFunc("/student/reserve", func(w http.ResponseWriter, rq *http.Request) {
		handler.Reserve(w, rq)
	})

	http.ListenAndServe("", nil)

}
