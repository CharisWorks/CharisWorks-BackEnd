package utils

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() (db *gorm.DB, err error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASS") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_HOST_PORT") + ")/" + os.Getenv("MYSQL_DB") + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Print("connect to ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})

	return
}

type IDBUtils interface {
}
type CartInDB struct {
	Id       int    `json:"id"`
	UserId   string `json:"purchaser_user_id"`
	ItemId   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type User struct {
	gorm.Model
	Id              string    `gorm:"id"`
	DisplayName     string    `gorm:"display_name"`
	Description     string    `gorm:"description"`
	StripeAccountId string    `gorm:"stripe_account_id"`
	HistoryUserId   string    `gorm:"history_user_id"`
	CreatedAt       time.Time `gorm:"created_at"`
}
type ItemInDB struct {
	Id                 int      `json:"id"`
	ManufacturerUserId string   `json:"manufacturer_user_id"`
	HisToryItemId      string   `json:"history_item_id"`
	Name               string   `json:"name"`
	Price              int      `json:"price"`
	Status             string   `json:"status"`
	Stock              int      `json:"stock"`
	Size               int      `json:"size"`
	Description        string   `json:"description"`
	Tags               []string `json:"tags"`
}
