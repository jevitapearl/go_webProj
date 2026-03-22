package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Perform DB operations as if you are interacting with an object
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// runs sutomatically when this package is imported
func init() {
	config.Connect()        // connects the db
	db = config.GetDB()     // retrieves db instance
	db.AutoMigrate(&Book{}) // auto-migrates the schema for Book
}

func (b *Book) CreateBook() *Book {
	if b.ID == 0 { // checks if the entry exists
		db.Create(&b)
	}
	return b //final state of the data after operation
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) // passes a pointer because passing the var will result in sending a copy of the slice and no changes are reflected
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID=?", Id).Find(&getBook) // '?' is a placeholder. It prevents SQL injections by safely filtering the inputs from the user
	return &getBook, result
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
