package rdb

import (
	"cram-school-reserve-server/back/infra"
	"time"
)

func Reserve(reservation Reservation) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Create(&reservation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservationByFormIDAndStudentID(FormID int, StudentID string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("form_id = ? AND student_id = ?", FormID, StudentID).Unscoped().Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservation() error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	Date := time.Now().Format("2006-01-02")

	if err := db.Where("date < ?", Date).Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservationCompletely() error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	Date := time.Now().AddDate(0, -6, 0).Format("2006-01-02")

	if err := db.Where("date < ?", Date).Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}
