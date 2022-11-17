package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"whitepaper233.top/WMBBackend/model"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("./db/comments.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Comment{}) 
	db.AutoMigrate(&model.Reply{})
}

func GetDB() *gorm.DB {
	return db
}