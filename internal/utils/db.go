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
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASS") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_HOST_PORT") + ")/" + os.Getenv("MYSQL_DB") + "?parseTime=True"
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
func HistoryDBInitTest() (db *gorm.DB, err error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3307)/CharisWorks?parseTime=true"
	log.Print("connect to ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Transaction{}, &TransactionItem{})

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
type InternalItem struct {
	Item Item `gorm:"embedded"`
	User User `gorm:"embedded"`
}
type InternalCart struct {
	Cart Cart `gorm:"embedded"`
	Item Item `gorm:"embedded"`
	User User `gorm:"embedded"`
}

type User struct {
	Id              string    `gorm:"id;type:varchar(100)"`
	DisplayName     string    `gorm:"display_name"`
	Description     string    `gorm:"description"`
	StripeAccountId string    `gorm:"stripe_account_id"`
	CreatedAt       time.Time `gorm:"created_at"`
}
type Item struct {
	Id                 string `gorm:"id"`
	ManufacturerUserId string `gorm:"manufacturer_user_id;type:varchar(100)"`
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
type InternalTransaction struct {
	Transaction      Transaction     `gorm:"embedded"`
	TransactionItems TransactionItem `gorm:"embedded"`
}
type Transaction struct {
	TransactionId   string    `gorm:"column:transaction_id;type:varchar(100);primaryKey"`
	PurchaserUserId string    `gorm:"purchaser_user_id;type:varchar(100)"`
	Email           string    `gorm:"email"`
	TrackingId      string    `gorm:"tracking_id;type:varchar(100)"`
	CreatedAt       time.Time `gorm:"created_at"`
	ZipCode         string    `gorm:"zip_code"`
	Address         string    `gorm:"address"`
	PhoneNumber     string    `gorm:"phone_number"`
	RealName        string    `gorm:"real_name"`
	Status          string    `gorm:"status"`
	TotalPrice      int       `gorm:"total_price"`
	TotalAmount     int       `gorm:"total_amount"`
}

type TransactionItem struct {
	Id                          int    `gorm:"id"`
	TransactionId               string `gorm:"transaction_id;type:varchar(100)"`
	ItemId                      string `gorm:"item_id;type:varchar(100)"`
	Name                        string `gorm:"name"`
	Price                       int    `gorm:"price"`
	Size                        int    `gorm:"size"`
	Quantity                    int    `gorm:"quantity"`
	Description                 string `gorm:"description"`
	Tags                        string `gorm:"tags"`
	ManufacturerUserId          string `gorm:"manufacturer_user_id;type:varchar(100)"`
	ManufacturerName            string `gorm:"manufacturer_name"`
	ManufacturerDescription     string `gorm:"manufacturer_description"`
	ManufacturerStripeAccountId string `gorm:"manufacturer_stripe_account_id;type:varchar(100)"`
	StripeTransferId            string `gorm:"stripe_transfer_id;type:varchar(100)"`
	Status                      string `gorm:"status"`
}
