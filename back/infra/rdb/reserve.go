package rdb

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Reserve(c *gin.Context, reservation Reservation) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Create(&reservation).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservationByFormIDAndStudentID(c *gin.Context, FormID int, StudentID string) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("form_id = ? AND student_id = ?", FormID, StudentID).Unscoped().Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservation(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	Date := time.Now().Format("2006-01-02")

	if err := db.Where("date < ?", Date).Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReservationCompletely(c *gin.Context) error {
	db := c.MustGet("db").(*gorm.DB)

	Date := time.Now().AddDate(0, -6, 0).Format("2006-01-02")

	if err := db.Where("date < ?", Date).Delete(&Reservation{}).Error; err != nil {
		return err
	}
	return nil
}
