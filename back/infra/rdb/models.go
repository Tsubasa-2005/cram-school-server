package rdb

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Student struct {
	PrimaryKey     uint `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ID             string `gorm:"unique;not null"`
	Name           string `gorm:"not null"`
	Password       string `gorm:"not null"`
	MaxFirstClass  int    `gorm:"not null"`
	MinFirstClass  int    `gorm:"not null"`
	MaxSecondClass int    `gorm:"not null"`
	MinSecondClass int    `gorm:"not null"`
	MaxThirdClass  int    `gorm:"not null"`
	MinThirdClass  int    `gorm:"not null"`
}

type Teacher struct {
	PrimaryKey uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ID         string `gorm:"unique;not null"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null"`
}

type Form struct {
	PrimaryKey       uint `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ID               int    `gorm:"unique;not null"`
	Name             string `gorm:"not null"`
	StartDate        string `gorm:"not null"`
	EndDate          string `gorm:"not null"`
	ReserveStartDate string `gorm:"not null"`
	ReserveEndDate   string `gorm:"not null"`
	ExceptionDate    string
}

type Reservation struct {
	PrimaryKey uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FormID     int    `gorm:"ForeignKey:FormID;References:ID;constraint:OnDelete:CASCADE"`
	StudentID  string `gorm:"ForeignKey:StudentID;References:ID;constraint:OnDelete:CASCADE"`
	Date       string `gorm:"not null"`
	Time       string
}
