package commons

import (
	"database/sql"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetConnection() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	psgDB, err := db.DB()
	if err != nil {
		log.Fatal("error getting db connection", err)
	}
	return db, psgDB
}

func Migrate() {
	db, psgDB := GetConnection()
	defer func() {
		psgDB.Close()
	}()
	log.Println("Migrate")
	db.AutoMigrate(&models.Person{})
}
