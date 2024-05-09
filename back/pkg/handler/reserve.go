package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Reserve(w http.ResponseWriter, rq *http.Request) {
	user := pkg.CheckLogin(w, rq)
	if user == nil {
		// Handle error here
		return
	}
	studentUser, ok := user.(*rdb.Student)
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
			http.Error(w, "Failed to get form", http.StatusInternalServerError)
			return
		}

		// Parse the reserve start and end dates.
		layout := "2006-01-02"
		start, err := time.Parse(layout, form.ReserveStartDate)
		if err != nil {
			http.Error(w, "Failed to parse reserve start date", http.StatusInternalServerError)
			return
		}
		end, err := time.Parse(layout, form.ReserveEndDate)
		if err != nil {
			http.Error(w, "Failed to parse reserve end date", http.StatusInternalServerError)
			return
		}

		// Generate all dates between the start and end dates.
		var dates []string
		for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
			dates = append(dates, d.Format("01-02"))
		}

		// Render the reserve page with the form data and the dates.
		item := struct {
			Title   string
			Message string
			Account string
			ID      string
			Form    *rdb.Form
			Dates   []string
		}{
			Title:   "フォームの予約",
			Message: "フォームの予約を行ってください。",
			Account: studentUser.Name,
			ID:      "student",
			Form:    form,
			Dates:   dates,
		}
		er := pkg.Page("reserve_form").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	} else if rq.Method == "POST" {
		// reserveのデータベースが決まってから実装
		http.Redirect(w, rq, "/student", http.StatusSeeOther)
	}
}
