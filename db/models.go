package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open("postgres", "host=localhost port=5432 user=yourusername dbname=yourdatabase password=yourpassword sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database")
	}
	return DB
}

type Order struct {
	Id             int        `json:"id"`
	PhoneNumber    string     `json:"phoneNumber"`
	Email          string     `json:"email"`
	LinkedId       *int       `json:"linkedId"`
	LinkPrecedence string     `json:"linkPrecedence"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
}
