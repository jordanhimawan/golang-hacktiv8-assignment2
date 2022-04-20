package configs

import (
	"fmt"
	"sesi8-assignment/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 5432
	user   string = "postgres"
	pass   string = "koinworks2022*" // diisi password masing-masing
	dbname string = "db-go-sql"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Order{}) {
		db.AutoMigrate(&models.Order{})
	}

	if !db.HasTable(&models.Item{}) {
		db.AutoMigrate(&models.Item{})
	}

	return db
}
