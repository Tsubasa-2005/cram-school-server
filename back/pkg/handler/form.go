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

func GetEditForm(c *gin.Context) {
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

	// Get the form ID from the query string.
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	// Get the form data from the database.
	form, err := rdb.GetForm(c, ID)
	if err != nil {
		http.Error(c.Writer, "The form ID is already in use. Please use a different ID.", http.StatusBadRequest)
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
	er := pkg.Page("edit_form").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostEditForm(c *gin.Context) {
	// Update the record in the database using the form data.
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	newName := c.PostForm("new_name")
	startDate := c.PostForm("start_date")
	endDate := c.PostForm("end_date")
	reserveStartDate := c.PostForm("reserve_start_date")
	reserveEndDate := c.PostForm("reserve_end_date")
	exceptionDates := c.PostFormArray("exception_date")

	// Join the exception dates into a single string with a comma as the delimiter
	exceptionDatesStr := strings.Join(exceptionDates, ",")
	err = rdb.UpdateForm(c, ID, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr)
	if err != nil {
		http.Error(c.Writer, "Failed to update form", http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}

func GetCreateForm(c *gin.Context) {
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

	// Get the form ID from the query string.
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	// Check if the formID already exists in the database
	_, err = rdb.GetForm(c, ID)
	if err == nil {
		// If the formID already exists, show an error message and return
		http.Error(c.Writer, "The form ID is already in use. Please use a different ID.", http.StatusBadRequest)
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
	er := pkg.Page("create_form").Execute(c.Writer, item)
	if er != nil {
		log.Fatal(er)
	}
}

func PostCreateForm(c *gin.Context) {
	// Update the record in the database using the form data.
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}
	newName := c.PostForm("new_name")
	startDate := c.PostForm("start_date")
	endDate := c.PostForm("end_date")
	reserveStartDate := c.PostForm("reserve_start_date")
	reserveEndDate := c.PostForm("reserve_end_date")
	exceptionDates := c.PostFormArray("exception_date")

	// Join the exception dates into a single string with a comma as the delimiter
	exceptionDatesStr := strings.Join(exceptionDates, ",")
	err = rdb.CreateForm(c, ID, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr)
	if err != nil {
		http.Error(c.Writer, "Failed to update form", http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}
func GetDeleteForm(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("form_id"))
	if err != nil {
		http.Error(c.Writer, "Invalid form ID", http.StatusBadRequest)
		return
	}

	err = rdb.DeleteForm(c, ID)
	if err != nil {
		http.Error(c.Writer, "Failed to delete form", http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher")
}
