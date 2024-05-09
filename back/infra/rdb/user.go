package rdb

import (
	"cram-school-reserve-server/back/infra"

	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateStudent(student Student) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func CreateTeacher(teacher Teacher) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(studentID string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("id = ?", studentID).Delete(&Student{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTeacher(teacherID string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("id = ?", teacherID).Delete(&Teacher{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudentNameAndPassword(studentID, newName, newPassword string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Model(&Student{}).Where("id = ?", studentID).Updates(Student{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeacherNameAndPassword(teacherID, newName, newPassword string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Model(&Teacher{}).Where("id = ?", teacherID).Updates(Teacher{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func EditStudentClass(studentID string, maxFirstClass, minFirstClass, maxSecondClass, minSecondClass, maxThirdClass, minThirdClass int, w http.ResponseWriter) error {
	// Update the student record in the database.
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	student := &Student{}
	if err := db.Model(student).Where("id = ?", studentID).Updates(Student{
		MaxFirstClass:  maxFirstClass,
		MinFirstClass:  minFirstClass,
		MaxSecondClass: maxSecondClass,
		MinSecondClass: minSecondClass,
		MaxThirdClass:  maxThirdClass,
		MinThirdClass:  minThirdClass,
	}).Error; err != nil {
		return err
	}

	return nil
}
