package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB mysql connector
func DB() *gorm.DB {
	var db *gorm.DB
	var err error

	address := MysqlAddress()

	db, err = gorm.Open("mysql", "dealioapi:dealioteam@tcp("+address+")/dealio?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection db failed")
	}
	return db
}
