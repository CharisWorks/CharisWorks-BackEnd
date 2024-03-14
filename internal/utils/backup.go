package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectTobackUp() (db *gorm.DB, err error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3306)/CharisWorks?parseTime=true"
	log.Print("connect to ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Backup{})

	return
}

type Backup struct {
	gorm.Model
	id   int    `gorm:"primaryKey"`
	from string `gorm:"type:varchar(255)"`
}
