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
	db.AutoMigrate(&User{}, &Item{}, &Cart{}, &Shipping{})

	return
}

type IDBUtils interface {
}
type Cart struct {
	Id       int    `gorm:"id"`
	UserId   string `gorm:"purchaser_user_id"`
	ItemId   string `gorm:"item_id"`
	Quantity int    `gorm:"quantity"`
}

type InternalCart struct {
	Cart       Cart   `gorm:"foreignKey:item_id"`
	ItemStock  int    `gorm:"stock"`
	ItemStatus string `gorm:"status"`
}

type User struct {
	Id              string    `gorm:"id"`
	DisplayName     string    `gorm:"display_name"`
	Description     string    `gorm:"description"`
	StripeAccountId string    `gorm:"stripe_account_id"`
	HistoryUserId   string    `gorm:"history_user_id"`
	CreatedAt       time.Time `gorm:"created_at"`
}
type Item struct {
	Id                 int      `gorm:"id"`
	ManufacturerUserId string   `gorm:"manufacturer_user_id"`
	HisToryItemId      string   `gorm:"history_item_id"`
	Name               string   `gorm:"name"`
	Price              int      `gorm:"price"`
	Status             string   `gorm:"status"`
	Stock              int      `gorm:"stock"`
	Size               int      `gorm:"size"`
	Description        string   `gorm:"description"`
	Tags               []string `gorm:"tags"`
}
type Shipping struct {
	Id            string `gorm:"id"`
	ZipCode       string `gorm:"zip_code"`
	Address1      string `gorm:"address_1"`
	Address2      string `gorm:"address_2"`
	Address3      string `gorm:"address_3" null:"true"`
	PhoneNumber   string `gorm:"phone_number"`
	FirstName     string `gorm:"first_name"`
	FirstNameKana string `gorm:"first_name_kana"`
	LastName      string `gorm:"last_name"`
	LastNameKana  string `gorm:"last_name_kana"`
}
