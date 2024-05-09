package rdb

import (
	"fmt"
	"time"
)

func ParseAndFormatDates(startDateStr, endDateStr string) (string, string, bool, error) {
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
