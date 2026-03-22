package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // holds DB connection
)

// opens a connection to the database
func Connect() {
	d, err := gorm.Open("mysql", "jevitapearl:pearl12345@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// Allows other packages to retrieve the active db connection
func GetDB() *gorm.DB { 
	return db
}
