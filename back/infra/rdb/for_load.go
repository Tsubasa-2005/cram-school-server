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

func GetReservationsByFormIDAndStudentID(FormID int, StudentID string) ([]Reservation, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var reservations []Reservation
	if err := db.Where("form_id = ? AND student_id = ?", FormID, StudentID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationByFormIDAndStudentIDAndDate(FormID int, StudentID string, Date string) (*Reservation, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var reservation Reservation
	if err := db.Where("form_id = ? AND student_id = ? AND date = ?", FormID, StudentID, Date).First(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func GetReservationByFormIDAndDate(FormID int, Date string) ([]Reservation, error) {
	db, err := infra.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var reservations []Reservation
	if err := db.Where("form_id = ? AND date = ?", FormID, Date).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
