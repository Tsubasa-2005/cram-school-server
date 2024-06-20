package rdb

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ReserveDay struct {
	ExceptionDay   bool
	Date           string
	Time           []string
	FirstReserved  []string
	SecondReserved []string
}

func FormatDate(startDateStr, endDateStr string) (string, string, bool, error) {
	// Parse the reserve start and end dates.
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDateStr)
	if err != nil {
		return "", "", false, fmt.Errorf("Failed to parse reserve start date: %w", err)
	}
	end, err := time.Parse(layout, endDateStr)
	if err != nil {
		return "", "", false, fmt.Errorf("Failed to parse reserve end date: %w", err)
	}

	// Check if the current date is within the reserve period.
	now := time.Now()
	isWithinReservePeriod := now.After(start) && now.Before(end)

	// Format the start and end dates in the new format.
	newLayout := "2006年1月2日"
	startDate := start.Format(newLayout)
	endDate := end.Format(newLayout)

	return startDate, endDate, isWithinReservePeriod, nil
}

func FormatDates(dates []string) ([]string, error) {
	var formattedDates []string
	layout := "01-02"
	newLayout := "1月2日"

	for _, dateStr := range dates {
		date, err := time.Parse(layout, dateStr)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse date: %w", err)
		}

		formattedDate := date.Format(newLayout)
		formattedDates = append(formattedDates, formattedDate)
	}

	return formattedDates, nil
}

func CreateReservationPeriodList(form Form) ([]string, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, form.ReserveStartDate)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse reserve start date: %w", err)
	}
	end, err := time.Parse(layout, form.ReserveEndDate)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse reserve end date: %w", err)
	}

	var dates []string
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("01-02"))
	}
	return dates, nil
}

func CreateExceptionDates(form Form) ([]string, error) {
	layout := "2006-01-02"
	newLayout := "01-02"
	exceptionDatesStr := strings.Split(form.ExceptionDate, ",")
	var exceptionDates []string
	var filteredExceptionDatesStr []string

	for _, dateStr := range exceptionDatesStr {
		if dateStr != "" {
			filteredExceptionDatesStr = append(filteredExceptionDatesStr, dateStr)
		}
	}

	exceptionDatesStr = filteredExceptionDatesStr
	for _, dateStr := range exceptionDatesStr {
		date, err := time.Parse(layout, dateStr)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse exception date: %w", err)
		}
		formattedDate := date.Format(newLayout)
		exceptionDates = append(exceptionDates, formattedDate)
	}

	return exceptionDates, nil
}

func CreateReserveDays(c *gin.Context, formID int, studentID string, formattedDates, exceptionDates []string) []ReserveDay {
	reserveDays := make([]ReserveDay, len(formattedDates))

	for i, date := range formattedDates {
		reserveDays[i].Date = date
		reserveDays[i].ExceptionDay = false

		for _, exceptionDate := range exceptionDates {
			if date == exceptionDate {
				reserveDays[i].ExceptionDay = true
				break
			}
		}
		reservations, err := GetReservationByFormIDAndStudentIDAndDate(c, formID, studentID, date)

		if err != nil {
			http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)

		}
		for _, reservation := range reservations {
			reservedTime := strings.Split(reservation.Time, "-")

			reserveDays[i].Time = reservedTime
		}

		if formID == 1 {
			reserveDays[i].FirstReserved = []string{}
			reserveDays[i].SecondReserved = []string{}
		} else if formID == 2 {
			reservations, err = GetReservationByFormIDAndStudentIDAndDate(c, 1, studentID, date)
			if err != nil {
				http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)
			}
			for _, reservation := range reservations {
				reservedTime := strings.Split(reservation.Time, "-")

				reserveDays[i].FirstReserved = reservedTime
			}
			reserveDays[i].SecondReserved = []string{}
		} else if formID == 3 {
			reservations, err = GetReservationByFormIDAndStudentIDAndDate(c, 1, studentID, date)
			if err != nil {
				http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)
			}
			for _, reservation := range reservations {
				reservedTime := strings.Split(reservation.Time, "-")

				reserveDays[i].FirstReserved = reservedTime
			}
			reservations, err = GetReservationByFormIDAndStudentIDAndDate(c, 2, studentID, date)
			if err != nil {
				http.Error(c.Writer, "Failed to get reservation", http.StatusInternalServerError)
			}
			for _, reservation := range reservations {
				reservedTime := strings.Split(reservation.Time, "-")

				reserveDays[i].SecondReserved = reservedTime
			}
		}
	}
	return reserveDays
}
