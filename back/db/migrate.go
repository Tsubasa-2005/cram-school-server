package db

import (
	"cram-school-reserve-server/back/infra/rdb"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Migrate() {
	db, err := gorm.Open("sqlite3", "data.sqlite3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Migrate the schema
	err = db.AutoMigrate(&rdb.Student{}).Error
	if err != nil {
		log.Fatalf("Failed to migrate Student model: %v", err)
		return
	}
	db.AutoMigrate(&rdb.Teacher{})
	db.AutoMigrate(&rdb.Form{})

	//RESERVATION modelは後ほど追加
	//db.AutoMigrate(&Reservation{})
}
