package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func EditForm(w http.ResponseWriter, rq *http.Request) {
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
		// Get the form ID from the query string.
		formID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}
		// Get the form data from the database.
		form, err := rdb.GetForm(formID)
		if err != nil {
			http.Error(w, "The form ID is already in use. Please use a different ID.", http.StatusBadRequest)
			return
		}
		// Render the edit form page with the form data.
		item := struct {
			Title   string
			Message string
			Account string
			ID      string
			Form    *rdb.Form
		}{
			Title:   "フォームの編集",
			Message: "フォームの名前と日付を変更してください。例外日付がある場合は、追加してください。",
			Account: teacherUser.Name,
			ID:      "teacher",
			Form:    form,
		}
		er := pkg.Page("edit_form").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	} else if rq.Method == "POST" {
		// Update the record in the database using the form data.
		ID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}
		newName := rq.FormValue("new_name")
		startDate := rq.FormValue("start_date")
		endDate := rq.FormValue("end_date")
		reserveStartDate := rq.FormValue("reserve_start_date")
		reserveEndDate := rq.FormValue("reserve_end_date")
		exceptionDates := rq.Form["exception_date"]

		// Join the exception dates into a single string with a comma as the delimiter
		exceptionDatesStr := strings.Join(exceptionDates, ",")
		err = rdb.UpdateForm(ID, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr)
		if err != nil {
			http.Error(w, "Failed to update form", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, rq, "/teacher", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func CreateForm(w http.ResponseWriter, rq *http.Request) {
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
		// Get the form ID from the query string.
		ID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}
		// Check if the formID already exists in the database
		_, err = rdb.GetForm(ID)
		if err == nil {
			// If the formID already exists, show an error message and return
			http.Error(w, "The form ID is already in use. Please use a different ID.", http.StatusBadRequest)
			return
		}
		form := &rdb.Form{ID: ID}

		// Render the edit form page with the form data.
		item := struct {
			Title   string
			Message string
			Account string
			ID      string
			Form    *rdb.Form
		}{
			Title:   "フォームの編集",
			Message: "フォームの名前と日付を変更してください。例外日付がある場合は、追加してください。",
			Account: teacherUser.Name,
			ID:      "teacher",
			Form:    form,
		}
		er := pkg.Page("create_form").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	} else if rq.Method == "POST" {
		// Update the record in the database using the form data.
		ID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}
		newName := rq.FormValue("new_name")
		startDate := rq.FormValue("start_date")
		endDate := rq.FormValue("end_date")
		reserveStartDate := rq.FormValue("reserve_start_date")
		reserveEndDate := rq.FormValue("reserve_end_date")
		exceptionDates := rq.Form["exception_date"]

		// Join the exception dates into a single string with a comma as the delimiter
		exceptionDatesStr := strings.Join(exceptionDates, ",")
		err = rdb.CreateForm(ID, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr)
		if err != nil {
			http.Error(w, "Failed to update form", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, rq, "/teacher", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DeleteForm(w http.ResponseWriter, rq *http.Request) {
	formID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
	if err != nil {
		http.Error(w, "Invalid form ID", http.StatusBadRequest)
		return
	}

	err = rdb.DeleteForm(formID)
	if err != nil {
		http.Error(w, "Failed to delete form", http.StatusInternalServerError)
		return
	}
	// 消されたフォームの予約データは全て消す。予約データ保存方法が決まってから実装。
	http.Redirect(w, rq, "/teacher", http.StatusSeeOther)
}
