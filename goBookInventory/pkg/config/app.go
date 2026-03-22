package config

import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //holds DB connection
)

// opens a connection to the database
func Connect(){
	d, err := gorm.Open("mysql", "jevitapearl:pearl12345@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{ // return db connection so that other files can use it
	return db
}