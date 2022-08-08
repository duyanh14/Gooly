package services

import (
	"gooly/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
}

const dsn = "kvrevbyu_gooly:ZOG7OkvG05kI@tcp(45.252.249.31:3306)/kvrevbyu_gooly?charset=utf8&parseTime=True&loc=Local"

var db *gorm.DB

func (d Database) Conn() *gorm.DB {
	if db != nil {
		return db
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	db.AutoMigrate(&models.Site{})
	return db
}
