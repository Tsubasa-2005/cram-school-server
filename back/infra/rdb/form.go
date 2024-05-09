package rdb

import (
	"cram-school-reserve-server/back/infra"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateForm(formID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	form := Form{
		ID:               formID,
		Name:             newName,
		StartDate:        startDate,
		EndDate:          endDate,
		ReserveStartDate: reserveStartDate,
		ReserveEndDate:   reserveEndDate,
		ExceptionDate:    exceptionDatesStr,
	}

	if err := db.Create(&form).Error; err != nil {
		return err
	}
	return nil
}

// UpdateForm updates a form's details in the database
func UpdateForm(formID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var form Form

	if err := db.Model(&form).Updates(Form{
		Name:             newName,
		StartDate:        startDate,
		EndDate:          endDate,
		ReserveStartDate: reserveStartDate,
		ReserveEndDate:   reserveEndDate,
		ExceptionDate:    exceptionDatesStr,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteForm(formID int) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("id = ?", formID).Delete(&Form{}).Error; err != nil {
		return err
	}
	return nil
}
