package rdb

import (
	"fmt"
	"strings"
	"time"
)

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
	exceptionDates := strings.Split(form.ExceptionDate, ",")
	formattedExceptionDates, err := FormatDates(exceptionDates)
	if err != nil {
		return nil, err
	}
	return formattedExceptionDates, nil
}
