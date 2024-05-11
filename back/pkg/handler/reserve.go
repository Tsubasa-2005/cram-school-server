package handler

import (
	"cram-school-reserve-server/back/infra/rdb"
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Reserve(w http.ResponseWriter, rq *http.Request) {
	user := pkg.CheckLogin(w, rq)
	if user == nil {

		return
	}
	studentUser, ok := user.(*rdb.Student)
	if !ok {

		return
	}

	if rq.Method == "GET" {

		ID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}

		form, err := rdb.GetForm(ID)
		if err != nil {
			http.Error(w, "Failed to get form", http.StatusInternalServerError)
			return
		}

		dates, err := rdb.CreateReservationPeriodList(*form)
		if err != nil {
			http.Error(w, "Failed to create reservation period list", http.StatusInternalServerError)
			return
		}

		formattedDates, err := rdb.FormatDates(dates)
		if err != nil {
			http.Error(w, "Failed to format dates", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		exceptionDates, err := rdb.CreateExceptionDates(*form)

		Max := 1
		Min := 1
		item := struct {
			Title          string
			Message        string
			Account        string
			ID             string
			Form           *rdb.Form
			Dates          []string
			ExceptionDates []string
			Max            int
			Min            int
		}{
			Title:          "フォームの予約",
			Message:        "フォームの予約を行ってください。",
			Account:        studentUser.Name,
			ID:             "student",
			Form:           form,
			Dates:          formattedDates,
			ExceptionDates: exceptionDates,
			Max:            Max,
			Min:            Min,
		}
		er := pkg.Page("reserve_form").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	} else if rq.Method == "POST" {
		ID, err := strconv.Atoi(rq.URL.Query().Get("form_id"))
		if err != nil {
			http.Error(w, "Invalid form ID", http.StatusBadRequest)
			return
		}
		var dates []string
		var times []string
		err = rq.ParseForm()
		if err != nil {
			return
		}
		reserves := rq.Form["reserve"]
		for _, reserve := range reserves {
			split := strings.Split(reserve, "_")
			date, time := split[0], split[1]
			if len(dates) > 0 && date == dates[len(dates)-1] {
				times[len(times)-1] = times[len(times)-1] + "-" + time
			} else {
				dates = append(dates, date)
				times = append(times, time)
			}
		}
		err = rdb.DeleteReservationByFormIDAndStudentID(ID, studentUser.ID)
		for i := 0; i < len(dates); i++ {
			reservation := rdb.Reservation{
				FormID:    ID,
				StudentID: studentUser.ID,
				Date:      dates[i],
				Time:      times[i],
			}
			err = rdb.Reserve(reservation)
			if err != nil {
				return
			}

		}

		http.Redirect(w, rq, "/student", http.StatusSeeOther)
	}
}
