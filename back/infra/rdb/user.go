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

func DeleteStudent(ID string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("id = ?", ID).Unscoped().Delete(&Student{}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTeacher(ID string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Where("id = ?", ID).Unscoped().Delete(&Teacher{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudentNameAndPassword(ID, newName, newPassword string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Model(&Student{}).Where("id = ?", ID).Updates(Student{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeacherNameAndPassword(ID, newName, newPassword string) error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Model(&Teacher{}).Where("id = ?", ID).Updates(Teacher{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func EditStudentClass(ID string, maxFirstClass, minFirstClass, maxSecondClass, minSecondClass, maxThirdClass, minThirdClass int, w http.ResponseWriter) error {
	// Update the student record in the database.
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	student := &Student{}
	if err := db.Model(student).Where("id = ?", ID).Updates(Student{
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
