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
func DBInitTest() (db *gorm.DB, err error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3306)/CharisWorks?parseTime=true"
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
	Id              int    `gorm:"id"`
	PurchaserUserId string `gorm:"purchaser_user_id; type:varchar(100)"`
	ItemId          string `gorm:"item_id;type:varchar(100)"`
	Quantity        int    `gorm:"quantity"`
}

type InternalCart struct {
	Cart Cart `gorm:"embedded"`
	Item Item `gorm:"embedded"`
}

type User struct {
	Id              string    `gorm:"id;type:varchar(100)"`
	DisplayName     string    `gorm:"display_name"`
	Description     string    `gorm:"description"`
	StripeAccountId string    `gorm:"stripe_account_id"`
	HistoryUserId   int       `gorm:"history_user_id"`
	CreatedAt       time.Time `gorm:"created_at"`
}
type Item struct {
	Id                 string `gorm:"id"`
	ManufacturerUserId string `gorm:"manufacturer_user_id;type:varchar(100)"`
	HistoryItemId      int    `gorm:"history_item_id"`
	Name               string `gorm:"name"`
	Price              int    `gorm:"price"`
	Status             string `gorm:"status"`
	Stock              int    `gorm:"stock"`
	Size               int    `gorm:"size"`
	Description        string `gorm:"description"`
	Tags               string `gorm:"type:json"`
}
type Shipping struct {
	Id            string `gorm:"id"`
	ZipCode       string `gorm:"zip_code"`
	Address_1     string `gorm:"address_1"`
	Address_2     string `gorm:"address_2"`
	Address_3     string `gorm:"address_3" null:"true"`
	PhoneNumber   string `gorm:"phone_number"`
	FirstName     string `gorm:"first_name"`
	FirstNameKana string `gorm:"first_name_kana"`
	LastName      string `gorm:"last_name"`
	LastNameKana  string `gorm:"last_name_kana"`
}

type Transaction struct {
	Id                  int       `gorm:"id"`
	PurchaserUserId     string    `gorm:"purchaser_user_id"`
	TrackingId          string    `gorm:"tracking_id"`
	CreatedAt           time.Time `gorm:"created_at"`
	ZipCode             string    `gorm:"zip_code"`
	Address             string    `gorm:"address"`
	PhoneNumber         string    `gorm:"phone_number"`
	RealName            string    `gorm:"real_name"`
	Status              string    `gorm:"status"`
	StripeTransactionId string    `gorm:"stripe_transaction_id"`
	TotalPrice          int       `gorm:"total_price"`
	TotalAmount         int       `gorm:"total_amount"`
}

type TransactionItem struct {
	Id                      int    `gorm:"id"`
	TransactionId           int    `gorm:"transaction_id"`
	ItemId                  string `gorm:"item_id"`
	Name                    string `gorm:"name"`
	Price                   int    `gorm:"price"`
	Quantity                int    `gorm:"quantity"`
	Description             string `gorm:"description"`
	Tags                    string `gorm:"tags"`
	ManufacturerUserId      string `gorm:"manufacturer_user_id"`
	ManufacturerName        string `gorm:"manufacturer_name"`
	ManufacturerDescription string `gorm:"manufacturer_description"`
}
