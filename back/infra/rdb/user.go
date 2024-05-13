package rdb

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateStudent(c *gin.Context, student Student) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func CreateTeacher(c *gin.Context, teacher Teacher) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(c *gin.Context, ID string) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", ID).Unscoped().Delete(&Student{}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTeacher(c *gin.Context, ID string) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", ID).Unscoped().Delete(&Teacher{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudentNameAndPassword(c *gin.Context, ID, newName, newPassword string) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Model(&Student{}).Where("id = ?", ID).Updates(Student{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeacherNameAndPassword(c *gin.Context, ID, newName, newPassword string) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Model(&Teacher{}).Where("id = ?", ID).Updates(Teacher{
		Name:     newName,
		Password: newPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func EditStudentClass(c *gin.Context, ID string, maxFirstClass, minFirstClass, maxSecondClass, minSecondClass, maxThirdClass, minThirdClass int, w http.ResponseWriter) error {
	db := c.MustGet("db").(*gorm.DB)

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
