package models

import (
	"github.com/jevitapearl/go_webProj/goBookInventory/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json: "author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() // from config module
	db = config.GetDB() // initialized DB instance and stores it in var db
	db.AutoMigrate(&Book{}) // Ensures schema matches struct
}

func (b *Book) CreateBook() *Book{ // Insert
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{ //
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}