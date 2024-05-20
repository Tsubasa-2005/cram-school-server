package rdb

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetStudent(c *gin.Context, studentID string) (*Student, error) {
	db := c.MustGet("db").(*gorm.DB)

	var student Student
	if err := db.Where("id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func GetTeacher(c *gin.Context, teacherID string) (*Teacher, error) {
	db := c.MustGet("db").(*gorm.DB)

	var teacher Teacher
	if err := db.Where("id = ?", teacherID).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func GetAllStudents(c *gin.Context) ([]Student, error) {
	db := c.MustGet("db").(*gorm.DB)

	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetAllTeachers(c *gin.Context) ([]Teacher, error) {
	db := c.MustGet("db").(*gorm.DB)

	var teachers []Teacher
	if err := db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func GetForm(c *gin.Context, formID int) (*Form, error) {
	db, err := gorm.Open("sqlite3", "data.sqlite3")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var form Form
	if err := db.Where("id = ?", formID).First(&form).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func GetReservationsByFormIDAndStudentID(c *gin.Context, FormID int, StudentID string) ([]Reservation, error) {
	db := c.MustGet("db").(*gorm.DB)

	var reservations []Reservation
	if err := db.Where("form_id = ? AND student_id = ?", FormID, StudentID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationByFormIDAndStudentIDAndDate(c *gin.Context, FormID int, StudentID string, Date string) ([]Reservation, error) {
	db := c.MustGet("db").(*gorm.DB)

	var reservation []Reservation
	if err := db.Where("form_id = ? AND student_id = ? AND date = ?", FormID, StudentID, Date).First(&reservation).Error; err != nil {
		return nil, err
	}
	return reservation, nil
}

func GetReservationByFormIDAndDate(c *gin.Context, FormID int, Date string) ([]Reservation, error) {
	db := c.MustGet("db").(*gorm.DB)

	var reservations []Reservation
	if err := db.Where("form_id = ? AND date = ?", FormID, Date).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
