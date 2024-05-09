package rdb

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Student model
type Student struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	ID             string     `gorm:"primary_key"`
	Name           string     `gorm:"not null"`
	Password       string     `gorm:"not null"`
	MaxFirstClass  int        `gorm:"not null"`
	MinFirstClass  int        `gorm:"not null"`
	MaxSecondClass int        `gorm:"not null"`
	MinSecondClass int        `gorm:"not null"`
	MaxThirdClass  int        `gorm:"not null"`
	MinThirdClass  int        `gorm:"not null"`
}

// Teacher model
type Teacher struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ID        string     `gorm:"primary_key"`
	Name      string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
}

// Form model
type Form struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	ID               int        `gorm:"primary_key"`
	Name             string     `gorm:"not null"`
	StartDate        string     `gorm:"not null"`
	EndDate          string     `gorm:"not null"`
	ReserveStartDate string     `gorm:"not null"`
	ReserveEndDate   string     `gorm:"not null"`
	ExceptionDate    string
}

// Reservation model
//定義は相談して決定
