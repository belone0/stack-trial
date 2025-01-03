package db

import (
  	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
	"github.com/belone0/stack-trial/go/balance/internal/models"
)


var DB *gorm.DB

func ConnectToDatabase(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = db

}

