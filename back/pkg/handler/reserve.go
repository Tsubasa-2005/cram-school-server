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

func GetReserve(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {

		return
	}
	studentUser, ok := user.(*rdb.Student)
	if !ok {

		return
	}

	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}

	form, err := rdb.GetForm(c, ID)
	if err != nil {
		http.Error(c.Writer, "Failed to get form", http.StatusInternalServerError)
		return
	}

	dates, err := rdb.CreateReservationPeriodList(*form)
	if err != nil {
		http.Error(c.Writer, "Failed to create reservation period list", http.StatusInternalServerError)
		return
	}

	formattedDates, err := rdb.FormatDates(dates)
	if err != nil {
		http.Error(c.Writer, "Failed to format dates", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	exceptionDates, err := rdb.CreateExceptionDates(*form)
	if err != nil {
		log.Println(err)
		http.Error(c.Writer, "Failed to create exception dates", http.StatusInternalServerError)
		return
	}

	formattedExceptionDates, err := rdb.FormatDates(exceptionDates)
	if err != nil {
		log.Println(err)
		http.Error(c.Writer, "Failed to format exception dates", http.StatusInternalServerError)
		return
	}

	reserveDays := rdb.CreateReserveDays(c, ID, studentUser.ID, formattedDates, formattedExceptionDates)
	Max := 0
	Min := 0
	if ID == 1 {
		Max = studentUser.MaxFirstClass
		Min = studentUser.MinFirstClass
	} else if ID == 2 {
		Max = studentUser.MaxSecondClass
		Min = studentUser.MinSecondClass
	} else if ID == 3 {
		Max = studentUser.MaxThirdClass
		Min = studentUser.MinThirdClass
	}

	item := struct {
		Title   string
		Message string
		Account string
		ID      string
		Form    *rdb.Form
		Dates   []rdb.ReserveDay
		Max     int
		Min     int
	}{
		Title:   "フォームの予約",
		Message: "フォームの予約を行ってください。",
		Account: studentUser.Name,
		ID:      "student",
		Form:    form,
		Dates:   reserveDays,
		Max:     Max,
		Min:     Min,
	}

	er := pkg.Page("reserve_form").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostReserve(c *gin.Context) {
	user := pkg.CheckLogin(c)
	if user == nil {

		return
	}
	studentUser, ok := user.(*rdb.Student)
	if !ok {

		return
	}
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	var dates []string
	var times []string
	reserves := c.PostFormArray("reserve")
	log.Println(reserves)
	for _, reserve := range reserves {
		log.Println(reserve)
		split := strings.Split(string(reserve), "_")
		date, time := split[0], split[1]
		if len(dates) > 0 && date == dates[len(dates)-1] {
			times[len(times)-1] = times[len(times)-1] + "-" + time
		} else {
			dates = append(dates, date)
			times = append(times, time)
		}
	}
	err = rdb.DeleteReservationByFormIDAndStudentID(c, ID, studentUser.ID)
	for i := 0; i < len(dates); i++ {
		reservation := rdb.Reservation{
			FormID:    ID,
			StudentID: studentUser.ID,
			Date:      dates[i],
			Time:      times[i],
		}
		err = rdb.Reserve(c, reservation)
		if err != nil {
			return
		}
	}
	c.Redirect(http.StatusSeeOther, "/student")
}
