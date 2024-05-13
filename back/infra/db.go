package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "data.sqlite3")
	if err != nil {
		return nil, err
	}
	return db, nil
}
