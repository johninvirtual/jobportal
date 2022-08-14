package database

import (
	"jobportal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Migrating...")
	db.AutoMigrate(&models.Job{})
	log.Println(("Migration done!"))

	DB = DbInstance{
		Db: db,
	}
}
