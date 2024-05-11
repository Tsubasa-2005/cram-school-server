package rdb

import (
	"cram-school-reserve-server/back/infra"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateForm(ID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	form := Form{
		ID:               ID,
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
func UpdateForm(ID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Model(&Form{}).Where("id=?", ID).Updates(Form{
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

	var form Form
	if err := db.Where("id = ?", formID).First(&form).Error; err != nil {
		return err
	}

	if err := db.Unscoped().Delete(&form).Error; err != nil {
		return err
	}

	return nil
}
