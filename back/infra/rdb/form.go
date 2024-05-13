package rdb

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateForm(c *gin.Context, ID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db := c.MustGet("db").(*gorm.DB)

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
func UpdateForm(c *gin.Context, ID int, newName, startDate, endDate, reserveStartDate, reserveEndDate, exceptionDatesStr string) error {
	db := c.MustGet("db").(*gorm.DB)

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

func DeleteForm(c *gin.Context, formID int) error {
	db := c.MustGet("db").(*gorm.DB)

	var form Form
	if err := db.Where("id = ?", formID).First(&form).Error; err != nil {
		return err
	}

	if err := db.Unscoped().Delete(&form).Error; err != nil {
		return err
	}

	return nil
}
