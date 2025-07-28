// order-service/db/db.go
package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var database *gorm.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	for i := 0; i < 10; i++ {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Waiting for DB to be ready... retrying in 2s")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic("Failed to connect to database after retries: " + err.Error())
	}

	DB = database
}