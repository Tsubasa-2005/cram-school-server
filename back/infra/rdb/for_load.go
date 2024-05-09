package rdb

import (
	"cram-school-reserve-server/back/infra"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetStudent(studentID string) (*Student, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var student Student
	if err := db.Where("id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func GetTeacher(teacherID string) (*Teacher, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var teacher Teacher
	if err := db.Where("id = ?", teacherID).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

// GetAllStudents retrieves all students from the database
func GetAllStudents() ([]Student, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// GetAllTeachers retrieves all teachers from the database
func GetAllTeachers() ([]Teacher, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var teachers []Teacher
	if err := db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

// GetForm retrieves a form from the database by its ID
func GetForm(formID int) (*Form, error) {
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
